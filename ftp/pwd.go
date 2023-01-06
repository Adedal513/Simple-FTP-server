package ftp

func (c *Conn) pwd(args []string) {
	if len(args) != 0 {
		c.respond(status501)
		return
	}

	c.respond(c.workDir)
}
