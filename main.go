// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin" // go mod tidy を実行すると go.mod に追加される
)

func main() {
	engine := gin.Default()
	// htmlのディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// htmlに渡す変数を定義
			"message": "hello gin",
		})
	})
	engine.Run(":3000")
}

// go run main.go で実行
