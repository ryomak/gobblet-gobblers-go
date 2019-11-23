package util

import "fmt"

func (g *Game) PrintBoard() {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			pieces := g.Board[y][x]
			for i := len(pieces) - 1; i >= 0; i-- {
				if pieces[i] != nil {
					user := pieces[i].User
					size := pieces[i].Size
					fmt.Printf(g.ColorAttr(user), size)
					break
				}
				if i == 0 && pieces[0] == nil {
					fmt.Print("-")
				}
			}
		}
		fmt.Println()
	}
}

func (g *Game) PrintRestPieces(name string) {
	user := 0
	if g.P1.Name == name {
		user = g.P1.MyTurn
	} else if g.P2.Name == name {
		user = g.P2.MyTurn
	}
	for _, v := range g.RestPiece {
		if user == v.User {
			fmt.Printf("ID:%+v,Size:%+v\n", v.ID, v.Size)
		}
	}
}

func (g *Game) ColorAttr(user int) string {
	if user == g.P1.MyTurn {
		return "\x1b[31m%d\x1b[0m"
	} else {
		return "\x1b[34m%d\x1b[0m"
	}
}
