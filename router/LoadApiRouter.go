package router

import (
	"bookstore/controller"

	"github.com/gin-gonic/gin"
)

// 注册路由、绑定处理函数
func LoadApiRouter(r *gin.Engine) {
	api := r.Group("/api/v1")

	// 用户信息
	api.POST("/register", controller.RegisterHandler)
	api.POST("/login", controller.LoginHandler)

	// 书籍信息
	api.POST("/book", controller.CreateBookHandler)
	api.GET("/book", controller.GetAllBookHandler)
	api.GET("/book/:id", controller.GetBookDetailHandler)
	api.PUT("/book/", controller.UpdateBookDetailHandler)
	api.DELETE("/book/:id", controller.DeleteBookHandler)
}
