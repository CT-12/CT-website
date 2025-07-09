---
create_at: 2024.02.14
update_at: 2024.02.14
draft: true
tags: 
 - Go
 - Gin
 - MiddleWare
---

MiddleWare會在每次有請求時執行。執行流程大概是這樣：client 發送請求 -> server 收到請求 -> 執行 middleware -> 進入路由處理請求 -> respond。

🎯 Gin Middleware 執行流程
- 1️⃣ Client 發送請求

    使用者發送 HTTP 請求，例如 GET /home
- 2️⃣ Server 收到請求

    Gin 會解析請求，然後進入 Middleware 流程
- 3️⃣ 執行 Middleware
    
    Middleware 會在請求進入 Handler 之前 執行
    Middleware 可以 攔截、修改請求，甚至終止請求
    例如，日誌記錄、身份驗證、錯誤攔截 都是常見的 Middleware 用途
- 4️⃣ 進入路由處理請求
    
    如果 Middleware 允許請求繼續，請求就會被路由 (router) 導向相應的 Handler 來處理
- 5️⃣ 返回 Response
    
    Handler 處理完請求後，Gin 會將結果回傳給 Client

```go
func main() {
	r := gin.Default()

	// 設置 HTML Template 位置
	r.LoadHTMLGlob("templates/**/*")

	// 設置靜態資源, 例如圖片、CSS、JavaScript 等
	r.Static("/static", "./static")

	if internal.HasInitErrors() {
		// Middleware，在每次請求時都會執行。
		r.Use(func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internal.GetInitErrors()})
			c.Abort() // 終止請求，以免請求進行其他 handler 繼續執行
		})
	}

	// 註冊路由
	routes.RegisterRoutes(r)

	r.Run()
}
```
> r.Use(...) 會在每次有 request 時都回傳 `c.JSON(http.StatusInternalServerError, gin.H{"error": internal.GetInitErrors()})` 然後中止 request 繼續執行下去。


另一個例子：

```go
package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware: 記錄請求日誌
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // 記錄請求開始時間

		// 執行請求
		c.Next()

		// 請求結束後計算時間
		latency := time.Since(start)
		log.Printf("[%s] %s - %v", c.Request.Method, c.Request.URL.Path, latency)
	}
}

func main() {
	r := gin.Default()

	// 使用 Middleware
	r.Use(LoggerMiddleware())

	// 註冊路由
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Home!"})
	})

	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "About Page"})
	})

	r.Run()
}

```