package service

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"practice-go/entity"
	"practice-go/models"
)

//func validateEmail(email string) (bool, error) {
//	ctx := context.Background()
//	cekEmail, _ := s.user_repository.FindByEmail(ctx, email)
//
//	if !cekEmail {
//		return false, errors.New("email sudah digunakan")
//	}
//	return cekEmail, nil
//}

func (b *sqlService) RegisterUser(input entity.UserInput) (bool, error) {
	ctx := context.Background()

	users, errValidation := b.sqlRepository.FindByEmail(ctx, input.Email)

	if users.ID != 0 {
		return false, errors.New("email sudah terpakai")
	}

	if errValidation != nil {
		return false, errValidation
	}

	user := models.User{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if errHash != nil {
		log.Printf(errHash.Error())
		return false, errHash
	}

	user.Password = string(hashedPassword)

	_, errQuery := b.sqlRepository.CreateUser(ctx, user)

	if errQuery != nil {
		log.Printf(errQuery.Error())
		return false, errQuery
	}

	return true, nil
}

func (b *sqlService) Login(input entity.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	ctx := context.Background()

	user, errCheck := b.sqlRepository.FindByEmail(ctx, email)

	if errCheck != nil {
		return user, errCheck
	}

	errMismatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errMismatch != nil {
		return user, errMismatch
	}

	return user, nil
}
