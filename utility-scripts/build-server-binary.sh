#!/bin/bash

# Set the output binary name
OUTPUT_BINARY="your_server_binary_name"

# Build the Go server
go build -o $OUTPUT_BINARY main.go

# Optionally, move the binary to a specific directory
# mkdir -p bin
# mv $OUTPUT_BINARY bin/
