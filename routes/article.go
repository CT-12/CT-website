package routes

import (
	"fmt"
	"net/http"
	"website/internal"

	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(r *gin.RouterGroup) {
	// 取得文章列表
	// GET /article/:topic
	r.GET("/:topic", func(c *gin.Context) {
		topic := c.Param("topic")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": topic,
			"Items": internal.Topic2Articles[topic],
		})
	})

	// 取得文章內容
	// GET /article/:topic/:article
	r.GET("/:topic/:article", func(c *gin.Context) {
		topic := c.Param("topic")
		article := c.Param("article")
		
		// 找出文章
		var articleObj internal.Article
		for _, tmp_articleObj := range internal.Topic2Articles[topic] {
			if tmp_articleObj.FileName == article {
				articleObj = tmp_articleObj
			}
		}
		if articleObj.FileName == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Article %s not found", article)})
			return
		}

		// 轉換 Markdown 內容到 HTML
		htmlContent, err := internal.ConvertMdToHtml(articleObj.Markdown)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Fail to convert to html content: ": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "article_page.html", gin.H{
			"Topic": topic,
			"Title": articleObj.Name,
			"Content": htmlContent,
		})
	})
}