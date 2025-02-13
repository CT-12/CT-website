package routes

import (
	"net/http"
	"path/filepath"
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
			// 萃取出 article 名稱
			originalName := name // Example name: 1_HelloWorld.md
			name = strings.Split(name, "_")[1] 
			name = strings.TrimSuffix(name, ".md")

			// 取得文章的 front matter	
			filePath := internal.CONTENT_DIR + "/" + topic + "/" + originalName
			frontMatter, _, err := internal.ParseMarkdown(filePath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse markdown": err.Error()})
				return
			}
			
			if frontMatter.Draft == true {
				continue 
			}

			// 建立 Article 物件
			article := internal.Article{
				Name: name,
				Path: "/article/" + topic + "/" + originalName,
				CreateAt: frontMatter.CreateAt,
				UpdateAt: frontMatter.UpdateAt,
			}
			articles = append(articles, article)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": topic,
			"Items": articles,
		})
	})

	// 取得文章內容
	// GET /article/:topic/:article
	r.GET("/:topic/:article", func(c *gin.Context) {
		topic := c.Param("topic")
		article := c.Param("article")
		md_file_path := filepath.Join(internal.CONTENT_DIR, topic, article)

		// 解析 Markdown 檔案，取得 Markdown 內容
		_, md_content, err := internal.ParseMarkdown(md_file_path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Fail to parse markdown": err.Error()})
			return
		}

		// 轉換 Markdown 內容到 HTML
		htmlContent, err := internal.ConvertMdToHtml(md_content)
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