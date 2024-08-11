package main

import (
	"KWeb/framework/gin"
	"fmt"
)

func SubjectAddController(c *gin.Context) {
	c.JSON(200, "ok, SubjectAddController")
}

func SubjectListController(c *gin.Context) {
	c.JSON(200, "ok, SubjectListController")
}

func SubjectDelController(c *gin.Context) {
	c.JSON(200, "ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.JSON(200, "ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId := c.Param("id")
	c.JSON(200, "ok, SubjectGetController:"+fmt.Sprint(subjectId))
}

func SubjectNameController(c *gin.Context) {
	c.JSON(200, "ok, SubjectNameController")
}
