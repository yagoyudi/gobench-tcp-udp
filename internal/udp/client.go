package udp

import (
	"net"
	"time"
)

type Client struct {
	conn net.Conn
}

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (c *Client) Send(packet *Packet) error {
	data, err := packet.Serialize()
	if err != nil {
		return err
	}
	_, err = c.conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Recv() (*Packet, error) {
	buf := make([]byte, PacketSize)
	c.conn.SetReadDeadline(time.Now().Add(Timeout)) // Define timeout para a leitura
	n, err := c.conn.Read(buf)
	if err != nil {
		return nil, err
	}

	packet, err := Deserialize(buf[:n])
	if err != nil {
		return nil, err
	}

	return packet, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
