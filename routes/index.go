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
		var topics []internal.Topic

		topicNames, err := internal.GetTopics()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"無法取得標題列表": err.Error()})
			return
		}

		for _, name := range topicNames {
			topic := internal.Topic{
				Name: name,
				Path: "/article/" + name,
			}
			topics = append(topics, topic)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "CT's website",
			"Items": topics,
		})
	})

	// About me 頁面
	// GET: /aboutme
	r.GET("/aboutme", func(c *gin.Context) {
		c.HTML(http.StatusOK, "aboutme.html", gin.H{})
	})
}