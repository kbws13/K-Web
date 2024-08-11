package middleware

import (
	"KWeb/framework/gin"
	"fmt"
)

func Test1() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test1")
		// 使用 Next 往下调用，会自增 context.index
		c.Next()
		fmt.Println("middleware post test1")
	}
}

func Test2() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
	}
}

func Test3() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
	}
}
