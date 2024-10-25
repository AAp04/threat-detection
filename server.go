package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "strings"
    "os"
)

// Function to handle incoming client connections
func handleClient(conn net.Conn) {
    defer conn.Close()

    // Create a buffer to read incoming data
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        log.Println("Error reading from connection:", err)
        return
    }

    message := string(buffer[:n])

    // Log received message (simulated threat detection)
    fmt.Printf("Received packet: %s", message)

    // Perform basic threat detection
    if strings.Contains(message, "192.168.1.1") { // Example check
        log.Println("Suspicious activity detected from 192.168.1.1")
        conn.Write([]byte("Threat detected!\n"))
    } else {
        log.Println("No threats detected")
        conn.Write([]byte("Packet received, no threat detected.\n"))
    }

    // Log packet to file
    logPacket(message)
}

// Function to log packet data to a file
func logPacket(packet string) {
    f, err := os.OpenFile("packets.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("Error opening log file:", err)
        return
    }
    defer f.Close()

    logger := log.New(f, "INFO: ", log.LstdFlags)
    logger.Println(packet)
}

// Function to handle HTTP requests (for testing purposes)
func handleHttpRequest(w http.ResponseWriter, r *http.Request) {
    log.Println("Received HTTP request:", r.Method, r.URL.Path)
    fmt.Fprintf(w, "This is a custom HTTP response. Try sending raw packets instead!")
}

// Main function to start the server
func main() {
    // Start TCP server for handling packet data
    go func() {
        fmt.Println("Starting TCP server on 127.0.0.1:65432...")
        listener, err := net.Listen("tcp", "127.0.0.1:65432")
        if err != nil {
            log.Fatal("Error starting TCP server:", err)
        }
        defer listener.Close()

        for {
            conn, err := listener.Accept()
            if err != nil {
                log.Println("Error accepting connection:", err)
                continue
            }

            go handleClient(conn)
        }
    }()

    // Start HTTP server for handling basic HTTP requests (optional)
    fmt.Println("Starting HTTP server on 127.0.0.1:8080...")
    http.HandleFunc("/", handleHttpRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
