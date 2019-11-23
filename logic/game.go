package logic

import (
	"fmt"

	"github.com/google/uuid"
)

var Room map[string]*Game

func init() {
	Room = make(map[string]*Game)
}

type Game struct {
	restPieces []*Piece
	allPieces  []*Piece
	board      Board
	player1    *Player
	player2    *Player
	turn       P
}

func (g *Game) GetRestPieces() []*Piece {
	return g.restPieces
}

func (g *Game) SetRestPieces(ps []*Piece) {
	g.restPieces = ps
}

func (g *Game) GetAllPieces() []*Piece {
	return g.allPieces
}

func (g *Game) SetAllPieces(ps []*Piece) {
	g.allPieces = ps
}

func (g *Game) GetBoard() Board {
	return g.board
}

func (g *Game) GetPlayer1() *Player {
	return g.player1
}

func (g *Game) GetPlayer2() *Player {
	return g.player2
}

func (g *Game) SetPlayer1(name string) {
	p1 := new(Player)
	p1.SetName(name)
	p1.SetP(P1)
	g.player1 = p1
}

func (g *Game) SetPlayer2(name string) {
	p2 := new(Player)
	p2.SetName(name)
	p2.SetP(P2)
	g.player2 = p2
}

func (g *Game) SetBoard(b Board) {
	g.board = b
}

func (g *Game) ChangeTurn() P {
	g.turn = g.turn * -1
	return g.turn
}

func (g *Game) GetTurn() P {
	return g.turn
}

func GameInit(p1, p2 string) (string, bool) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", false
	}
	game := new(Game)
	game.SetBoard(BoardInit())
	game.SetPlayer1(p1)
	game.SetPlayer2(p2)
	game.turn = 1

	game.SetRestPieces([]*Piece{
		NewPiece(1, P1, SMALL),
		NewPiece(2, P1, SMALL),
		NewPiece(3, P1, MEDIUM),
		NewPiece(4, P1, MEDIUM),
		NewPiece(5, P1, BIG),
		NewPiece(6, P1, BIG),
		NewPiece(7, P2, SMALL),
		NewPiece(8, P2, SMALL),
		NewPiece(9, P2, MEDIUM),
		NewPiece(10, P2, MEDIUM),
		NewPiece(11, P2, BIG),
		NewPiece(12, P2, BIG),
	})

	game.SetAllPieces([]*Piece{
		NewPiece(1, P1, SMALL),
		NewPiece(2, P1, SMALL),
		NewPiece(3, P1, MEDIUM),
		NewPiece(4, P1, MEDIUM),
		NewPiece(5, P1, BIG),
		NewPiece(6, P1, BIG),
		NewPiece(7, P2, SMALL),
		NewPiece(8, P2, SMALL),
		NewPiece(9, P2, MEDIUM),
		NewPiece(10, P2, MEDIUM),
		NewPiece(11, P2, BIG),
		NewPiece(12, P2, BIG),
	})

	Room[u.String()] = game
	return u.String(), true
}

func GetGame(uid string) *Game {
	room, ok := Room[uid]
	if !ok {
		return nil
	}
	return room
}

func DeleteGame(uid string) {
	delete(Room, uid)
}

func (g *Game) IsMyTurn(name string) bool {
	p1 := g.GetPlayer1()
	p2 := g.GetPlayer2()
	if p1.GetName() == name {
		return g.GetTurn() == p1.GetP()
	} else if p2.GetName() == name {
		return g.GetTurn() == p2.GetP()
	}
	return false
}

func (g *Game) IsMyPiece(name string, id uint) bool {
	p1 := g.GetPlayer1()
	p2 := g.GetPlayer2()
	var p P
	if p1.GetName() == name {
		p = p1.GetP()
	} else if p2.GetName() == name {
		p = p2.GetP()
	} else {
		return false
	}
	for _, b := range g.GetAllPieces() {
		if b.GetId() == id {
			return p == b.GetUser()
		}
	}
	return false
}

func (g *Game) PutPiece(id uint, x, y uint) bool {
	//validation
	if id > 12 || x >= BOARD_MAX || y >= BOARD_MAX {
		fmt.Println("bad param", id, x, y)
		return false
	}
	piece, index, exist := g.GetPieceInRestPieces(id)
	if exist {
		if ok := g.GetBoard().Put(piece, x, y); !ok {
			return false
		} else {
			g.SetRestPieces(append(g.GetRestPieces()[:index], g.GetRestPieces()[index+1:]...))
			return true
		}
	}
	piece, bx, by, ok := g.GetBoard().GetPieceInBoard(id)
	if !ok {
		return false
	}
	return g.GetBoard().Move(piece, bx, by, x, y)
}

func (g *Game) GetPieceInRestPieces(id uint) (*Piece, int, bool) {
	for i, v := range g.GetRestPieces() {
		if v.GetId() == id {
			return v, i, true
		}
	}
	return nil, -1, false
}

func (g *Game) GetPlayerInP(p *P) *Player {
	if p == nil {
		return nil
	}
	p1 := g.GetPlayer1()
	p2 := g.GetPlayer2()
	if p1.GetP() == *p {
		return p1
	} else if p2.GetP() == *p {
		return p2
	}
	return nil
}
