package main

import (
	"fmt"
	"net"
)

// global con pool
var ConnMap map[string]net.Conn

func checkErr(err error) int {
	if err != nil {
		//fmt.Println("Err:" + err.Error())
		return 1
	}
	return 0
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		total, err := conn.Read(data)
		if err != nil {
			fmt.Println(string(data[:total]), err)
		} else {
			fmt.Println(string(data[:total]))
		}
		// check
		flag := checkErr(err)
		if flag != 0 {
			delete(ConnMap, conn.RemoteAddr().String())
			break
		}
		// broadcast
		for _, con := range ConnMap {
			if con.RemoteAddr().String() == conn.RemoteAddr().String() {
				continue
			}
			con.Write(data[:total])
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
