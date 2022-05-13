package repositories

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"practice-go/models"
	"testing"
)

func TestUserInsert(t *testing.T) {
	userRepository := NewRepository()

	ctx := context.Background()

	user := models.User{
		Name:     "Mamang Racing",
		Phone:    "08123456789",
		Password: "test123",
		Email:    "mamang@mailinator.com",
	}

	testinsert, err := userRepository.CreateUser(ctx, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(testinsert)
}
