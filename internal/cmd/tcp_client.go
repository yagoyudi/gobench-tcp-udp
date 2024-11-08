package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var tcpClientCmd = &cobra.Command{
	Use:   "client [address]",
	Short: "Run TCP client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := tcpClient(addr)
		if err != nil {
			log.Println(err)
		}
	},
}

func tcpClient(address string) error {
	// Connect to TCP server.
	fmt.Printf("Connecting to TCP server at %s...\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Connected to server. Sending message...")

	// Sends message.
	message := []byte("Hello from TCP client")
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
