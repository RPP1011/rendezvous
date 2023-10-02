//go:build ignore
// +build ignore

//go:generate go run generate.go -proto messages/packets.proto -output .\shared\packet_handler.go -template .\shared\handler.tmpl

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

	protoFilePath := flag.String("proto", "example.proto", "Path to the .proto file")
	templateFilePath := flag.String("template", "handler.tmpl", "Path to the template file")
	outputFilePath := flag.String("output", "generated_handler.go", "Path to the output Go file")
	flag.Parse()

	// Load the .proto file
	fd, err := protoregistry.GlobalFiles.FindFileByPath(*protoFilePath)
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

	// Extract information about the ReliablePacket message and its oneof field
	oneof := messageDesc.Fields().Get(1).ContainingOneof()
	oneofFields := oneof.Fields()

	// Convert the oneof fields to a slice of pointers to FieldDescriptor
	oneofFieldSlice := make([]protoreflect.FieldDescriptor, oneofFields.Len())
	for i := 0; i < oneofFields.Len(); i++ {
		oneofFieldSlice[i] = oneofFields.Get(i)
	}

	messageNames := make([]MessageName, oneofFields.Len())
	for i := 0; i < oneofFields.Len(); i++ {
		// Cast the field's message's name to a string
		messageNames[i].Name = string(oneofFields.Get(i).Message().Name())

	}

	// Prepare data for the template
	data := struct {
		PackageName string
		OneofFields []MessageName
	}{
		PackageName: "shared",
		OneofFields: messageNames,
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
