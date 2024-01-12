#!/bin/bash

# Set the output binary name
OUTPUT_BINARY="x509-validation-server"

# Build the Go server
go build -o $OUTPUT_BINARY ./cmd
