package main

import (
	"KWeb/framework/gin"
	"KWeb/framework/middleware"
)

// 注册路由规则
func registerRouter(core *gin.Engine) {

	// 静态路由+HTTP方法匹配 为单个路由增加中间件
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		// 在 group 中使用 middleware.Test3() 为单个路由增加中间件
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
