package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
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
	logger.PrintInfo(fmt.Sprintf("Starting TCP server at %s...", addr))

	// Listen on specified address.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	logger.PrintInfo("Server started. Waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.PrintError(err)
			continue
		}
		logger.PrintInfo("Connection established. Handling client...")

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
		logger.PrintError(err)
		return
	}
	logger.PrintInfo(fmt.Sprintf("Received from client: %s", string(buffer[:n])))

	// Send response.
	message := []byte("Hello from TCP server")
	_, err = conn.Write(message)
	if err != nil {
		logger.PrintError(err)
		return
	}
	logger.PrintInfo("Response sent to client.")
}
