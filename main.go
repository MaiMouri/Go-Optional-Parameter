package main

import (
	"app/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	// gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DBマイグレート
func dbInit() {
	db := GetDBConnection()
	db.AutoMigrate(&Todo{}, &User{})
	defer db.Close()
}

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:secret@tcp(db:3306)/sample?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	dbInit()

	//Index
	r.GET("/", controller.DbGetAll)

	//Create
	r.POST("/new", controller.DbInsert)

	//Detail
	r.GET("/detail/:id", controller.DbGetOne)

	//削除確認
	r.GET("/delete_check/:id", controller.DbDeleteCheck)

	//Delete
	r.POST("/delete/:id", controller.DbDelete)

	r.Run(":8080")
}
