package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadTestRouter(r *gin.Engine) {
	r.GET("/hello", TestHandler)

}

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
