package http

import (
	"KWeb/app/http/module/demo"
	"KWeb/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
