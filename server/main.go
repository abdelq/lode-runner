package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	addr := flag.String("addr", ":1337", "listener address")
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go newClient(conn)
	}
}
