package server

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

// verifySignature verifies the signature using the public key from the certificate
func VerifySignature(checkcodesign bool, signature string, script string, cert *x509.Certificate) (bool, error) {
	if checkcodesign && !HasCodeSigningExtension(cert) {
		log.Printf("Error in EKU")
		return false, errors.New("Certificate does not have the required code signing Extended Key Usage")
	}
	
	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.Printf("Failed to decode signature: %v\n", err)
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
		return false, fmt.Errorf("Unsupported public key type: %T", pubKey)
	}
}

func ScriptHash(scriptContent string) []byte {
	hashed := sha256.Sum256([]byte(scriptContent))
	return hashed[:]
}

// Check if the certificate has the Code Signing extension
func HasCodeSigningExtension(cert *x509.Certificate) bool {
	for _, extKeyUsage := range cert.ExtKeyUsage {
		// x509.ExtKeyUsageCodeSigning == 3 (for code signing)
		if extKeyUsage == x509.ExtKeyUsageCodeSigning {
			return true
		}
	}
	return false
}