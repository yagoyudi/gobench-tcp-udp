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
	SequenceID int
	Data       []byte
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
	logger.PrintInfo("Connected to server")

	numPackets := totalDataSize / packetSize
	data := make([]byte, packetSize)
	for i := 0; i < packetSize; i++ {
		data[i] = 'a'
	}

	// Sends message to server.
	start := time.Now()
	for i := 0; i < numPackets; i++ {
		packet := PacketUDP{
			SequenceID: i,
			Data:       data,
		}

		packetBytes, err := packet.Serialize()
		if err != nil {
			return err
		}

		//timeoutDuration := 2 * time.Second
		//conn.SetDeadline(time.Now().Add(timeoutDuration))

		for {
			_, err = conn.Write(packetBytes)
			if err != nil {
				return err
			}

			ack := make([]byte, packetSize)
			_, err := conn.Read(ack)
			if err != nil {
				return err
			}

			ackPacket, err := Deserialize(ack)
			if err != nil {
				return err
			}

			if ackPacket.SequenceID == i {
				break
			}
		}
	}
	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %v seconds \n", totalDurationSeconds)

	return nil
}

func ServerUDP(address string) error {
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

	buf := make([]byte, packetSize)
	for {
		_, client, err := conn.ReadFrom(buf)
		if err != nil {
			logger.PrintError(err)
			continue
		}

		packet, err := Deserialize(buf)
		if err != nil {
			return err
		}

		ackPacket := PacketUDP{SequenceID: packet.SequenceID}
		ackBytes, err := ackPacket.Serialize()
		if err != nil {
			return err
		}

		_, err = conn.WriteTo(ackBytes, client)
		if err != nil {
			logger.PrintError(err)
		}
	}

	return nil
}
