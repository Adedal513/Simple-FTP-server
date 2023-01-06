package ftp

import (
	"fmt"
	"log"
)

type dataPort struct {
	h1, h2, h3, h4 int // host
	p1, p2         int // port
}

func dataPortFromHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort

	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d",
		&dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &dp, nil
}

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}

	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
}

func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}
	dataPort, err := dataPortFromHostPort(args[0])
	if err != nil {
		log.Print(err)
		c.respond(status501)
		return
	}
	c.dataPort = dataPort
	c.respond(status200)
}
