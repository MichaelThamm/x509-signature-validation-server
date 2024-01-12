package server

import (
	"crypto/x509"
    "fmt"
	"io/ioutil"
    "net"
)

// startServer starts the server and listens for incoming connections
func StartServer(cert *x509.Certificate, port int) {
    // Create a listener on the specified port
    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        fmt.Println("Error creating listener:", err)
        return
    }
    defer listener.Close()

    fmt.Printf("Server listening on port %d...\n", port)

    for {
        // Accept incoming connections
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        // Handle the connection in a goroutine (concurrent)
        go handleConnection(conn, cert)
    }
}

// handleConnection handles incoming connections and data
func handleConnection(conn net.Conn, cert *x509.Certificate) {
    defer conn.Close()

    // Read the incoming data (bash script)
    script, err := ioutil.ReadAll(conn)
    if err != nil {
        fmt.Println("Error reading script:", err)
        return
    }
    
    fmt.Println("script:", script)
    // WRITE code to split the signature from the script
    // signature := "123321"
    // // Verify the signature and execute the script
    // isValid := verifySignature(signature, script, cert)
    // if isValid {
    //     output, err := executeScript(script)
    //     if err != nil {
    //         fmt.Println("Error executing script:", err)
    //         return
    //     }

    //     // Send the status line and script output back through the connection
    //     // (You need to implement this based on your protocol)
    //     _, err = conn.Write([]byte("Valid\n"))
    //     if err != nil {
    //         fmt.Println("Error writing status:", err)
    //         return
    //     }
    //     _, err = conn.Write(output)
    //     if err != nil {
    //         fmt.Println("Error writing output:", err)
    //         return
    //     }
    // } else {
    //     // Send an invalid status line back through the connection
    //     // (You need to implement this based on your protocol)
    //     _, err := conn.Write([]byte("Invalid\n"))
    //     if err != nil {
    //         fmt.Println("Error writing status:", err)
    //         return
    //     }
    // }
}
