package main

import (
	"fmt"
	"log"
	"os"
	"test_fullstack/database"
	"test_fullstack/handlers"
	"test_fullstack/pkg/mysql"
	"test_fullstack/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.DatabaseInit()
	database.RunMigration()

	if err := handlers.SeedDummyCredentials(mysql.DB); err != nil {
		// Tangani error jika ada
		panic(err)
	}

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("server running localhost:5000")
	// e.Logger.Fatal(e.Start("localhost:5000"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
