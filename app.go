package main

import (
	"net/http"
	"website/internal"
	"website/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 設置 HTML Template 位置
	r.LoadHTMLGlob("templates/**/*")

	// 設置靜態資源, 例如圖片、CSS、JavaScript 等
	r.Static("/static", "./static")

	if internal.HasInitErrors() {
		// Middleware，在每次 request 時都會執行。
		r.Use(func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internal.GetInitErrors()})
			c.Abort() // 終止請求，以免 request 進入其他 handler 繼續執行
		})
	}

	// 註冊路由
	routes.RegisterRoutes(r)

	r.Run("0.0.0.0:8080")
}