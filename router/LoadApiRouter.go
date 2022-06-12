package router

import (
	"bookstore/controller"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	api := r.Group("/api/v1")

	// user
	api.POST("/register", controller.RegisterHandler)
	api.POST("/login", controller.LoginHandler)

	// book
	api.POST("/book", controller.CreateBookHandler)
	api.GET("/book", controller.GetAllBookHandler)
	api.GET("/book/:id", controller.GetBookDetailHandler)
	api.PUT("/book/", controller.UpdateBookDetailHandler)
	api.DELETE("/book/:id", controller.DeleteBookHandler)
}
