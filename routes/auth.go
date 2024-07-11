package routes

import (
	"test_fullstack/handlers"
	"test_fullstack/pkg/middleware"
	"test_fullstack/pkg/mysql"
	"test_fullstack/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
