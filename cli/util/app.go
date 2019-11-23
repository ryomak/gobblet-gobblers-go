package util

import (
	"log"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type App struct {
	client   *Client
	me       *User
	op       *User
	roomId   string
	host     string
	game     *Game
	inputStr []rune
	errStr   string
}

type User struct {
	Name string
	Turn int
}

func Init(u, r, m, o *string) *App {
	client, _ := NewClient(*u)
	app := new(App)
	app.me = new(User)
	app.op = new(User)
	app.client = client
	app.host = *u
	app.roomId = *r
	app.me.Name = *m
	app.op.Name = *o
	app.game = new(Game)
	if app.roomId == "" {
		r, err := app.client.CreateRoom(app.me.Name, app.op.Name)
		if err != nil {
			panic(err)
		}
		app.roomId = r
	} else {
		app.client.SetRoomId(app.roomId)
		app.client.SetUser(app.me.Name)
	}

	gg, err := app.client.GetRoom()
	if err != nil {
		panic(err)
	}
	app.game = gg
	app.me.Turn = gg.GetMyturn(app.me.Name)
	app.op.Turn = gg.GetMyturn(app.op.Name)
	return app
}

func (app *App) Run() {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	go app.fetchLoop()
	go app.DisplayAll()
  if app.me.Turn == 0{
    app.WatchControl()
  }else{
	  app.Control()
  }
}

func (app *App) fetchLoop() {
	for {
		//更新
		gg, err := app.client.GetRoom()
		if err == nil {
			app.game = gg
		}
		time.Sleep(3 * time.Second)
	}
}
