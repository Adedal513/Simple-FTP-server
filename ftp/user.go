package ftp

import (
	"fmt"
	"strings"
)

type Authenticator interface {
	Login(string) (string, error)
	Retrive(string) (string, error)
	Register(string) (string, error)
}

func (c *Conn) user(args []string) {
	c.respond(fmt.Sprintf(status230, strings.Join(args, " ")))
}
