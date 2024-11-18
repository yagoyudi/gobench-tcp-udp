package udp

import (
	"net"
)

type Server struct {
	conn *net.UDPConn
}

func NewServer(address string) (*Server, error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{conn: conn}, nil
}

func (s *Server) Send(packet *Packet, clientAddr *net.UDPAddr) error {
	data, err := packet.Serialize()
	if err != nil {
		return err
	}
	_, err = s.conn.WriteToUDP(data, clientAddr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Recv() (*Packet, *net.UDPAddr, error) {
	buf := make([]byte, PacketSize)
	n, clientAddr, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, nil, err
	}

	packet, err := Deserialize(buf[:n])
	if err != nil {
		return nil, nil, err
	}

	return packet, clientAddr, nil
}

func (s *Server) Close() {
	s.conn.Close()
}
