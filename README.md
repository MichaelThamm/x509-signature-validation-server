# x509-validation-server
A server which will verify the signature of a bash script in order to decide if the script will be executed or not

# Requirements
* Golang - go1.18.1 linux/amd64
  * go get github.com/gin-gonic/gin
* Ubuntu - 22.04.1 LTS
  * openssl - 3.0.2

# Repository Explanation
* The __utility-scripts__ directory contains the bash scripts which do the following:
  1. __create-certs.sh__ -> Create the private and x509 certificate using the openssl library
  2. __sign-scripts.sh__ -> Sign the scripts in the usigned-scripts directory using the SHA-256 hash function and the private key
     * They are then placed in the test-scripts directory
  3. __build-server-binary.sh__ -> Build the binary
  4. __run-tests.sh__ -> Execute the test suite
* __However__, for convenience the bash script __run.sh__ in the project root does all of this for you

* To build the server binary:
  * In the project root directory, execute: 
  ```
  ./build-server-binary.sh
  ```

* Using the binary file:
  * In the project root directory, execute: 
  ```
  ./x509-validation-server
  ```
  * Within another shell, execute:
  ```
  curl -X POST --data-binary @<bash_script_path> http://localhost:8080/execute
  ```

# Test Case Inspection
| Script | Status | Purpose/Reason |
| -- | ---- | ------------ |
| 1 | Pass | Basic script |
| 2 | Pass | RSA functions |
| 3 | Pass | New lines |
| 4 | Pass | Special characters |
| 5 | Fail | 500 status code - Error executing script |
| 6 | Fail | RSA verification error |
| 7 | Fail | RSA verification error |
| 8 | Fail | RSA verification error |
| 9 | Fail | RSA verification error |
| 10 | Fail | 500 status code - Illegal signature format |

# Stretch Goals
* __Switch to stretch-goals branch to test stretch goals__