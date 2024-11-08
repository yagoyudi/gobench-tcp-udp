package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var udpServerCmd = &cobra.Command{
	Use:   "server [address]",
	Short: "Run UDP server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := udpServer(addr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func udpServer(address string) error {
	fmt.Printf("Starting UDP server on %s...\n", address)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Server is listening for incoming messages...")

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(buffer[:n]))
		fmt.Printf("Message from %s\n", clientAddr)

		response := []byte("Message received by UDP server")
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			log.Printf("Error sending response: %v\n", err)
			continue
		}
		fmt.Println("Response sent to client.")
	}
	return nil
}
