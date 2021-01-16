package main

import (
	"github.com/gin-gonic/gin"
	funcs "github.com/mytysoldier/kappa_web/funcs/todo"
)

// パスルーティングの設定
func router() {
	router := gin.Default()
	// テンプレートファイルの読み込み
	router.LoadHTMLGlob("templates/*.html")
	// 静的ファイルの読み込み(CSS, JS)
	router.Static("/assets", "./assets")

	// トップ画面
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "kappa_list.html", gin.H{
			"title": 0,
		})
	})

	// かっぱ召喚画面
	router.GET("/kappa_create", func(ctx *gin.Context) {
		ctx.HTML(200, "kappa_create.html", gin.H{
			"title": 0,
		})
	})

	// TODO管理画面
	router.GET("/kappa_todo", funcs.TodoList())

	// TODO追加
	router.POST("/add_todo", funcs.AddTodo())

	// TODO追加
	router.POST("/update_todo", funcs.UpdateTodo())

	router.Run()
}
