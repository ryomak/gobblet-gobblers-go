package util

import (
	"fmt"
	"strconv"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var TerminalWidth, TerminalHeight = 0, 0
var panes = map[string]Pane{}

type Pane interface {
	X() int
	Y() int
	MaxX() int
	MaxY() int
}

const (
	InfoWidth      = 30
	RestPieceWidth = 30
)

type BoardPane struct{}

func (a BoardPane) X() int    { return 0 }
func (a BoardPane) Y() int    { return 0 }
func (a BoardPane) MaxX() int { return TerminalWidth - RestPieceWidth - InfoWidth }
func (a BoardPane) MaxY() int { return TerminalHeight - 3 }

type InputPane struct{}

func (a InputPane) X() int    { return 0 }
func (a InputPane) Y() int    { return TerminalHeight - 3 }
func (a InputPane) MaxX() int { return TerminalWidth }
func (a InputPane) MaxY() int { return TerminalHeight }

type RestPiecePane struct{}

func (a RestPiecePane) X() int    { return TerminalWidth - RestPieceWidth - InfoWidth }
func (a RestPiecePane) Y() int    { return 0 }
func (a RestPiecePane) MaxX() int { return TerminalWidth - InfoWidth }
func (a RestPiecePane) MaxY() int { return TerminalHeight - 3 }

type InfoPane struct{}

func (a InfoPane) X() int    { return TerminalWidth - InfoWidth }
func (a InfoPane) Y() int    { return 0 }
func (a InfoPane) MaxX() int { return TerminalWidth }
func (a InfoPane) MaxY() int { return TerminalHeight - 3 }

type LineOption struct {
	Reverse   bool
	Length    int
	CharColor termbox.Attribute
	BGColor   termbox.Attribute
}

type SentenceOption struct {
	X         int
	Y         int
	MaxX      int
	MaxY      int
	CharColor termbox.Attribute
	BGColor   termbox.Attribute
}

func writeLine(x, y int, op LineOption) {
	width := TerminalWidth
	if !op.Reverse {
		if op.Length != 0 && width > (x+op.Length) {
			width = x + op.Length
		}
		for col := x; col < width; col++ {
			termbox.SetCell(col, y, '-', op.CharColor, op.BGColor)
		}
	} else {
		if op.Length != 0 {
			width = x - op.Length
			if width < 0 {
				width = 0
			}
		}
		for col := x; col > width; col-- {
			termbox.SetCell(col, y, '-', op.CharColor, op.BGColor)
		}
	}
}

func writeVerticalLine(x, y int, op LineOption) {
	height := TerminalHeight
	if !op.Reverse {
		if op.Length != 0 && height > (y+op.Length) {
			height = y + op.Length
		}
		for row := y; row < height; row++ {
			termbox.SetCell(x, row, '|', op.CharColor, op.BGColor)
		}
	} else {
		if op.Length != 0 {
			height = x - op.Length
			if height < 0 {
				height = 0
			}
		}
		for row := y; row > height; row-- {
			termbox.SetCell(x, row, '|', op.CharColor, op.BGColor)
		}
	}
}

func writeSentense(str string, op SentenceOption) {
	runes := []rune(str)
	for _, r := range runes {
		termbox.SetCell(op.X, op.Y, r, op.CharColor, op.BGColor)
		op.X += runewidth.RuneWidth(r)
		if (op.X) > op.MaxX {
			return
		}
	}
}

func writeRune(r rune, op SentenceOption) {
	termbox.SetCell(op.X, op.Y, r, op.CharColor, op.BGColor)
	op.X += runewidth.RuneWidth(r)
	if (op.X) > op.MaxX {
		return
	}
}

func (g *Game) writePiece(r rune, user int, op SentenceOption) {
	if user == g.P1.MyTurn {
		op.CharColor = termbox.ColorRed
	} else {
		op.CharColor = termbox.ColorBlue
	}
	termbox.SetCell(op.X, op.Y, r, op.CharColor, op.BGColor)
	op.X += runewidth.RuneWidth(r)
	if (op.X) > op.MaxX {
		return
	}
}

func (app *App) DisplayBoard() {
	writeSentense(fmt.Sprintf("ルームID -> %s", app.roomId), SentenceOption{
		X:    panes["BoardPane"].X() + 5,
		Y:    panes["BoardPane"].Y() + 1,
		MaxX: panes["BoardPane"].MaxX() - 1,
	})
	topx := (panes["BoardPane"].MaxX()+panes["BoardPane"].X())/2 - 5
	topy := (panes["BoardPane"].MaxY()+panes["BoardPane"].Y())/2 - 5
	//TODO 上手くボード出
	writeSentense("A    B    C", SentenceOption{
		X:    topx,
		Y:    topy - 2,
		MaxX: panes["BoardPane"].MaxX() - 1,
	})
	for y := 0; y < 3; y++ {
		writeRune(intToChar(y), SentenceOption{
			X: topx - 3,
			Y: topy + y*4,
		})
		for x := 0; x < 3; x++ {
			pieces := app.game.Board[y][x]
			for i := len(pieces) - 1; i >= 0; i-- {
				if pieces[i] != nil {
					user := pieces[i].User
					size := pieces[i].Size
					app.game.writePiece([]rune(strconv.Itoa(size))[0], user, SentenceOption{
						X: topx + x*5,
						Y: topy + y*4,
					})
					break
				} else if i == 0 && pieces[0] == nil {
					writeRune('-', SentenceOption{
						X: topx + x*5,
						Y: topy + y*4,
					})
				}
			}
		}
	}
}

func (app *App) DisplayInfo() {
	winner := ""
	if app.game.Win != nil {
		winner = fmt.Sprintf("winnter:%s", app.game.Win.Name)
	}
	turn := ""
	if app.me.Turn == 0 {
		turn = "あなたは閲覧者です"
	} else if app.me.Turn == app.game.Turn {
		turn = "あなたの番です"
	} else {
		turn = "相手の番です"
	}
	sWithHeader := []string{
		"■ 情報",
		fmt.Sprintf("あなた:%s", app.me.Name),
		fmt.Sprintf("相手:%s", app.op.Name),
		"■ ターン",
		turn,
		"■ エラー",
		app.errStr,
		"■ 結果",
		winner,
	}
	cnY := 3
	for i, str := range sWithHeader {
		writeSentense(str, SentenceOption{
			X:    panes["InfoPane"].X() + 5,
			MaxX: panes["InfoPane"].MaxX() - 1,
			Y:    cnY + i*2,
		})
	}
}

func (app *App) DisplayRestPieces() {
	cnY := 3 + 3
	writeSentense("My ピース", SentenceOption{
		X:         panes["RestPiecePane"].X() + 5,
		Y:         3,
		MaxX:      panes["RestPiecePane"].MaxX() - 1,
		CharColor: termbox.ColorGreen,
	})
	for i, v := range app.game.AllPiece {
		if app.me.Turn != v.User {
			continue
		}
		exist := false
		for _, rest := range app.game.RestPiece {
			if rest.ID == v.ID {
				exist = true
				break
			}
		}
		color := termbox.ColorCyan
		if exist {
			color = termbox.ColorWhite
		}
		writeSentense(fmt.Sprintf("ID:%d,Size:%d", v.ID, v.Size), SentenceOption{
			CharColor: color,
			X:         panes["RestPiecePane"].X() + 5,
			Y:         cnY + i*2,
			MaxX:      panes["RestPiecePane"].MaxX() - 1,
		})
	}
}

func (app *App) DisplayInput() {
	writeSentense(fmt.Sprintf("INPUT(ex:a,A,2)-> %s", string(app.inputStr)), SentenceOption{
		X:    panes["InputPane"].X() + 5,
		Y:    panes["InputPane"].Y() + 2,
		MaxX: panes["InputPane"].MaxX() - 1,
	})
}
