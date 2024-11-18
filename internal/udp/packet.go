package udp

import (
	"bytes"
	"encoding/binary"
)

const (
	TypeAck uint8 = iota
	TypeData
	TypeEnd
)

const (
	MaxDataLen = 1024
	PacketSize = 1 + 4 + MaxDataLen*1 // Size of struct Packet in bytes.
)

type Packet struct {
	Type uint8
	Seq  uint32
	Data [MaxDataLen]uint8
}

func (p *Packet) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, p)
	return buf.Bytes(), err
}

func Deserialize(data []byte) (*Packet, error) {
	var p Packet
	buf := bytes.NewBuffer(data)
	err := binary.Read(buf, binary.BigEndian, &p)
	return &p, err
}
