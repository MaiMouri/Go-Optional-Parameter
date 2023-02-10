package main

import (
	database "app/config/database"
	"app/controller"
	"app/models"
	"fmt"

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
	db := models.GetDBConnection()
	db.AutoMigrate(&Todo{}, &User{})
	defer db.Close()
}

type GreetOpts struct {
	GreetingWord *string
}

// オプショナルパラメータを構造体で受け取る
func Greet(name string, opts *GreetOpts) {
	greetingWord := "Hello"
	if opts.GreetingWord != nil {
		// 引数がnilだったら未指定なのでデフォルト値で埋める
		greetingWord = *opts.GreetingWord
	}
	fmt.Printf("%s, %s!\n", greetingWord, name)
}

func main() {
	db := database.New()
	connect := db.DB()
	defer connect.Close()

	//DI
	// var customerRepository repository.CustomerRepository
	// customerPersistance := persistance.NewCustomerPersistance(db, customerRepository)
	// customerUseCase := usecase.NewCustomerUseCase(customerPersistance)
	// customerController := controller.NewCustomerController(customerUseCase)

	Greet("gopher", &GreetOpts{}) // Hello, gopher!

	word := "Hey"
	Greet("gopher", &GreetOpts{GreetingWord: &word}) // Hey, gopher!

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
