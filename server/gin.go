package server

import (
	"crypto/x509"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func StartServer(cert *x509.Certificate) {
	router := gin.Default()

	// Define the bash execution route
	router.POST("/execute", func(c *gin.Context) {
		// Read the incoming data (bash script)
		entireScript, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error reading script")
			return
		}
		signature, script, err := SplitScript(string(entireScript))

        // Verify the signature
		isValid, err := VerifySignature(signature, script, cert)
		log.Println("Validity", isValid)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error verifying signature")
			return
		}

		if isValid {
			// Execute the script
			output, err := ExecuteScript(string(script))
			if err != nil {
				c.String(http.StatusInternalServerError, "Error executing script")
				return
			}

			// Send the status line and script output back through the response
			c.String(http.StatusOK, "The script is VALID and was executed:\n%s", output)
		} else {
			// Send an invalid status line back through the response
			c.String(http.StatusOK, "The script is INVALID and will NOT be executed\n")
		}
	})

	// Run the server
	router.Run(":8080")
}

func SplitScript(script string) (signature string, restOfScript string, err error) {
	// Split the content into lines
	lines := strings.Split(script, "\n")

	// Extract the signature from the first line
	if len(lines) == 0 {
		return "", "", errors.New("Script with signature is empty")
	}
	signature = strings.TrimSpace(lines[0])
	// Join the remaining lines
	restOfScript = strings.Join(lines[1:], "\n")
	return signature, restOfScript, nil
}