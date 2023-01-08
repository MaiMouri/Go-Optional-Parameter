package controller

import (
	"app/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DB追加
func DbInsert(c *gin.Context) {
	var todo models.Todo
	todo.Text = c.PostForm("text")
	todo.Status = c.PostForm("status")
	fmt.Println(todo.Text, todo.Status)
	models.CreateTodo(&todo)
	c.Redirect(301, "/")
}

// DB全取得
func DbGetAll(c *gin.Context) {
	var todos []models.Todo
	todos = models.GetTodos()
	c.HTML(200, "index.html", gin.H{"todos": todos})
	// c.JSON(http.StatusOK, gin.H{"todos", todos})
}

// DB一つ取得
func DbGetOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	todo = models.GetTodo(id)
	c.HTML(200, "detail.html", gin.H{"todo": todo})
}

// func TestGetOne(t *testing.T) {

// }

// DB更新
func DbUpdate(c *gin.Context) {
	// var todo models.Todo
	// if err != c.ShouldBindJSON(&todo); err = nil {

	// }
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.GetTodo(id)
	todo.Text = c.PostForm("text")
	todo.Status = c.PostForm("status")
	models.UpdateTodo(&todo)
	c.Redirect(301, "/")
}

// DB削除
func DbDeleteCheck(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.GetTodo(id)
	models.DeleteTodo(&todo)
	c.HTML(200, "delete.html", gin.H{"todo": todo})
}

// DB削除
func DbDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := models.GetTodo(id)
	models.DeleteTodo(&todo)
	c.Redirect(301, "/")
}
