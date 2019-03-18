package main

import "fmt"
import "net"

var ConnMap map[string]net.Conn

type message struct {
	content string
	from    string
}

var messages = make(chan message)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 256)
		n, err := conn.Read(data)
		if err != nil {
			//fmt.Println(err.Error())
			delete(ConnMap, conn.RemoteAddr().String())
			break
		} else {
			var msg message
			msg.content = string(data[:n])
			msg.from = conn.RemoteAddr().String()
			messages <- msg
		}
	}
}

func broadcaster() {
	for {
		msg := <-messages
		fmt.Println(msg.content)
		for k, con := range ConnMap {
			if k == msg.from {
				continue
			}
			con.Write([]byte(msg.content))
		}
	}
}

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:9999")
	ConnMap = make(map[string]net.Conn)
	go broadcaster()
	for {
		conn, _ := listen.Accept()
		defer conn.Close()
		ConnMap[conn.RemoteAddr().String()] = conn
		fmt.Println("New client:", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
