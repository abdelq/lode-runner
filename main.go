package main

import (
	"flag"
	"log"
	"net"

	"github.com/abdelq/lode-runner/web"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	// Command-line flags
	tcpAddr := flag.String("tcp", ":1337", "TCP network address")
	flag.Parse()

	// Listen on TCP
	ln, err := net.Listen("tcp", *tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s %s", ln.Addr().Network(), ln.Addr())

	// Listen on HTTP
	go web.Listen()

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
