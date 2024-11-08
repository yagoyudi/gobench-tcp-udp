package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var udpClientCmd = &cobra.Command{
	Use:   "client [address]",
	Short: "Run UDP client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := udpClient(addr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func udpClient(addr string) error {
	// Connect to server
	fmt.Printf("Connecting to UDP server at %s...\n", addr)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Connected to server. Sending message...")

	// Sends message to server.
	message := []byte("Hello from UDP client")
	_, err = conn.Write(message)
	if err != nil {
		return err
	}
	fmt.Println("Message sent to server. Waiting for response...")

	// Wait response.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	fmt.Printf("Received from server: %s\n", string(buffer[:n]))
	return nil
}
