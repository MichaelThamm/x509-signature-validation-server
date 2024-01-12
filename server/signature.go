package server

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
)

// verifySignature verifies the signature using the public key from the certificate
func VerifySignature(signature string, script string, cert *x509.Certificate) (bool, error) {
	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.Printf("Error decoding signature: %v\n", err)
		return false, err
	}
	
	// Extract the public key from the certificate
	switch pubKey := cert.PublicKey.(type) {
	case *rsa.PublicKey:
		// Continue with verification using RSA public key
		hashed := ScriptHash(script)
		err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, signatureBytes)
		if err != nil {
			log.Println("Verification failed:", err)
			return false, nil
		}
		return true, nil
	default:
		return false, fmt.Errorf("unsupported public key type: %T", pubKey)
	}
}

func ScriptHash(scriptContent string) []byte {
	hashed := sha256.Sum256([]byte(scriptContent))
	return hashed[:]
}
