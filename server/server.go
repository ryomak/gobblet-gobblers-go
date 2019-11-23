package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ryomak/reversi-ex-go/server/src"
)

func main() {
	e := echo.New()
	e.GET("/api/room/:uid", src.GetData)
	e.POST("/api/room", src.InitGame)
	e.POST("/api/room/:uid", src.PutPiece)
	e.Logger.Fatal(e.Start(":8080"))
}
