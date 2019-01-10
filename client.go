package main

import (
	"fmt"
	"net"
	"os"
)

var ch chan int = make(chan int)
var nickname string

func read(conn *net.TCPConn) {
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			ch <- 1
			break
		}
		fmt.Printf("%s\n", buff[0:n])
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	server := os.Args[1]
	TcpAdd, _ := net.ResolveTCPAddr("tcp", server)
	//TcpAdd, _ := net.ResolveTCPAddr("tcp", "localhost:9999")
	conn, err := net.DialTCP("tcp", nil, TcpAdd)
	if err != nil {
		fmt.Println("Server closed.")
		os.Exit(1)
	}
	defer conn.Close()
	go read(conn)
	fmt.Println("Input nickname:")
	fmt.Scanln(&nickname)
	fmt.Println("Your nickname:", nickname)
	for {
		var msg string
		fmt.Scan(&msg)
		//fmt.Print("[" + nickname + "]" + ":")
		//fmt.Println(msg)
		b := []byte("[" + nickname + "]" + ":" + msg)
		conn.Write(b)
		select {
		case <-ch:
			fmt.Println("Server error.")
			os.Exit(2)
		default:
		}
	}
}
