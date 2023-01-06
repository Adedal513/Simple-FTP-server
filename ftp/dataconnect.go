package ftp

import "net"

func (c *Conn) dataConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.dataPort.toAddress())

	if err != nil {
		return nil, err
	}

	return conn, nil
}
