package main

import (
	"gentleman/bootstrap"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	err := r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
