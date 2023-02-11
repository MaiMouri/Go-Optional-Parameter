package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

var Conn *gorm.DB

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:secret@tcp(db:3306)/sample?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetTodos() []Todo {
	db := GetDBConnection()
	var todos []Todo
	db.Find(&todos)
	db.Close()
	fmt.Println(todos, "from models")
	return todos
}

func GetTodo(id int) Todo {
	db := GetDBConnection()
	var todo Todo
	db.Find(&todo, id)
	db.Close()
	fmt.Println(todo, "from models")
	return todo
}

func CreateTodo(t *Todo) {
	db := GetDBConnection()
	db.Create(&t)
}

func UpdateTodo(t *Todo) {
	db := GetDBConnection()
	db.Update(&t)
}

func DeleteTodo(t *Todo) {
	db := GetDBConnection()
	db.Delete(&t)
}

func RawSQL(q string) []Todo {
	db := GetDBConnection()
	var todos []Todo
	db.Raw(q).Scan(&todos)
	return todos
}
