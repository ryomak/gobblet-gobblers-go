package util

import (
	termbox "github.com/nsf/termbox-go"
)

func init() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	TerminalWidth, TerminalHeight = termbox.Size()

	panes["BoardPane"] = new(BoardPane)
	panes["InputPane"] = new(InputPane)
	panes["RestPiecePane"] = new(RestPiecePane)
	panes["InfoPane"] = new(InfoPane)
}

func (app *App) DisplayAll() {
	for {
		UpdateTerminal()
		DisplayBoardFrame(app)
		DisplayPane()
		DisplayInfoFrame(app)
		DisplayRestPiecesFrame(app)
		DisplayInputFrame(app)
		termbox.Flush()
	}
}

func UpdateTerminal() {
	TerminalWidth, TerminalHeight = termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func DisplayPane() {
	//縦のライン
	writeVerticalLine(panes["RestPiecePane"].MaxX(), 0, LineOption{
		Length: panes["RestPiecePane"].MaxY(),
	})
	//縦のライン
	writeVerticalLine(panes["BoardPane"].MaxX(), 0, LineOption{
		Length: panes["RestPiecePane"].MaxY(),
	})
	//下のライン
	writeLine(0, panes["BoardPane"].MaxY(), LineOption{})
}

func DisplayBoardFrame(app *App) {
	app.DisplayBoard()
}

func DisplayInfoFrame(app *App) {
	app.DisplayInfo()
}

func DisplayRestPiecesFrame(app *App) {
	app.DisplayRestPieces()
}

func DisplayInputFrame(app *App) {
	app.DisplayInput()
}
