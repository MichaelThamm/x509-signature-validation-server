#!/bin/bash

KEYS_DIR="keys"
mkdir -p "$KEYS_DIR"

# Create a private key
openssl genpkey -algorithm RSA -out $KEYS_DIR/private-key.pem

# Extract the public key from the private key
openssl rsa -pubout -in $KEYS_DIR/private-key.pem -out $KEYS_DIR/public-key.pem

# Generate a self-signed X.509 certificate
openssl req -x509 -newkey rsa:4096 -key $KEYS_DIR/private-key.pem -out $KEYS_DIR/x509-certificate.pem

