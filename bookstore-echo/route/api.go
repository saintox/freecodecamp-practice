package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"practice-go/handler"
)

func User(route *echo.Echo, controller handler.ParentController) {
	api := route.Group("/api/v1/")
	api.POST("register", controller.Register)
	api.POST("login", controller.Login)
}

func Book(route *echo.Echo, controller handler.ParentController) {
	api := route.Group("/api/v1/")
	api.POST("new-book", controller.NewBook, middleware.JWT([]byte("iloveindonesia")))
}
