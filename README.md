# reversi-ex-go

Gobblet Gobblers clone written in Go

## Introduction
ガキ使で紹介されていたGobblet Gobblersのクローンです
ロジックはAPIサーバで実装し、クライアントはCLI/webで実装します
現在:CLIのみのみなので
是非WebやCLI修正プルリクお願いいたします

## cli


### Usage
1. ```go run server/server.go```
starting game server
2. ```go run cli/cli.go -me player1 -op player2```
player1 create room 
3. ```go run cli/cli.go -me player2 -room xxxx-xxxxxxx-xxxx```
player2 join room with roomId(have player2 tell player2 roomId)

## License
MIT

