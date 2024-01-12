package main

import (
	"fmt"
	"x509-validation-server/server"
)

const CERT_PATH string = "keys/x509-certificate.pem"

func main() {

	// Load the certificate
	cert, err := server.LoadCertificate(CERT_PATH)
	if err != nil {
		fmt.Println("Error loading certificate:", err)
		return
	}

	// Start the server and handle incoming requests
	server.StartServer(cert)
}
