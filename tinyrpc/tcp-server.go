package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/kaiya/goutils/tinyrpc/rpc"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()
	bufconn := bufio.NewReader(conn)
	msg, err := rpc.DecodePacket(bufconn)
	if err != nil {
		fmt.Printf("decode packet error:%s\n", err)
		return
	}

	fmt.Println("got msg len: ", len(msg))
	fmt.Println("msg: ", string(msg))
}
