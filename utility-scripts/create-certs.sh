#!/bin/bash

KEYS_DIR="keys" 
SERVER_DIR="server" 
mkdir -p "$KEYS_DIR"

## Cert # 1 -> For Code Signing
# Create a CA key and certificate
openssl genpkey -algorithm RSA -out $KEYS_DIR/ca-key.pem
openssl req -new -key $KEYS_DIR/ca-key.pem -x509 -days 365 -out $KEYS_DIR/ca-cert.pem -config $KEYS_DIR/ca.conf

# Create a code signing key and CSR
openssl genpkey -algorithm RSA -out $KEYS_DIR/codesign-key.pem
openssl req -new -key $KEYS_DIR/codesign-key.pem -out $KEYS_DIR/codesign.csr -config $KEYS_DIR/codesign.conf

# Sign the code signing CSR with CA
# codesign-key.pem is private key & codesign-cert.pem is certificate
openssl x509 -req -in $KEYS_DIR/codesign.csr -CA $KEYS_DIR/ca-cert.pem -CAkey $KEYS_DIR/ca-key.pem -CAcreateserial -out $KEYS_DIR/codesign-cert.pem -days 365 -extensions codesign_reqext -extfile $KEYS_DIR/codesign.conf


## Cert # 2 -> Self-signed Without Code Signing
# Create a private key
openssl genpkey -algorithm RSA -out $KEYS_DIR/private-key.pem

# Generate a self-signed X.509 certificate
openssl req -x509 -newkey rsa:4096 -key $KEYS_DIR/private-key.pem -out $KEYS_DIR/x509-certificate.pem
