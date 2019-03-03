package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var nconnections int = 0

type connection struct {
	name               string
	connection_string  string
	network_connection net.Conn
}

func main() { // client to multiple servers.
	var conns []connection
	// read port numbers from the command line.
	for _, a := range os.Args[1:] {
		fields := strings.Split(a, "=")
		conns = append(conns, connection{name: fields[0], connection_string: fields[1]})
	}
	// make all the connection before starting to read form them
	for i, con := range conns {
		fmt.Printf("connecting to %s\n", con.connection_string)
		conn, err := net.Dial("tcp", con.connection_string)
		if err != nil {
			log.Fatal(err)
		}
		conns[i].network_connection = conn
	}
	for _, con := range conns {
		go handler(con.name, con.network_connection)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func handler(name string, conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("%s %s\n", name, scanner.Text())
	}
}
