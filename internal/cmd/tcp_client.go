package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

var tcpClientCmd = &cobra.Command{
	Use:   "client [address]",
	Short: "Run TCP client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := tcpClient(addr)
		if err != nil {
			logger.PrintError(err)
		}
	},
}

func tcpClient(address string) error {
	// Connect to TCP server.
	logger.PrintInfo(fmt.Sprintf("Connecting to TCP server at %s...", address))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.PrintInfo("Connected to server. Sending message...")

	// Sends message.
	message := []byte("Hello from TCP client")
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
