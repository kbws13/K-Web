package framework

import "net/http"

// Core 框架核心结构
type Core struct {
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

// ServerHTTP 框架核心结构实现 Handler 接口
func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// TODO
}
