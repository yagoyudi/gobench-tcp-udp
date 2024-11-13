package benchmark

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"github.com/yagoyudi/gobench-tcp-udp/internal/logger"
)

type PacketUDP struct {
	Type string
	Data []byte
}

func (p *PacketUDP) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(p)
	return buf.Bytes(), err
}

func Deserialize(data []byte) (*PacketUDP, error) {
	var p PacketUDP
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(&p)
	return &p, err
}

func ClientUDP(addr string, totalDataSize int) error {
	// Connect to server.
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.PrintInfo("Connected to server.")

	// Sends message to server.
	numPackets := totalDataSize / packetSize
	msg := make([]byte, packetSize)
	for i := 0; i < packetSize; i++ {
		msg[i] = 'a'
	}
	start := time.Now()
	for i := 0; i < numPackets; i++ {
		//timeoutDuration := 2 * time.Second
		//conn.SetDeadline(time.Now().Add(timeoutDuration))

		for {
			_, err = conn.Write(msg)
			if err != nil {
				return err
			}

			ack := make([]byte, 3)
			_, err := conn.Read(ack)
			if err != nil {
				return err
			}

			if string(ack) == "ack" {
				break
			}
		}
	}
	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %vs\n", totalDurationSeconds)
	fmt.Printf("Throughput: %v bytes/s\n", float64(totalDataSize)/totalDurationSeconds)

	return nil
}

func ServerUDP(address string) error {
	logger.PrintInfo(fmt.Sprintf("Starting UDP server on %s.", address))

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	logger.PrintInfo("Server is listening for incoming messages.")

	for {
		buf := make([]byte, packetSize)
		_, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			logger.PrintError(err)
			continue
		}

		var ack []byte
		ack = []byte("ack")

		_, err = conn.WriteToUDP(ack, clientAddr)
		if err != nil {
			logger.PrintError(err)
		}
	}

	return nil
}
