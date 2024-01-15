package server

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

// LoadCertificate loads the x509 certificate from the specified path
func LoadCertificate(certPath string) (*x509.Certificate, error) {
	// Read the X.509 certificate
	certPEM, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	
	// Parse the certificate PEM block
	block, _ := pem.Decode(certPEM)
	if block == nil {
		log.Println("Error: failed to parse certificate PEM")
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Println("failed to parse certificate: %v", err)
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}
	return cert, nil
}
