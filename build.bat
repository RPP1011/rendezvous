@echo off

rem Clear build directory
rmdir /s /q build

rem Populate container type
go generate packet_merger.go

rem Generate Go code from protobuf file
protoc --go_out=. messages/packets.proto

rem Run Go generate to generate code from templates
go generate generate.go

rem Build the main Go program
go build -o build/ main.go