package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
)

// echo the dara recieved on frontend
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	log.Println("Serving")
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
