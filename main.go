package main

import (
	"backEnd/database"
	"backEnd/pkg/mysql"
	"backEnd/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
