package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const CERT_PATH string = "keys/x509-certificate.pem"

func readScript(filename string) (signature string, result string, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	// Convert the content to string
	scriptContent := string(content)

	// Split the content into lines
	lines := strings.Split(scriptContent, "\n")

	// Extract the signature from the first line
	if len(lines) > 0 {
		signature = strings.TrimSpace(lines[0])
	}

	// Join the remaining lines to reconstruct the script content
	result = strings.Join(lines[1:], "\n")

	return
}

func verifySignature(signature, scriptContent, certificatePath string) (bool, error) {
	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	// Read the X.509 certificate
	certPEM, err := ioutil.ReadFile(certificatePath)
	if err != nil {
		return false, err
	}

	// Parse the certificate PEM block
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return false, fmt.Errorf("failed to parse certificate PEM")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse certificate: %v", err)
	}

	// Extract the public key from the certificate
    switch pubKey := cert.PublicKey.(type) {
    case *rsa.PublicKey:
        // Continue with verification using RSA public key
        hashed := scriptHash(scriptContent)
        err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signatureBytes)
        return err == nil, nil
    default:
        return false, fmt.Errorf("unsupported public key type: %T", pubKey)
    }
}

func scriptHash(scriptContent string) []byte {
	// In this example, we use SHA-256 hash function
	hashed := sha256.Sum256([]byte(scriptContent))
	return hashed[:]
}

func executeScript(script string) ([]byte, error) {
	cmd := exec.Command("/bin/bash", "-c", script)
	output, err := cmd.CombinedOutput()
	return output, err
}

func main() {
	if len(os.Args) != 2 {
		err := fmt.Errorf("Usage: ./main <bash_script>")
		fmt.Println("Error reading inputs:", err)
		return
	}

	scriptFile := os.Args[1]
	
	signature, scriptContent, err := readScript(scriptFile)
	
	if err != nil {
		fmt.Println("Error reading inputs:", err)
		os.Exit(1)
	}

	isValid, err := verifySignature(signature, scriptContent, CERT_PATH)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Signature is", map[bool]string{true: "valid", false: "invalid"}[isValid])

	if isValid {
		output, err := executeScript(scriptContent)
		if err != nil {
			// Handle the error
			fmt.Println("Error executing script:", err)
		} else {
			// Process the output
			fmt.Println("Script output:")
			fmt.Println(string(output))
		}
	}
}
