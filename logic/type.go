package logic

const (
	_ Size = iota
	SMALL
	MEDIUM
	BIG
)

const P1 P = 1
const P2 P = -1

type Size uint
type P int

type Piece struct {
	id   uint
	size Size
	user P
}

type Square struct {
	top    int
	pieces [3]*Piece
}

const BOARD_MAX = 3

type Board [BOARD_MAX][BOARD_MAX]*Square
