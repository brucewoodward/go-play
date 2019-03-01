package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var nconnections int = 0

func main() { // client to multiple servers.
	var ports []int
	// read port numbers from the command line.
	for _, a := range os.Args[1:] {
		v, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}
		ports = append(ports, v)
	}
	for _, port := range ports {
		fmt.Printf("connecting to localhost %d\n", port)
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
		fmt.Printf("%v\n", conn)
		if err != nil {
			fmt.Printf("%v\n", err)
			log.Fatal(err)
			continue
		}
		go handler(conn)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func handler(conn net.Conn) {
	fmt.Println("in handler")
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Print(err)
	}
}
