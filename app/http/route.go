package http

import (
	"KWeb/app/http/module/demo"
	"KWeb/framework/gin"
	"KWeb/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	// 先去 ./dist 目录下查找文件是否存在，找到就使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	demo.Register(r)
}
