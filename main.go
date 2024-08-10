package main

import (
	"KWeb/framework"
	"log"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// 自定义的请求处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8888",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
