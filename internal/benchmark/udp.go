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
	base := 0
	nextSeq := 0

	packets := make([]udp.Packet, totalPackets)
	for i := 0; i < totalPackets; i++ {
		var data [udp.MaxDataLen]uint8
		for j := 0; j < udp.MaxDataLen; j++ {
			data[j] = 'a'
		}
		packets[i] = udp.Packet{
			Type: udp.TypeData,
			Seq:  uint32(i),
			Data: data,
		}
	}

	start := time.Now()
	for base < totalPackets {
		// Envia pacotes na janela
		for nextSeq < base+udp.WindowLen && nextSeq < totalPackets {
			err := client.Send(&packets[nextSeq])
			if err != nil {
				log.Println(err)
			}
			//log.Printf("Sent pkg %d\n", nextSeq)
			nextSeq++
		}

		// Tenta receber um ACK
		packet, err := client.Recv()
		if err != nil {
			if opErr, ok := err.(net.Error); ok && opErr.Timeout() {
				log.Println("Timeout occurred, retransmitting window")
				// Retransmite pacotes na janela
				for i := base; i < nextSeq; i++ {
					err := client.Send(&packets[i])
					if err != nil {
						log.Printf("Failed to resend packet %d: %v\n", i, err)
					}
				}
				continue
			} else {
				log.Printf("Error receiving packet: %v\n", err)
				continue
			}
		}

		if packet.Type != udp.TypeAck {
			log.Printf("Unexpected packet type: %d\n", packet.Type)
			continue
		}

		//log.Printf("Received ACK %d\n", packet.Seq)
		if packet.Seq >= uint32(base) {
			base = int(packet.Seq + 1)
		}
	}

	endPacket := udp.Packet{
		Type: udp.TypeEnd,
	}
	err = client.Send(&endPacket)
	if err != nil {
		return err
	}
	log.Println("Sent end pkg")

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

		log.Printf("Received packet %d\n", packet.Seq)

		var ackSeq uint32
		if packet.Seq == expectedSeq {
			expectedSeq++
			ackSeq = packet.Seq
		} else {
			// Ignore unordered packets.
			continue
		}

		ackPacket := udp.Packet{
			Type: udp.TypeAck,
			Seq:  ackSeq,
		}
		err = server.Send(&ackPacket, clientAddr)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Sent ACK %d\n", ackSeq)
	}

	return nil
}
