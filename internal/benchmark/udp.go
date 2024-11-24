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
	log.Println("Connected to server.")

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
	fmt.Printf("Bytes sent: %v bytes/s\n", totalData)

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
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		count += n
		fmt.Printf("Bytes recieved: %d\n", count)
	}

	return nil
}
