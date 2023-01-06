package ftp

import (
	"bufio"
	"log"
)

const (
	status150 = "150 File status okay; about to open data connection."
	status200 = "200 Command okay."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status230 = "230 User %s logged in, proceed."
	status257 = "257 %s path created."
	status332 = "332 Authentication needed to proceed."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Command not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
)

func (c *Conn) respond(s string) {
	log.Print(">> ", s)

	response := []byte(s)
	buffer := bufio.NewWriter(c.conn)

	for _, b := range response {
		err := buffer.WriteByte(b)

		if err != nil {
			log.Print(err)
		}
	}

	buffer.WriteString(c.EOL())

	err := buffer.Flush()

	if err != nil {
		log.Print(err)
	}
}

func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}
