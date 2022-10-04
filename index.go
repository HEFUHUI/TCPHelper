package tcphelper

import (
	"net"
)

func Run(port int) {
	// port 优雅转 string
	portStr := ":" + string(rune(port))
	// 监听端口
	if connect, err := net.Listen("tcp", portStr); err != nil {
		panic(err)
	} else {
		for {
			// connect.Accept()
			if conn, err := connect.Accept(); err != nil {
				panic(err)
			} else {
				go handleConnect(conn)
			}
		}
	}
}

func handleConnect(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	state := false
	// 读取数据
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		panic(err)
	}
	// 判断数据
	if string(data[:n]) == "hello" {
		state = true
	}
	// 返回数据
	if state {
		conn.Write([]byte("world"))
	} else {
		conn.Write([]byte("bye"))
	}
}
