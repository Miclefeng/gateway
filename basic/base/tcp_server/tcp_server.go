package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、监听端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	// 2、建立套接字
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		go func(accept net.Conn) {
			defer accept.Close()
			// 3、读取数据
			for {
				var buf [1024]byte
				readNum, err := accept.Read(buf[:])
				if err != nil {
					fmt.Printf("read from connect failed, err: %v\n", err)
					break
				}
				str := string(buf[:readNum])
				fmt.Printf("receive from client, data: %v\n", str)
			}

		}(accept)
	}
}
