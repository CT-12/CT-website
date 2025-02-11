package main

import (
	"website/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 設置 HTML Template 位置
	r.LoadHTMLGlob("templates/*")

	// 設置靜態資源, 例如圖片、CSS、JavaScript 等
	r.Static("/static", "./static")

	// 註冊路由
	routes.RegisterRoutes(r)

	r.Run(":8080")
}