package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	// Command-line flags
	addr := flag.String("addr", ":1337", "network address")
	flag.Parse()

	var ln net.Listener
	var err error

	// Listen on TCP
	if _, port, _ := net.SplitHostPort(*addr); port == "443" {
		crt, err := tls.LoadX509KeyPair("server.crt", "server.key")
		if err != nil {
			log.Fatal(err)
		}
		ln, err = tls.Listen("tcp", *addr, &tls.Config{
			Certificates: []tls.Certificate{crt},
		})
	} else {
		ln, err = net.Listen("tcp", *addr)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s %s", ln.Addr().Network(), ln.Addr())

	defer ln.Close()
	for {
		// Wait for connection
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go newClient(conn)
	}
}
