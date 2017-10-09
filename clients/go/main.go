package main

import (
	"crypto/tls"
	"flag"
	"log"
)

func main() {
	// Command-line flags
	addr := flag.String("addr", ":443", "network address")
	flag.Parse()

	// Dial on TCP
	conn, err := tls.Dial("tcp", *addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to %s", conn.RemoteAddr())

	defer conn.Close()
	// TODO Join a room
}
