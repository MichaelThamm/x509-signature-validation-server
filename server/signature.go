package server

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// verifySignature verifies the signature using the public key from the certificate
func VerifySignature(signature, scriptContent string, cert *x509.Certificate) (bool, error) {
	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	
	// Extract the public key from the certificate
	switch pubKey := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		// Continue with verification using RSA public key
		hashed := ScriptHash(scriptContent)
		err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signatureBytes)
		return err == nil, nil
	default:
		return false, fmt.Errorf("unsupported public key type: %T", pubKey)
	}
}

func ScriptHash(scriptContent string) []byte {
	// In this example, we use SHA-256 hash function
	hashed := sha256.Sum256([]byte(scriptContent))
	return hashed[:]
}
