package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	rootRouter := r.Group("/")
	RegisterIndexRoutes(rootRouter)
	
	articleRouter := r.Group("/article")
	RegisterArticleRoutes(articleRouter)
}