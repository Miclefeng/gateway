package main

import (
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":8099"
)

func main() {
	// 1、创建路由
	mux := http.NewServeMux()
	// 2、设置路由规则
	mux.HandleFunc("/hello", sayBye)
	// 3、配置 http 服务器
	server := &http.Server{
		Addr:         Addr,
		Handler:      mux,
		WriteTimeout: 3 * time.Second,
	}
	// 4、监听端口并启动服务
	log.Println("Starting http_server at " + Addr)
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("bye, this is a http_server."))
}
