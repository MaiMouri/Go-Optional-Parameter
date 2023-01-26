package infrastructure

import (
	"app/models"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type UserSqlHandler struct {
	DB *gorm.DB
}

func (handler *UserSqlHandler) GetUsers() (result []models.User, err error) {
	var users []models.User
	if result := handler.DB.Find(&users); err != nil {
		err := result.Error
		return nil, err
	}

	fmt.Println(users, "from infra")
	return users, nil
}

func (handler *UserSqlHandler) GetUser(id int) (result *models.User, err error) {
	var user models.User
	handler.DB.Find(&user, id)

	return &user, nil
}

func (handler *UserSqlHandler) CreateUser(t models.User) {
	handler.DB.Create(&t)
}

func (handler *UserSqlHandler) UpdateUser(t models.User) {
	handler.DB.Update(&t)
}

func (handler *UserSqlHandler) DeleteUser(t models.User) {
	handler.DB.Delete(&t)
}

func LeapYear(year int) {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}

// Validate the input email address is valid or not
// Valid email should have @ between letters and followed by .
// . should not be the first or last character
// @ should not be the first or last character
func ValidateEmail(email string) bool {

}
