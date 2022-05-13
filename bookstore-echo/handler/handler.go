package handler

import (
	"github.com/labstack/echo/v4"
	"practice-go/configs"
	"practice-go/service"
)

type (
	parentController struct {
		sqlService  service.SqlService
		authService configs.AuthService
	}

	Formatter struct {
		Message interface{} `json:"message"`
		Data    interface{} `json:"data"`
	}

	ParentController interface {
		Register(c echo.Context) error
		Login(c echo.Context) error
		NewBook(c echo.Context) error
	}
)

func NewParentController(sqlService service.SqlService, authService configs.AuthService) *parentController {
	return &parentController{
		sqlService,
		authService,
	}
}
