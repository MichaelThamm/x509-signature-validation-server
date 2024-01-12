# x509-validation-server
A server which will verify the signature of a bash script in order to decide if the script will be executed or not

# Requirements
* Golang - go1.18.1 linux/amd64
  * go get github.com/gin-gonic/gin
* Ubuntu - 22.04.1 LTS
  * openssl - 3.0.2

# Repository Explanation
* The utility-scripts folder contains the shell scripts which do the following:
  1. __create-certs.sh__ -> Create the private and x509 certificate using the openssl library
  2. __sign-scripts.sh__ -> Sign the scripts in the usigned-scripts directory using the SHA-256 hash function and the private key
    * They are then placed in the signed-scripts directory
  3. __build-server-binary.sh__ -> Build the binary

* To build the server binary:
  * Execute _./build-server-binary.sh_ in the project root directory

* Using the binary file:
  * Execute _./x509-validation-server_ in the project root directory 
  * Execute _cat signed-scripts/<choose a script>.sh | curl -X POST --data-binary @- http://localhost:8080/execute_

# Test Case Inspection
| Pass | Fail |
| --- | ----------- |
| Header | Title |
| Paragraph | Text |