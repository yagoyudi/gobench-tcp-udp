package benchmark

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ClientUDP(addr string, totalData int) error {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Println("Starting to send data")

	totalPackets := totalData / 1024

	data := make([]byte, 1024)
	for i := range data {
		data[i] = 'a'
	}

	start := time.Now()

	for i := 0; i < totalPackets; i++ {
		conn.Write(data)
	}

	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %vs\n", totalDurationSeconds)
	fmt.Printf("Sent: %v bytes\n", totalData)

	return nil
}

func ServerUDP(address string) error {
	log.Printf("Starting UDP server on %s\n", address)

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Println("Server is listening for incoming messages")

	buf := make([]byte, 1024)
	count := 0

	// Read first message.
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	count += n
	// Start timer.
	start := time.Now()
	for {
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Printf("Received: %d bytes\n", count)
				log.Printf("Total duration: %vs\n", time.Since(start).Seconds()-1)
				return nil
			} else {
				return err
			}
		}
		count += n
	}

	return nil
}
