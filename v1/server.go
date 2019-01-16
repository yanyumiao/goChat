package main

import "fmt"
import "net"

var ConnMap map[string]net.Conn
var messages = make(chan string)

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
			messages <- string(data[:n])
		}
	}
}

func broadcaster() {
	for {
		msg := <-messages
		fmt.Println(msg)
		for _, con := range ConnMap {
			con.Write([]byte(msg))
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
