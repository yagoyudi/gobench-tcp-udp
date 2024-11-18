package benchmark

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	TypeAck uint8 = iota
	TypeData
	TypeEnd
)

const (
	dataSize = 1024
	pkgSize  = 1 + 4 + dataSize // Size of struct Package.
	winSize  = 5                // Length of window
	timeout  = 2 * time.Second
)

type Package struct {
	Type uint8
	Seq  uint32
	Data [dataSize]uint8
}

func (p *Package) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, p)
	return buf.Bytes(), err
}

func Deserialize(data []byte) (*Package, error) {
	var p Package
	buf := bytes.NewBuffer(data)
	err := binary.Read(buf, binary.BigEndian, &p)
	return &p, err
}

func ClientUDP(addr string, totalDataSize int) error {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Println("Connected to server.")

	numPkgs := totalDataSize / dataSize
	base := 0
	nextSeq := 0

	var pkgs []Package
	for i := 0; i < numPkgs; i++ {
		var msg [dataSize]uint8
		for j := 0; j < dataSize; j++ {
			msg[j] = 'a'
		}
		pkg := Package{
			Type: TypeData,
			Seq:  uint32(i),
			Data: msg,
		}
		pkgs = append(pkgs, pkg)
	}

	done := make(chan bool)

	start := time.Now()

	// Gerenciar retransmissões.
	go func() {
		for {
			select {
			case <-done:
				return
			case <-time.After(timeout):
				fmt.Println("Timeout occurred, retransmitting window")
				for i := base; i < nextSeq; i++ {
					data, err := pkgs[i].Serialize()
					if err != nil {
						log.Printf("serialize: %s\n", err.Error())
					}
					_, err = conn.Write(data)
					if err != nil {
						log.Printf("write: %s\n", err.Error())
					}
				}
			}
		}
	}()

	for base < numPkgs {
		// Envia pacotes na janela
		for nextSeq < base+winSize && nextSeq < numPkgs {
			data, _ := pkgs[nextSeq].Serialize()
			_, err := conn.Write(data)
			if err != nil {
				log.Printf("write: %s\n", err.Error())
			}
			log.Printf("Sent pkg %d\n", nextSeq)
			nextSeq++
		}

		// Espera ACK
		conn.SetReadDeadline(time.Now().Add(timeout))
		buf := make([]byte, pkgSize)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("read: %s\n", err.Error())
		}

		pkg, err := Deserialize(buf)
		if err != nil {
			log.Printf("deserialize: %s\n", err.Error())
		}
		if pkg.Type != TypeAck {
			log.Printf("want ack, got %d\n", pkg.Type)
		}

		log.Printf("Received ACK %d\n", pkg.Seq)
		base = int(pkg.Seq + 1)
	}

	endPkg := Package{
		Type: TypeEnd,
		Seq:  uint32(numPkgs),
	}
	data, _ := endPkg.Serialize()
	conn.Write(data)
	fmt.Println("Sent end pkg")

	done <- true

	totalDurationSeconds := time.Since(start).Seconds()
	fmt.Printf("Total duration: %vs\n", totalDurationSeconds)
	fmt.Printf("Throughput: %v bytes/s\n", float64(totalDataSize)/totalDurationSeconds)

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
	log.Println("Server is listening for incoming messages.")

	expectedSeq := uint32(0)
	for {
		buf := make([]byte, pkgSize)
		_, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("ReadFromUDP: %s\n", err.Error())
			continue
		}

		pkg, err := Deserialize(buf)
		if err != nil {
			log.Printf("Deserialize: %s\n", err.Error())
			continue
		}

		if pkg.Type == TypeEnd {
			log.Println("Received type END")
			continue
		}

		log.Printf("Received packet %d\n", pkg.Seq)

		var ackSeq uint32
		if pkg.Seq == expectedSeq {
			expectedSeq++
			ackSeq = pkg.Seq
		} else if pkg.Seq < expectedSeq {
			ackSeq = expectedSeq - 1
		} else {
			// Ignore unordered packets.
			continue
		}

		ack := Package{
			Type: TypeAck,
			Seq:  ackSeq,
		}
		data, err := ack.Serialize()
		if err != nil {
			log.Printf("Serialize: %s\n", err.Error())
			continue
		}

		_, err = conn.WriteToUDP(data, clientAddr)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Sent ACK %d\n", ackSeq)
	}

	return nil
}
