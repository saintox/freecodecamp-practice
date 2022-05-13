package main

import (
	"github.com/labstack/echo/v4"
	"practice-go/configs"
	"practice-go/handler"
	"practice-go/repositories"
	"practice-go/route"
	"practice-go/service"
)

func main() {
	repo := repositories.NewRepository()
	sqlService := service.NewSqlService(repo)
	authService := configs.NewServiceAuth()
	controller := handler.NewParentController(sqlService, authService)

	e := echo.New()
	route.User(e, controller)
	route.Book(e, controller)

	e.Logger.Fatal(e.Start(":" + "8080"))
}
