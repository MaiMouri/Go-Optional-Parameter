package infrastructure

import (
	"app/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	DB *gorm.DB
}

func (handler *SqlHandler) GetTodos() (result []models.Todo, err error) {
	var todos []models.Todo
	if result := handler.DB.Find(&todos); err != nil {
		err := result.Error
		return nil, err
	}

	fmt.Println(todos, "from infra")
	return todos, nil
}

func (handler *SqlHandler) GetTodo(id int) (result *models.Todo, err error) {
	var todo models.Todo
	handler.DB.Find(&todo, id)

	return &todo, nil
}

func (handler *SqlHandler) CreateTodo(t models.Todo) {
	handler.DB.Create(&t)
}

func (handler *SqlHandler) UpdateTodo(t models.Todo) {
	handler.DB.Update(&t)
}

func (handler *SqlHandler) DeleteTodo(t models.Todo) {
	handler.DB.Delete(&t)
}

func (handler *SqlHandler) RawSQL(q string) []models.Todo {
	var todos []models.Todo
	handler.DB.Raw(q).Scan(&todos)
	return todos
}
