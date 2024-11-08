package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var tcpServerCmd = &cobra.Command{
	Use:   "server [address]",
	Short: "Run TCP server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		tcpServer(addr)
	},
}

// TcpServer starts a TCP server on specified address.
func tcpServer(addr string) error {
	fmt.Printf("Starting TCP server at %s...\n", addr)

	// Listen on specified address.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		fmt.Println("Connection established. Handling client...")

		// Create a goroutine to handle client concurrently.
		go handleConnection(conn)
	}
}

// HandleConnection handles communication with client.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read clients message.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from connection: %v\n", err)
		return
	}
	fmt.Printf("Received from client: %s\n", string(buffer[:n]))

	// Send response.
	message := []byte("Hello from TCP server")
	_, err = conn.Write(message)
	if err != nil {
		log.Printf("Error writing to connection: %v\n", err)
		return
	}
	fmt.Println("Response sent to client.")
}
