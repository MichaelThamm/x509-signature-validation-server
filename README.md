# x509-signature-validator-server
A server which will verify the signature of a bash script in order to decide if the script will be executed or not

# Requirements
* Golang - go1.18.1 linux/amd64
* Ubuntu - 22.04.1 LTS
  * openssl - 3.0.2
* [XCA](https://apps.microsoft.com/detail/9N9K02V7XR1B?hl=en-US&gl=US)

# XCA Explanation
Let's say I want to browse to [www.markdownguide.org](https://www.markdownguide.org/extended-syntax/) to learn about markdown syntax for writing a nice README.md for this repo.
* I can export the certificate and inspect it using XCA to provide the trust chain
* ![Alt text](markdown-pem.png)

# Code Explanation
* The utility-scripts folder contains the shell scripts which do the following:
  * Create the private, public, and x509 certificate using the openssl library
  * Sign the scripts in the usigned-scripts directory using the SHA-256 hash function and the private key
    * They are then placed in the signed-scripts directory
* The main.go file accepts 2 arguments:
  * script file name
	* certification path

# Test Case Inspection
| Pass | Fail |
| --- | ----------- |
| Header | Title |
| Paragraph | Text |