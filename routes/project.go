package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(r *gin.RouterGroup) {
	// 專案列表
	// GET: /project/:project
	r.GET("/:project", func(c *gin.Context) {
		project := c.Param("project")
		c.HTML(http.StatusOK, project + ".html", gin.H{})
	})
}