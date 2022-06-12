package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {

	// 1.生成路由
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 2.路由分层
	LoadTestRouter(r)
	LoadApiRouter(r)
	return r
}
