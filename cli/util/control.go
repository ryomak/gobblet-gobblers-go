package util

import (
	termbox "github.com/nsf/termbox-go"
)

func (app *App) Control() {
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyCtrlC:
				break mainloop
			case termbox.KeyBackspace:
				app.inputStr = app.inputStr[:len(app.inputStr)-1]
			case termbox.KeyBackspace2:
				app.inputStr = app.inputStr[:len(app.inputStr)-1]
			case termbox.KeyEnter:
				id, x, y, ok := split(string(app.inputStr))
				if !ok {
					app.errStr = "入力データがおかしいです"
					app.inputStr = []rune{}
					continue
				}
				_, err := app.client.PutPiece(id, x, y)
				if err != nil {
					app.errStr = err.Error()
				} else {
					app.errStr = ""
				}
				app.inputStr = []rune{}
				gg, err := app.client.GetRoom()
				if err == nil {
					app.game = gg
				}
			default:
				app.inputStr = append(app.inputStr, ev.Ch)
			}
		}
	}
}

func (app *App) WatchControl() {
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyCtrlC:
				break mainloop
			}
		}
	}
}
