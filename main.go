package main

import (
	"KWeb/framework"
	"KWeb/framework/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	core.Use(middleware.Timeout(1 * time.Second))
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
