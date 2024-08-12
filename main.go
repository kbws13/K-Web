package main

import (
	"KWeb/app/provider/demo"
	"KWeb/framework/gin"
	"KWeb/framework/middleware"
	"KWeb/framework/provider/app"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	// 绑定具体的服务
	core.Bind(&app.KAppProvider{})
	core.Bind(&demo.DemoProvider{})

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		// 自定义的请求处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8888",
	}
	go func() {
		server.ListenAndServe()
	}()
	// 当前的 goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 goroutine 等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
