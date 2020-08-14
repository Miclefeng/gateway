package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {

		return
	}
	//defer dial.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		trimString := strings.TrimSpace(readString)
		if trimString == "Q" {
			break
		}
		writeNum, err := dial.Write([]byte(trimString))
		if err != nil {
			break
		}
		fmt.Println("send number: ", writeNum)
	}
}
