package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、创建 listen 监听
	listen, err := net.Listen("tcp", ":9090")

	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	for  {
		// 2、建立连接套接字
		accept, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		//3. 创建处理协程(接收数据)
		go func(accept net.Conn) {
			// defer accept.Close() // 出现 close_wait | FIN_WAIT_2 状态
			for {
				var buf [1024]byte
				readNum, err := accept.Read(buf[:])
				if err != nil {
					fmt.Printf("read from connect failed, err: %v\n", err)
					break
				}
				fmt.Printf("receive from client, data: %v\n", string(buf[:readNum]))
			}
		}(accept)
	}
}
