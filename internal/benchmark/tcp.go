package benchmark

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ClientTCP(address string, totalData int) error {
	// Connect to TCP server.
	log.Printf("Connecting to TCP server at %s", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Println("Connected to server")

	totalPackets := totalData / 1024
	data := make([]byte, 1024)

	// Sends message.
	for i := 0; i < 1024; i++ {
		data[i] = 'a'
	}
	start := time.Now()
	for i := 0; i < totalPackets; i++ {
		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}
	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %vs\n", totalDurationSeconds)
	fmt.Printf("Sent: %v bytes\n", 1024*totalPackets)

	return nil
}

// ServerTCP starts a TCP server on specified address.
func ServerTCP(addr string) error {
	log.Printf("Starting TCP server at %s", addr)

	// Listen on specified address.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept: %s\n", err.Error())
			continue
		}
		log.Println("Connection established")

		// Create a goroutine to handle client concurrently.
		go handleConnection(conn)
	}
}

// HandleConnection handles communication with client.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read clients message.
	count := 0
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Read: %s\n", err.Error())
			return
		}
		count += n
		fmt.Printf("Received: %d bytes\n", count)
	}
}
