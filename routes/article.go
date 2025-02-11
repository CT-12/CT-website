package routes

import (
	"net/http"
	"strings"
	"website/internal"

	"github.com/gin-gonic/gin"
)

func RegisterArticleRoutes(r *gin.RouterGroup) {
	// 取得文章列表
	// GET /article/:topic
	r.GET("/:topic", func(c *gin.Context) {
		topic := c.Param("topic")

		var articles []internal.Article

		articleNames, err := internal.GetArticles(topic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"無法取得文章列表": err.Error()})
			return
		}

		for _, name := range articleNames {
			originalName := name
			name = strings.Split(name, "_")[1] // Example name: 1_HelloWorld.md
			name = strings.TrimSuffix(name, ".md")

			article := internal.Article{
				Name: name,
				Path: "/article/" + topic + "/" + originalName,
			}
			articles = append(articles, article)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "CT's website - " + topic,
			"Items": articles,
		})
	})

	// 取得文章內容
	// GET /article/:topic/:article
	r.GET("/:topic/:article", func(c *gin.Context) {
		topic := c.Param("topic")
		article := c.Param("article")

		htmlContent, err := internal.ConvertMdToHtml(topic, article)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Fail to convert to html content: ": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "article_page.html", gin.H{
			"Topic": topic,
			"Title": strings.TrimSuffix(strings.Split(article, "_")[1], ".md"),
			"Content": htmlContent,
		})
	})
}