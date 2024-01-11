#!/bin/bash

# Create a private key
openssl genpkey -algorithm RSA -out private-key.pem

# Create a certificate signing request to openssl:
# openssl req -new -key private-key.pem -out csr.pem

# Self-sign the certificate
openssl x509 -req -in csr.pem -signkey private-key.pem -out x509-certificate.pem