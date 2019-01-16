package main

import "fmt"
import "net"

var ConnMap map[string]net.Conn

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
			fmt.Println(string(data[:n]))
		}
		// broadcast
		for _, con := range ConnMap {
			if con.RemoteAddr().String() == conn.RemoteAddr().String() {
				continue
			}
			con.Write(data[:n])
		}
	}
}

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:9999")
	ConnMap = make(map[string]net.Conn)
	for {
		conn, _ := listen.Accept()
		defer conn.Close()
		ConnMap[conn.RemoteAddr().String()] = conn
		fmt.Println("New client:", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
