package ftp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var default_mode = os.FileMode(int(0777))

func (c *Conn) mkd(args []string) {
	if len(args) < 1 {
		c.respond(status501)
		return
	}

	absPath := filepath.Join(c.rootDir, c.workDir)
	newPath := filepath.Join(absPath, args[0])

	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		err := os.Mkdir(newPath, default_mode)

		if err != nil {
			log.Print(err)
			c.respond(status550)
			return
		}

		c.respond(fmt.Sprintf(status257, strings.Join(args, " ")))
	}
}
