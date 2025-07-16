package routes

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	"website/config"
	urlshorten "website/internal/projects/url_shorten"
	"website/internal/redis"
)

func RegisterProjectRoutes(r *gin.RouterGroup) {
	// 專案列表
	// GET: /project/:project
	r.GET("/:project", func(c *gin.Context) {
		project := c.Param("project")
		c.HTML(http.StatusOK, project + ".html", gin.H{})
	})

	// POST: /project/url_shorten
	r.POST("/url_shorten", func(ctx *gin.Context) {
		var req urlshorten.UrlShortenRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Request error: " + err.Error()})
			return
		}

		short_url := urlshorten.GenerateShortURL(req.LongUrl)

		if err := redis.SetValue(short_url, req.LongUrl, 3); err != nil {
			log.Printf("Error: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL."})
			return
		} 

		ctx.JSON(http.StatusOK, gin.H{
			"short_url": fmt.Sprintf("https://%s/short/%s", config.Domain_name, short_url),
		})
	})
}