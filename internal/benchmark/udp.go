package benchmark

import (
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
	log.Printf("Total duration: %vs\n", totalDurationSeconds)
	log.Printf("Sent: %v bytes\n", totalData)

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

	var start time.Time
	for {
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, _ := err.(net.Error); netErr.Timeout() {
				if count != 0 {
					log.Printf("Received: %d bytes\n", count)
					log.Printf("Total duration: %vs\n", time.Since(start).Seconds()-1)
				}
				start = time.Now()
				count = 0
			} else {
				log.Println(err)
			}
			continue
		}

		// Only starts to count when first packet arrives.
		// Every packet has size 1024.
		if count == 1024 {
			start = time.Now()
		}

		count += n
	}

	return nil
}
