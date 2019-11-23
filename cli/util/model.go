package util

type Game struct {
	RoomId string `json:"room_id"`
	Board  [][][]*struct {
		Size int `json:"size"`
		User int `json:"user"`
	} `json:"board"`
	P1 struct {
		Name   string `json:"name"`
		MyTurn int    `json:"my_turn"`
	} `json:"p1"`
	P2 struct {
		Name   string `json:"name"`
		MyTurn int    `json:"my_turn"`
	} `json:"p2"`
	Turn      int `json:"turn"`
	RestPiece []struct {
		ID   int `json:"id"`
		Size int `json:"Size"`
		User int `json:"User"`
	} `json:"rest_piece"`
	AllPiece []struct {
		ID   int `json:"id"`
		Size int `json:"Size"`
		User int `json:"User"`
	} `json:"all_piece"`
	Win *struct {
		Name   string `json:"name"`
		MyTurn int    `json:"my_turn"`
	} `json:"win"`
}

type Result struct {
	Message string `json:"message"`
}

func (g *Game) GetMyturn(name string) int {
	if g.P1.Name == name {
		return g.P1.MyTurn
	}
	if g.P2.Name == name {
		return g.P2.MyTurn
	}
	return 0
}
