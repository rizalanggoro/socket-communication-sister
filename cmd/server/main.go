package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Message struct {
	sender net.Conn
	text   string
}

var (
	clients   = make(map[net.Conn]bool)
	mutex     sync.Mutex
	broadcast = make(chan Message)
)

func main() {
	// listen tcp pada port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	fmt.Println("tcp server running on :8080")

	go handleBroadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		mutex.Lock()
		clients[conn] = true
		mutex.Unlock()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	reader := bufio.NewScanner(conn)
	for reader.Scan() {
		msg := reader.Text()
		fmt.Println("received:", msg)
		// kirim pesan beserta info pengirim
		broadcast <- Message{sender: conn, text: msg}
	}
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for conn := range clients {
			if conn == msg.sender {
				continue // skip pengirim sendiri
			}
			fmt.Fprintln(conn, msg.text)
		}
		mutex.Unlock()
	}
}
