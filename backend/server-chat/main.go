package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"tcp-chat/backend/common"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "backend"
)

var (
	connections []net.Conn
)

func main() {
	// listen incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	//close the listener when the application closes
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		//Save connection
		connections = append(connections, conn)
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for {
		msg, err := common.ReadMsg(conn)
		if err != nil {
			if err == io.EOF {
				// close the connection when you're done with it
				removeConn(conn)
				conn.Close()
				return
			}
			log.Println("Error reading:", err.Error())
			return
		}
		fmt.Printf("Message Recieved: %s\n", msg)
		broadcast(conn, msg)
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = 0; i < len(connections); i++ {
		if connections[i] == conn {
			break
		}
	}
	connections = append(connections[:i], connections[i+1:]...)
}

func broadcast(conn net.Conn, msg string) {
	for i := range connections {
		if connections[i] != conn {
			err := common.WriteMsg(connections[i], msg)
			if err != nil {
				log.Println("Error broadcasting:", err.Error())
			}
		}
	}
}
