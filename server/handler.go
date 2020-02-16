package main

import (
	"net/http"
  "log"

	echo "github.com/labstack/echo/v4"
	"gobblet-gobblers-go/logic"
)

func InitGame(c echo.Context) error {
	var param struct {
		P1 string `json:"p1"`
		P2 string `json:"p2"`
	}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"user param error"})
	}
	uid, _ := logic.GameInit(param.P1, param.P2)
	var res struct {
		Id string `json:"room_id"`
	}
	res.Id = uid
  log.Printf("Create Game : \n- RoomID:%s \n- P1:%s \n- P2:%s\n",uid,param.P1,param.P2)
	return c.JSON(http.StatusOK, res)
}

func GetData(c echo.Context) error {
	room := c.Param("uid")
	if room == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{"roomname not found"})
	}
	game := logic.GetGame(room)
	if game == nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"roomname not found"})
	}
	res := GameReponse{
		RoomId:    room,
		Board:     BoardToResponse(game.GetBoard()),
		P1:        *PlayerToResponse(game.GetPlayer1()),
		P2:        *PlayerToResponse(game.GetPlayer2()),
		Turn:      game.GetTurn(),
		RestPiece: RestPiecesToRestPiecesResponse(game.GetRestPieces()),
		AllPiece:  RestPiecesToRestPiecesResponse(game.GetAllPieces()),
		Win:       PlayerToResponse(game.GetPlayerInP(game.GetBoard().Judge())),
	}
	return c.JSON(http.StatusOK, res)
}

func PutPiece(c echo.Context) error {
	room := c.Param("uid")
	if room == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{"roomname not found"})
	}
	game := logic.GetGame(room)
	if game == nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{"roomname not found"})
	}
	var param struct {
		Name string `json:"name"`
		Id   uint   `json:"piece_id"`
		X    uint   `json:"piece_x"`
		Y    uint   `json:"piece_y"`
	}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"cannot parse parameter"})
	}
	if !game.IsMyTurn(param.Name) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"no your turn"})
	}
	if !game.IsMyPiece(param.Name, param.Id) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"the piece is not yours"})
	}
	if ok := game.PutPiece(param.Id, param.X, param.Y); !ok {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"can not put"})
	}
	game.ChangeTurn()
	return c.JSON(http.StatusBadRequest, OkResponse{"OK"})
}
