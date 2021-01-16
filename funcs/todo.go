package funcs

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Todo ：テーブル格納用
type Todo struct {
	ID        int
	Todo      string
	DeleteFlg string
}

// TodoList ：現在のTODO一覧を取得し画面表示
func TodoList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todos := searchTodo()
		ctx.HTML(200, "kappa_todo.html", gin.H{
			"uncompTodos": todos["uncompTodos"],
			"compTodos":   todos["compTodos"],
		})
	}
}

// AddTodo ：TODOを新たに追加
func AddTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reserveTodo(ctx.PostForm("todo"))
		ctx.JSON(200, gin.H{})
	}
}

// UpdateTodo ：TODOを更新する（完了、もしくは未完了に更新）
func UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		updateTodos(strings.Split(ctx.PostForm("todos"), ","), ctx.PostForm("mode"))
		ctx.JSON(200, gin.H{})
	}
}

// TODO一覧取得
func searchTodo() map[string][]Todo {
	db, err := sql.Open("mysql", "root@/kappa")
	if err != nil {
		panic("データベース接続でエラー発生")
	}
	defer db.Close()

	// TODO一覧取得
	rows, err := db.Query("SELECT id, todo, delete_flg FROM todo")
	defer rows.Close()
	if err != nil {
		panic("クエリー発行でエラー発生")
	}
	// 未完了のTODO格納用
	uncompTodos := []Todo{}
	// 完了のTODO格納用
	compTodos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Todo, &todo.DeleteFlg)
		if err != nil {
			panic("検索結果読み込みでエラー発生")
		}
		if todo.DeleteFlg == "0" {
			uncompTodos = append(uncompTodos, todo)
		} else {
			compTodos = append(compTodos, todo)
		}
	}
	return map[string][]Todo{
		"uncompTodos": uncompTodos,
		"compTodos":   compTodos,
	}
}

// 新規TODO登録
func reserveTodo(todo string) {
	db, err := sql.Open("mysql", "root@/kappa")
	if err != nil {
		panic("データベース接続でエラー発生")
	}
	defer db.Close()

	// TODO存在チェック
	var exists bool
	err = db.QueryRow("SELECT if(count(1), 'true', 'false') FROM todo where todo = ? and delete_flg = '0'", todo).Scan(&exists)
	if err != nil {
		panic("クエリー発行でエラー発生")
	}
	if exists {
		log.Println("同名TODOは登録できません。", todo)
		return
	}

	// TODO登録
	ins, err := db.Prepare("INSERT INTO todo (todo, delete_flg, create_data, update_date) VALUES (?, '0', now(), now())")
	if err != nil {
		panic("クエリー発行でエラー発生")
	}
	ins.Exec(todo)
}

// TODO更新
func updateTodos(todos []string, updateMode string) {
	db, err := sql.Open("mysql", "root@/kappa")
	if err != nil {
		panic("データベース接続でエラー発生")
	}
	defer db.Close()

	// TODO更新
	query := "UPDATE todo SET delete_flg = '1' where todo IN (?) and delete_flg = '0'"
	if updateMode == "uncomp" {
		query = "UPDATE todo SET delete_flg = '0' where todo IN (?) and delete_flg = '1'"
	}
	upd, err := db.Prepare(query)
	if err != nil {
		panic("クエリー発行でエラー発生")
	}
	for _, todo := range todos {
		upd.Exec(todo)
	}
}
