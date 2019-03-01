package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() { // server.
	if len(os.Args) != 2 {
		log.Fatal("Need port number to listen to")
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleconn(conn)
	}
}

func handleconn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
