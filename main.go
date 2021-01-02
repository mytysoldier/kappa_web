package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// テンプレートファイルの読み込み
	router.LoadHTMLGlob("templates/*.html")
	// 静的ファイルの読み込み
	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "kappa_create.html", gin.H{
			"title": 0,
		})
	})

	router.Run()
}
