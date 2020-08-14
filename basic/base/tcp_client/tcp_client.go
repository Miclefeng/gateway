package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1、连接服务器
	dial, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	defer dial.Close()

	// 2、读取命令行输入
	reader := bufio.NewReader(os.Stdin)
	for {
		// 3、一只读取到 \n
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			return
		}

		// 4、到 Q 就退出
		trimString := strings.TrimSpace(readString)
		if trimString == "Q"{
			break
		}

		// 5、将输入传送给tcp_server
		writeNum, err := dial.Write([]byte(trimString))
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}

		fmt.Println(writeNum)
	}
}
