package main

import (
	"bookstore/dao/mysql"
	"bookstore/router"
)

func main() {

	// 1.Connect Database
	mysql.InitMysql()

	// 2.Routing distribution
	r := router.InitRouter()

	// 3.start service
	r.Run(":8080")

}
