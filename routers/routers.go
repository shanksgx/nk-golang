package routers

import (
	"github.com/gin-gonic/gin"
	"nk-golang/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)
	// v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/post", controller.CreateHander)
	}
	return r
}
