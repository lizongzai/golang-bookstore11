package main

import (
	"bookstore/dao/mysql"
	"bookstore/router"
	"fmt"
)

func main() {

	// 1.连接数据库
	mysql.InitMysql()

	// 2.路由分发
	r := router.InitRouter()

	// 3.启动服务
	r.Run(":8080")

	fmt.Println("http://localhost:8080")

}
