package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn, i int) {
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)

	if err != nil {
		log.Fatal(err)
	}

	str := fmt.Sprintf("HTTP/1.1 200 OK \r\n\r\n Hello, connection %d\r\n", i)
	time.Sleep(2 * time.Second)

	conn.Write([]byte(str))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7699")

	if err != nil {
		log.Fatal(err)
	}

	i := 1

	for {
		fmt.Printf("Waiting for a client\n")
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Client %d connected\n", i)
		i += 1
		go handle(conn, i)

	}
}
