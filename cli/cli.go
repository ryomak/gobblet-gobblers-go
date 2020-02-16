package main

import (
	"flag"

	"gobblet-gobblers-go/cli/util"
)

var (
	url  = flag.String("url", "http://localhost:8080", "server host")
	room = flag.String("room", "", "if exist roomId")
	me   = flag.String("me", "player1", "me")
	op   = flag.String("op", "player2", "相手")
)

func init() {
	flag.Parse()
}

func main() {
	app := util.Init(url, room, me, op)
	app.Run()
}
