package database

import (
	"app/models"

	"github.com/jinzhu/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", "root:secret@tcp(db:3306)/sample?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Todo{}, &models.User{})

	return db
}
