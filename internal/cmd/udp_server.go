package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

var udpServerCmd = &cobra.Command{
	Use:   "server [address]",
	Short: "Run UDP server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		err := udpServer(addr)
		if err != nil {
			logger.FatalError(err)
		}
	},
}

func udpServer(address string) error {
	logger.PrintInfo(fmt.Sprintf("Starting UDP server on %s...", address))
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.PrintInfo("Server is listening for incoming messages...")

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			logger.PrintError(err)
			continue
		}
		logger.PrintInfo(fmt.Sprintf("Received message: %s", string(buffer[:n])))
		logger.PrintInfo(fmt.Sprintf("Message from %s", clientAddr))

		response := []byte("Message received by UDP server")
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			logger.PrintError(err)
			continue
		}
		logger.PrintInfo("Response sent to client.")
	}
	return nil
}
