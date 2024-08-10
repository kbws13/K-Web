package framework

import (
	"log"
	"net/http"
)

// Core 框架核心结构
type Core struct {
	router map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// ServerHTTP 框架核心结构实现 Handler 接口
func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, writer)

	// 一个简单的路由选择器，这里写死为测试路由 foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	err := router(ctx)
	if err != nil {
		return
	}
}
