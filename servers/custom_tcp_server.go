package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(connection, "Hello my Gs!\n")
		io.WriteString(connection, "Na so I turn framework engr o\n")
		io.WriteString(connection, "Nice!!\n")

		connection.Close()
	}
}
