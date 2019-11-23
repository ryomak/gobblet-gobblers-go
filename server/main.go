package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/ryomak/reversi-ex-go/logic"
)

func display(g *Game) {
	for x := 0; x < BOARD_MAX; x++ {
		for y := 0; y < BOARD_MAX; y++ {
			piece := g.GetBoard()[y][x].GetTopPiece()
			if piece == nil {
				fmt.Print("-")
			} else {
				user := piece.GetUser()
				size := piece.GetSize()
				fmt.Printf(ColorFmt(user), size)
			}
		}
		fmt.Println()
	}
}

func ColorFmt(user P) string {
	if user == P1 {
		return "\x1b[31m%d\x1b[0m"
	} else {
		return "\x1b[34m%d\x1b[0m"
	}
}

func main() {
	uid, ok := GameInit("user1", "user2")
	if !ok {
		fmt.Println("err")
	}
	game := GetGame(uid)
	for {
		var s string
		display(game)
		fmt.Scan(&s)
		mm(game.GetBoard(), s)
		p := new(Piece)
		p.SetUser(P1)
		p.SetSize(SMALL)
	}
}

func mm(b Board, str string) {
	p, bx, by, ax, ay, ok := split(str)
	if !ok {
		return
	}
	if bx == -1 || by == -1 {
		b.Put(p, uint(ax), uint(ay))
		fmt.Println("put")
	} else {
		fmt.Println("move")
		b.Move(p.GetUser(), uint(bx), uint(by), uint(ax), uint(ay))
	}
}

//user,size,bx,by,ax,ay
func split(str string) (*Piece, int, int, int, int, bool) {
	ss := strings.Split(str, ",")
	user, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil, 0, 0, 0, 0, false
	}
	size, err := strconv.Atoi(ss[1])
	if err != nil {
		return nil, 0, 0, 0, 0, false
	}
	bx, err := strconv.Atoi(ss[2])
	if err != nil {
		bx = -1
	}
	by, err := strconv.Atoi(ss[3])
	if err != nil {
		by = -1
	}
	ax, err := strconv.Atoi(ss[4])
	if err != nil {
		return nil, 0, 0, 0, 0, false
	}
	ay, err := strconv.Atoi(ss[5])
	if err != nil {
		return nil, 0, 0, 0, 0, false
	}
	piece := new(Piece)
	piece.SetUser(P(user))
	piece.SetSize(Size(size))
	return piece, bx, by, ax, ay, true
}
