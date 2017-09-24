package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conn, err := tls.Dial("tcp", ":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Read(make([]byte, 100))
	if err != nil {
		log.Println(err)
		return
	}
}
