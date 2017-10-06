package main

import (
	"crypto/tls"
	"flag"
	"log"
)

func main() {
	// Flags
	addr := flag.String("addr", ":443", "listener's network address")
	flag.Parse()
	log.SetFlags(log.Ltime | log.Lshortfile)

	// Load public/private key pair
	crt, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	conf := tls.Config{Certificates: []tls.Certificate{crt}}

	// Listen on TCP
	ln, err := tls.Listen("tcp", *addr, &conf)
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
