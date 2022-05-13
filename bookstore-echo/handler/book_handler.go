package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"practice-go/entity"
)

func (u *parentController) NewBook(c echo.Context) error {
	var input entity.BookInput
	errInput := c.Bind(&input)

	if errInput != nil {
		return c.JSON(http.StatusUnprocessableEntity, Formatter{
			errInput.Error(),
			false,
		})
	}

	createBook, errCreate := u.sqlService.InsertBook(input)

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, Formatter{
			errCreate.Error(),
			createBook,
		})
	}

	return c.JSON(http.StatusOK, Formatter{
		nil,
		createBook,
	})
}
