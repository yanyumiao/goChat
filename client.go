package main

import "fmt"
import "net"
import "os"
import "bufio"

var ch chan int = make(chan int)
var nickname string

func read(conn net.Conn) {
	buff := make([]byte, 256)
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
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	go read(conn)
	fmt.Println("Input nickname:")
	fmt.Scanln(&nickname)
	fmt.Println("Your nickname:", nickname)
	var msg []byte
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, _, _ = reader.ReadLine()
		msg := "[" + nickname + "]:" + string(msg)
		conn.Write([]byte(msg))
		select {
		case <-ch:
			fmt.Println("Server error")
			os.Exit(2)
		default:
		}
	}
}
