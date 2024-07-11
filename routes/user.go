package routes

import (
	"test_fullstack/handlers"
	"test_fullstack/pkg/middleware"
	"test_fullstack/pkg/mysql"
	"test_fullstack/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)
	e.GET("/users", middleware.Auth(h.FindUsers))
	e.GET("/user/:id", middleware.Auth(h.GetUser))
}
