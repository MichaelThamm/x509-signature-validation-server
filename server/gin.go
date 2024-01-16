package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const CERT_DIR string = "keys"
const CERT_PATH string = CERT_DIR + "/codesign-cert.pem"

func StartServer() {
	router := gin.Default()

	// Define the bash execution route
	router.POST("/execute", func(c *gin.Context) {
		// Read the request arguments
		scriptID := c.Query("id")
		if scriptID == "" {
			scriptID = "1"
		}
		checkCodeSign := c.Query("code-sign") == "true"
		checkKeyInDir := c.Query("key-dir") == "true"
		
		// Load the certificate
		if !checkKeyInDir {
			certPath := CERT_PATH
			scriptContents, err := GetRawData(c)
			if err != nil { return }
			ExecuteFromCert(checkCodeSign, scriptID, certPath, scriptContents, c)
		} else {
			// Read the contents of the parent directory
			parentDirectory := CERT_DIR
			entries, err := os.ReadDir(parentDirectory)
			if err != nil {
				log.Printf("Error reading the certificate directory %s: %v\n", parentDirectory, err)
				return
			}
		
			// Iterate through certificates in the directory
			scriptContents, err := GetRawData(c)
			if err != nil { return }
			for _, entry := range entries {
				if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".pem") {
					continue
				}
				certPath := filepath.Join(parentDirectory, entry.Name())
				ExecuteFromCert(checkCodeSign, scriptID, certPath, scriptContents, c)
			}
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

func ExecuteFromCert(checkCodeSign bool, scriptID string, certPath string, entireScript string, c *gin.Context) {
	// Load the certificate from the file path
	cert, err := LoadCertificate(certPath)
	if err != nil {
		log.Println("Error loading certificate:", err)
		return
	}

	// Verify the signature
	signature, script, err := SplitScript(entireScript)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error splitting script")
		return
	}
	isValid, err := VerifySignature(checkCodeSign, signature, script, cert)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error verifying signature")
		return
	}

	log.Println("Validity:", isValid)
	if isValid {
		// Execute the script
		output, err := ExecuteScript(string(script))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error executing script")
			return
		}
		// Send the status line and script output back through the response
		c.String(http.StatusOK, "The script (ID %v) is VALID (using certificate: %s) and was executed:\n%s", scriptID, certPath, output)
	} else {
		// Send an invalid status line back through the response
		c.String(http.StatusOK, "The script (ID %v) is INVALID (using certificate: %s) and will NOT be executed\n", scriptID, certPath)
	}
}

func GetRawData(c *gin.Context) (strDataContents string, err error) {
	// Read the incoming data (bash script)
	dataContents, err := c.GetRawData()
	if err != nil {
		log.Println("Error reading script")
		return "", errors.New("Error reading script")
	}
	strDataContents = string(dataContents)
	return strDataContents, nil
}