package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"tcp-chat/backend/common"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("backend", CONN_HOST+":"+CONN_PORT)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("backend", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
}

func writeInput(conn *net.TCPConn) {
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	username = username[:len(username)-1]
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter text: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		err = common.WriteMsg(conn, username+": "+text)
		if err != nil {
			log.Println(err)
		}
	}
}

func printOutput(conn *net.TCPConn) {
	for {
		msg, err := common.ReadMsg(conn)
		if err == io.EOF {
			//close connection and exit
			err := conn.Close()
			if err != nil {
				return
			}
			fmt.Println("Connection close. Bye-bye ;)")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}
}
