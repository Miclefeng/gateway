package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 1、创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, //探活时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
	// 2、创建客户端
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	// 3、请求服务器获取数据
	resp, err := client.Get("http://127.0.0.1:8099/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 4、读入内容
	reader, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(nil)
	}
	fmt.Println(string(reader))
}
