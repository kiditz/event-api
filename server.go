package main

import (
	"github.com/kiditz/spgku-job/db"
	r "github.com/kiditz/spgku-job/router"
	trans "github.com/kiditz/spgku-job/translate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	db.Connect()
	db.DB.LogMode(true)
}
func main() {

	e := echo.New()
	trans.InitTranslate(e)
	r.InitRoutes(e)
	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":8000"))
}
