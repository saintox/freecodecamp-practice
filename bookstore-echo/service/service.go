package service

import (
	"practice-go/entity"
	"practice-go/models"
	"practice-go/repositories"
)

type (
	sqlService struct {
		sqlRepository repositories.SqlRepository
	}

	SqlService interface {
		RegisterUser(input entity.UserInput) (bool, error)
		Login(input entity.LoginInput) (models.User, error)
		InsertBook(input entity.BookInput) (bool, error)
	}
)

func NewSqlService(sqlRepository repositories.SqlRepository) *sqlService {
	return &sqlService{sqlRepository}
}
