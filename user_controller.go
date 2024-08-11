package main

import (
	"KWeb/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo := c.DefaultQuery("foo", "def")
	// 等待10s结束
	time.Sleep(10 * time.Second)
	// 输出结果
	c.JSON(200, "ok, UserLoginController: "+foo)
}
