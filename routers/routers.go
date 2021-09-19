package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 默认路由
	r := gin.Default()

	// 加载模版文件
	r.LoadHTMLGlob("templates/*")

	// 加载静态文件
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler())

	// 路由组v1
	v1 := r.Group("/v1")
	{
		v1.GET("/todo/:id", controller.GetToDo())
		v1.POST("/todo", controller.CreateTodo())
		v1.GET("/todo", controller.GetAllTodo())
		v1.PUT("/todo/:id", controller.UpdateTodo())
		v1.DELETE("/todo/:id", controller.DeleteTodo())
	}
	return r
}
