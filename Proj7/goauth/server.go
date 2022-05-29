package main

import (
	"github.com/labstack/echo/v4"
	"goauth/database"
	"goauth/handlers/auth/github"
	"goauth/handlers/auth/google"
)

func main() {
	e := echo.New()
	database.GetDatabase()
	e.GET("/login/google", google.LoginHandler)
	e.GET("/auth/google", google.AuthHandler)

	e.GET("/login/github", github.LoginHandler)
	e.GET("/auth/github", github.AuthHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
