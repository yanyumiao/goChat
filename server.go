// server
package main

import (
	"fmt"
	"net"
)

// 链接池 声明 全局
var ConnMap map[string]*net.TCPConn

// 检查错误
func checkErr(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("用户退出")
			return 0
		}
		fmt.Println("发生错误")
		return -1
	}
	return 1
}

// 业务
func say(tcpConn *net.TCPConn) {
	for {
		data := make([]byte, 256) // 256???
		// https://golang.org/pkg/net/#Buffers.Read
		// func (v *Buffers) Read(p []byte) (n int, err error)
		total, err := tcpConn.Read(data)
		if err != nil {
			fmt.Println(string(data[:total]), err) // 类型转换
			// TODO err in for
		} else {
			fmt.Println(string(data[:total]))
		}
		// 错误检查
		flag := checkErr(err)
		if flag == 0 {
			break
		}
		// 广播收到的消息
		for _, conn := range ConnMap {
			// 消息不发送给自己
			if conn.RemoteAddr().String() == tcpConn.RemoteAddr().String() {
				continue
			}
			conn.Write(data[:total])
		}
	}
}
func main() {
	//var conn net.TCPConn
	//localAddr :=conn.LocalAddr().String()
	//fmt.Println(localAddr)
	//tcpAddr, _ := net.ResolveTCPAddr("tcp",localAddr)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9000") // 注意端口选择
	tcpListen, _ := net.ListenTCP("tcp", tcpAddr)
	ConnMap = make(map[string]*net.TCPConn) // 创建
	for {
		tcpConn, _ := tcpListen.AcceptTCP()
		defer tcpConn.Close()
		ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
		fmt.Println("连接客户端信息：", tcpConn.RemoteAddr().String())
		go say(tcpConn) // 支持并发
	}
}
