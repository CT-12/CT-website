package routes

import (
	"log"
	"net/http"
	"website/internal"
	"website/internal/redis"

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

	// 短網址重定向
	// GET: /s/:short_url
	r.GET("/s/:short_url", func(ctx *gin.Context) {
		short_url := ctx.Param("short_url")

		if short_url == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Short URL is required"})
			return
		}

		long_url, err := redis.GetValue(short_url)
		if err != nil {
			log.Printf("Error retrieving long URL for short URL %s: %v", short_url, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve long URL."})
			return
		}

		ctx.Redirect(http.StatusMovedPermanently, long_url)
	})
}