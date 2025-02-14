package routes

import (
	"net/http"
	"website/internal"

	"github.com/gin-gonic/gin"
)


func RegisterIndexRoutes(r *gin.RouterGroup) {
	// 首頁
	// GET: /
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "CT's website",
			"Items": internal.Topics,
		})
	})

	// About me 頁面
	// GET: /aboutme
	r.GET("/aboutme", func(c *gin.Context) {
		c.HTML(http.StatusOK, "aboutme.html", gin.H{})
	})
}