# gobblet-gobblers-go(Gobblet Gobblers)
[![GoDoc](https://godoc.org/github.com/ryomak/reversi-ex-go?status.svg)](https://godoc.org/github.com/ryomak/reversi-ex-go)
[![GoReport](https://goreportcard.com/badge/github.com/ryomak/reversi-ex-go)](https://goreportcard.com/report/github.com/ryomak/reversi-ex-go)

Gobblet Gobblers clone written in Go

## Introduction
ガキ使で紹介されていたボードゲームで面白そうだったGobblet Gobblersのクローンです
ロジックはAPIサーバで実装し、クライアントはCLI/webで実装します
現在:CLIのみのみなので
是非WebやCLIリファクタや実装お待ちしております。

Gobblet Gobblersの紹介サイトを見つけたので、共有します
https://kyoheiomi.com/board-game/recommendation/gobblet-gobblers
## cli
<img width="1275" alt="スクリーンショット 2019-11-24 17 17 44" src="https://user-images.githubusercontent.com/21288308/69492032-1f5eef80-0ee0-11ea-91fb-d62c0cddab85.png">
<img width="1277" alt="スクリーンショット 2019-11-24 17 24 36" src="https://user-images.githubusercontent.com/21288308/69492034-21c14980-0ee0-11ea-9cc2-3dbcf477f5aa.png">


### Usage
#### 1. ```go run server/server.go```
starting game server
#### 2. ```go run cli/cli.go -me player1 -op player2```
player1 create room 
#### 3. ```go run cli/cli.go -me player2 -room xxxx-xxxxxxx-xxxx```
player2 join room with roomId(have player2 tell player2 roomId)

## License
MIT
