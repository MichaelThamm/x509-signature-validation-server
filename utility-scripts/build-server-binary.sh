#!/bin/bash

cd ../server

# Set the output binary name
OUTPUT_BINARY="x509-signature-validation-server"

# Build the Go server
go build -o $OUTPUT_BINARY
