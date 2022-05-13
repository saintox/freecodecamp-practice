package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"practice-go/entity"
)

func (u *parentController) Register(c echo.Context) error {
	var input entity.UserInput
	errInput := c.Bind(&input)

	if errInput != nil {
		return c.JSON(http.StatusUnprocessableEntity, Formatter{
			errInput.Error(),
			false,
		})
	}

	createUser, errRegister := u.sqlService.RegisterUser(input)

	if errRegister != nil {
		return c.JSON(http.StatusInternalServerError, Formatter{
			errRegister.Error(),
			createUser,
		})
	}

	return c.JSON(http.StatusOK, Formatter{
		nil,
		createUser,
	})
}

func (u *parentController) Login(c echo.Context) error {
	var input entity.LoginInput

	errInput := c.Bind(&input)

	if errInput != nil {
		return c.JSON(http.StatusBadRequest, Formatter{
			errInput.Error(),
			false,
		})
	}

	loginUser, errLogin := u.sqlService.Login(input)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, Formatter{
			errLogin.Error(),
			false,
		})
	}

	token, errToken := u.authService.GenerateTokenUser(loginUser.ID)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, Formatter{
			errToken.Error(),
			false,
		})
	}

	return c.JSON(http.StatusOK, Formatter{
		nil,
		token,
	})
}
