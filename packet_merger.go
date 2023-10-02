//go:build ignore
// +build ignore

//go:generate go run packet_merger.go -proto messages/packets.proto

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	protoFilePath := flag.String("proto", "messages/packets.proto", "Path to the .proto file")
	containerName := flag.String("containerName", "ReliablePacket", "name of the container message")

	file, err := os.Open(*protoFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered reader from the file
	reader := bufio.NewReader(file)

	rep := reporter.NewReporter(errorReporter, warningReporter)
	handler := reporter.NewHandler(rep)
	node, parse_err := parser.Parse(*protoFilePath, reader, handler)
	if parse_err != nil {
		panic(parse_err)
	}

	// Extract information about the ReliablePacket message and its oneof field
	result, ast_err := parser.ResultFromAST(node, false, handler)
	if ast_err != nil {
		panic(ast_err)
	}

	proto := result.FileDescriptorProto()
	messages := (*proto).MessageType

	// Make list to store all but container message
	//var filteredMessages []*descriptorpb.DescriptorProto = make([]*descriptorpb.DescriptorProto, len(messages)-1)
	container, filteredMessages := retrieveGenericMessages(messages)
	if container == nil {
		panic("No container message found")
	}

	// Extract names from filtered messages
	replace_container(*protoFilePath, *containerName, filteredMessages)

}

func warningReporter(ePos reporter.ErrorWithPos) {
	println(ePos.Error())
}

func errorReporter(ePos reporter.ErrorWithPos) error {
	println(ePos.Error())
	return ePos.Unwrap()
}

func retrieveGenericMessages(messages []*descriptorpb.DescriptorProto) (*descriptorpb.DescriptorProto, []string) {
	// Make list to store all but container message
	var container *descriptorpb.DescriptorProto = nil
	var filteredMessages []string = make([]string, 0)
	for _, message := range messages {
		if message.Options != nil {
		Outerloop:
			// Iterate through the options
			for _, option := range message.Options.UninterpretedOption {
				for _, namePart := range option.Name {
					if *namePart.NamePart == "is_generic" {
						container = message
						break Outerloop
					}
				}
			}
		}

		if message != container && message != nil {
			filteredMessages = append(filteredMessages, message.GetName())
		}
	}
	println("Container: " + container.GetName())
	return container, filteredMessages
}

func replace_container(targetFilePath string, containerName string, messages []string) {
	file, err := os.Open(targetFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result []string
	var inside bool
	var braceCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, fmt.Sprintf("message %s", containerName)) {

			inside = true
		}
		if inside {
			braceCount += strings.Count(line, "{")
			braceCount -= strings.Count(line, "}")

			if braceCount == 0 {
				inside = false
				result = append(result, fmt.Sprintf("message %s {", containerName)) // Add replacement text
				result = append(result, "    option (is_generic) = true;")
				result = append(result, "    oneof value {")

				for i, messageName := range messages {
					name := messageName
					lowercaseName := decapitalize(messageName)
					result = append(result, fmt.Sprintf("        %s %s = %d;", name, lowercaseName, i+1))
				}
				result = append(result, "    }")
				result = append(result, "}")
				continue
			}
		}
		if !inside {
			result = append(result, line)
		}
	}

	output := strings.Join(result, "\n")
	os.WriteFile(targetFilePath, []byte(output), 0644)
}

func decapitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToLower(str[:1]) + str[1:]
}
