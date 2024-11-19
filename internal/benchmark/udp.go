package benchmark

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/yagoyudi/gobench-tcp-udp/internal/udp"
)

func ClientUDP(addr string, totalData int) error {
	client, err := udp.NewClient(addr)
	if err != nil {
		return err
	}
	defer client.Close()
	log.Println("Connected to server.")

	totalPackets := totalData / udp.MaxDataLen
	base, nextSeq := 0, 0

	data := [udp.MaxDataLen]uint8{}
	for i := range data {
		data[i] = 'a'
	}

	packets := make([]udp.Packet, totalPackets)
	for i := 0; i < totalPackets; i++ {
		packets[i] = udp.Packet{
			Type: udp.TypeData,
			Seq:  uint32(i),
			Data: data,
		}
	}

	start := time.Now()

	ackCh := make(chan uint32, udp.WindowLen)
	defer close(ackCh)

	go func() {
		for base < totalPackets {
			for nextSeq < base+udp.WindowLen && nextSeq < totalPackets {
				_ = client.Send(&packets[nextSeq])
				//log.Printf("Sent pkt %d\n", nextSeq)
				nextSeq++
			}
		}
	}()

	for base < totalPackets {
		// Tenta receber um ACK
		packet, err := client.Recv()
		if err != nil {
			if opErr, ok := err.(net.Error); ok && opErr.Timeout() {
				for i := base; i < nextSeq; i++ {
					_ = client.Send(&packets[i])
				}
				continue
			}
			log.Printf("Error receiving packet: %v\n", err)
			continue
		}

		if packet.Type == udp.TypeAck && packet.Seq >= uint32(base) {
			//log.Printf("Received ACK %d\n", packet.Seq)
			base = int(packet.Seq + 1)
		}
	}

	endPacket := udp.Packet{Type: udp.TypeEnd}
	_ = client.Send(&endPacket)
	//log.Println("Sent end pkg")

	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %vs\n", totalDurationSeconds)
	fmt.Printf("Throughput: %v bytes/s\n", float64(totalData)/totalDurationSeconds)

	return nil
}

func ServerUDP(addr string) error {
	log.Printf("Starting UDP server on %s\n", addr)

	server, err := udp.NewServer(addr)
	if err != nil {
		return err
	}
	defer server.Close()
	log.Println("Server is listening for incoming messages")

	expectedSeq := uint32(0)
	for {
		packet, clientAddr, err := server.Recv()
		if err != nil {
			log.Println(err)
			continue
		}

		if packet.Type == udp.TypeEnd {
			expectedSeq = 0
			log.Println("Received type END")
			continue
		}

		//log.Printf("Received packet %d\n", packet.Seq)

		if packet.Seq == expectedSeq {
			expectedSeq++
		}

		ackPacket := udp.Packet{
			Type: udp.TypeAck,
			Seq:  packet.Seq,
		}
		_ = server.Send(&ackPacket, clientAddr)
		//log.Printf("Sent ACK %d\n", packet.Seq)
	}

	return nil
}
