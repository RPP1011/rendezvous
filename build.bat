@echo off

rem Clear build directory
rmdir /s /q build
del /F "packets\packet_handler.go"

rem Populate container types
go run packet_merger.go -proto messages/reliable_packets.proto -containerName ReliablePacket
go run packet_merger.go -proto messages/unreliable_packets.proto -containerName UnreliablePacket

rem Generate Go code from protobuf file
protoc --go_out=. messages/shared.proto
protoc --go_out=. messages/reliable_packets.proto
protoc --go_out=. messages/unreliable_packets.proto

rem Run Go generate to generate code from templates
go generate generate.go

rem Build the main Go program
go build -o build/ main.go