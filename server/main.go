package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	// Command-line flags
	addr := flag.String("addr", ":1337", "listener's network address")
	flag.Parse()

	// Listen on TCP
	ln, err := net.Listen("tcp", *addr)
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
		log.Printf("New connection from %s", conn.RemoteAddr())

		go newClient(conn)
	}
}
