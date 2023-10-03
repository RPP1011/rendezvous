//go:build ignore
// +build ignore

//go:generate go run generate.go -rproto messages/reliable_packets.proto -uproto messages/unreliable_packets.proto -output .\shared\packet_handler.go  -template .\shared\handler.tmpl

package main

import (
	"flag"
	"os"
	"text/template"

	"github.com/RPP1011/rendezvous/packets"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type MessageName struct {
	Name string
}

func main() {

	reliableProto := flag.String("rproto", "example.proto", "Path to the reliable packet .proto file")
	unreliableProto := flag.String("uproto", "example.proto", "Path to the unreliable packet .proto file")
	templateFilePath := flag.String("template", "handler.tmpl", "Path to the template file")
	outputFilePath := flag.String("output", "generated_handler.go", "Path to the output Go file")
	flag.Parse()

	// Load the .proto file
	// Extract information about the ReliablePacket message and its oneof field
	// Convert the oneof fields to a slice of pointers to FieldDescriptor

	reliable_fields := getOneOfFields(reliableProto)
	unreliable_fields := getOneOfFields(unreliableProto)

	// Cast the field's message's name to a string
	reliableMessageNames := getFieldNames(reliable_fields)
	unreliableMessageNames := getFieldNames(unreliable_fields)

	// Prepare data for the template
	data := struct {
		PackageName            string
		ReliableMessageNames   []MessageName
		UnreliableMessageNames []MessageName
	}{
		PackageName:            "shared",
		ReliableMessageNames:   reliableMessageNames,
		UnreliableMessageNames: unreliableMessageNames,
	}

	// Load the template
	tmpl, err := template.ParseFiles(*templateFilePath)
	if err != nil {
		panic(err)
	}

	// Execute the template and write the generated code to a file
	file, err := os.Create(*outputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

func getFieldNames(fields protoreflect.FieldDescriptors) []MessageName {
	messageNames := make([]MessageName, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		messageNames[i].Name = string(fields.Get(i).Message().Name())
	}
	return messageNames
}

func getOneOfFields(reliableProto *string) protoreflect.FieldDescriptors {
	fd, err := protoregistry.GlobalFiles.FindFileByPath(*reliableProto)
	if err != nil {
		panic(err)
	}

	messages := make([]protoreflect.MessageDescriptor, fd.Messages().Len())
	for i := 0; i < fd.Messages().Len(); i++ {
		messages[i] = fd.Messages().Get(i)
	}

	var messageDesc protoreflect.MessageDescriptor = nil
	for md := range messages {
		opts := messages[md].Options()
		if proto.GetExtension(opts, packets.E_IsGeneric) != nil {
			messageDesc = messages[md]
			break
		}
	}

	oneof := messageDesc.Fields().Get(0).ContainingOneof()
	oneofFields := oneof.Fields()

	oneofFieldSlice := make([]protoreflect.FieldDescriptor, oneofFields.Len())
	for i := 0; i < oneofFields.Len(); i++ {
		oneofFieldSlice[i] = oneofFields.Get(i)
	}
	return oneofFields
}
