package main

import (
	"github.com/kiditz/spgku-api/db"
	r "github.com/kiditz/spgku-api/router"
	trans "github.com/kiditz/spgku-api/translate"

	_ "github.com/kiditz/spgku-api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	db.Connect()
	db.DB.LogMode(true)
}

// @title Spgku Application
// @description This is event staffing application management
// @version 1.0
// @host localhost:8000
// @BasePath /api/v1
func main() {

	e := echo.New()
	trans.InitTranslate(e)

	v1 := e.Group("/api/v1")
	r.InitRoutes(v1)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	// e.Use(middleware.Gzip())
	e.Logger.Fatal(e.Start(":8000"))
}
