package seed

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jaswdr/faker"

	"app/models"
)

func Seed() {

	faker := faker.New()

	fmt.Println("Creating user Starting....")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		user := models.User{}
		user.Name = faker.Person().Name()
		min := 10
		max := 30
		user.Age = rand.Intn(max-min+1) + min
		user.CreatedAt = time.Now()
		user.Email = faker.Internet().Email()

		models.CreateUser(user)

	}

	fmt.Println("Creating user Finished....")
}
