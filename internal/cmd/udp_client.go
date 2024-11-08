package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

var udpClientCmd = &cobra.Command{
	Use:   "client [address]",
	Short: "Run UDP client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := udpClient(addr)
		if err != nil {
			logger.FatalError(err)
		}
	},
}

func udpClient(addr string) error {
	// Connect to server
	logger.PrintInfo(fmt.Sprintf("Connecting to UDP server at %s...", addr))
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.PrintInfo("Connected to server. Sending message...")

	// Sends message to server.
	message := []byte("Hello from UDP client")
	_, err = conn.Write(message)
	if err != nil {
		return err
	}
	logger.PrintInfo("Message sent to server. Waiting for response...")

	// Wait response.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	logger.PrintInfo(fmt.Sprintf("Received from server: %s", string(buffer[:n])))
	return nil
}
