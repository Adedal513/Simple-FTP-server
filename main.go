package main

import (
	"crypto/rand"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"gopl.io/ch8/ftpserver/ftp"
)

// TODO:
// 	- AUTH
// 		Possible auth with password and username
//	- UPLOAD TO SEVER
// 	- ROOT PROTECTION

var port int
var rootDir string
var certDir string = "./cert"
var tlsConfig tls.Config

func init() {
	certPath := filepath.Join(certDir, "server.pem")
	keyPath := filepath.Join(certDir, "server.key")

	if _, err := os.Stat(certPath); err != nil {
		log.Fatal("Unable to locate certificate")
	}

	if _, err := os.Stat(keyPath); err != nil {
		log.Fatal("Unable to locate private key")
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)

	if err != nil {
		log.Fatal(err)
	}

	tlsConfig = tls.Config{Certificates: []tls.Certificate{cert}}
	tlsConfig.Rand = rand.Reader

	flag.IntVar(&port, "port", 8080, "port number")
	flag.StringVar(&rootDir, "rootDir", "./public", "root directory")
	flag.Parse()
}

func main() {
	server := fmt.Sprintf(":%d", port)
	listener, err := tls.Listen("tcp", server, &tlsConfig)

	log.Print(fmt.Sprintf("Started listening on port %d", port))

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Print(fmt.Sprintf("Accepted connection from %s", conn.RemoteAddr()))
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer CloseConnection(c)
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}

func CloseConnection(c net.Conn) {
	log.Printf("Connection closed for %s", c.RemoteAddr())
	c.Close()
}
