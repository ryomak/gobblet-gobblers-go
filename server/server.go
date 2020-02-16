package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/api/room/:uid", GetData)
	e.POST("/api/room", InitGame)
	e.POST("/api/room/:uid", PutPiece)
	e.Logger.Fatal(e.Start(":8080"))
}
