package benchmark

import (
	"fmt"
	"net"
	"time"

	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

func ClientTCP(address string, totalDataSize int) error {
	// Connect to TCP server.
	logger.PrintInfo(fmt.Sprintf("Connecting to TCP server at %s...", address))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.PrintInfo("Connected to server. Sending message...")

	numPkgs := totalDataSize / pkgSize
	data := make([]byte, pkgSize)

	// Sends message.
	for i := 0; i < pkgSize; i++ {
		data[i] = 'a'
	}
	start := time.Now()
	for i := 0; i < numPkgs; i++ {
		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}
	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %vs\n", totalDurationSeconds)
	fmt.Printf("Throughput: %v bytes/s\n", float64(totalDataSize)/totalDurationSeconds)

	return nil
}

// ServerTCP starts a TCP server on specified address.
func ServerTCP(addr string) error {
	logger.PrintInfo(fmt.Sprintf("Starting TCP server at %s.", addr))

	// Listen on specified address.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	logger.PrintInfo("Server started.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.PrintError(err)
			continue
		}
		logger.PrintInfo("Connection established.")

		// Create a goroutine to handle client concurrently.
		go handleConnection(conn)
	}
}

// HandleConnection handles communication with client.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read clients message.
	for {
		buffer := make([]byte, pkgSize)
		_, err := conn.Read(buffer)
		if err != nil {
			//logger.PrintInfo(err.Error())
			return
		}
	}
}
