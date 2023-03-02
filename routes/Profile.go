package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	ProfileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(ProfileRepository)

	e.GET("/profile/:id", h.GetProfile)
	e.GET("/profile", h.FindProfile)
	e.POST("/profile", h.CreateProfil)

}
