package http

import (
	"KWeb/app/http/module/demo"
	"KWeb/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
