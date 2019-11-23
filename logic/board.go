package logic

func BoardInit() Board {
	board := Board{}
	for x := 0; x < BOARD_MAX; x++ {
		for y := 0; y < BOARD_MAX; y++ {
			board[x][y] = SqureInit()
		}
	}
	return board
}

func (b Board) CanPut(p *Piece, x, y uint) bool {
	square := b[x][y]
	if square.GetTop() == -1 {
		return true
	}
	if square.GetTopPiece().GetUser() == p.GetUser() {
		return false
	}
	return square.GetTopPiece().GetSize() < p.GetSize()
}

func (b Board) Put(p *Piece, x, y uint) bool {
	if ok := b.CanPut(p, x, y); !ok {
		return false
	}
	return b[x][y].SetTopPiece(p)
}

func (b Board) Pick(p *Piece, x, y uint) (*Piece, bool) {
	piece := b[x][y].GetTopPiece()
	if piece == nil {
		return nil, false
	}
	if piece.GetUser() != p.GetUser() {
		return nil, false
	}
	return b[x][y].RemoveTopPiece()
}

func (b Board) Move(p *Piece, bx, by, ax, ay uint) bool {
	piced, ok := b.Pick(p, bx, by)
	if !ok {
		return false
	}
	return b.Put(piced, ax, ay)
}

func (b Board) GetPieceInBoard(id uint) (*Piece, uint, uint, bool) {
	for x := 0; x < BOARD_MAX; x++ {
		for y := 0; y < BOARD_MAX; y++ {
			topPiece := b[x][y].GetTopPiece()
			if topPiece == nil {
				continue
			}
			if id == topPiece.GetId() {
				return topPiece, uint(x), uint(y), true
			}
		}
	}
	return nil, 0, 0, false
}

/*
判定ロジック
  0|  1|  2
-----------
  3|  4|  5
-----------
  6|  7|  8
*/
func (b Board) Judge() *P {
	for _, v := range winBoard {
		var ss []P
		for _, t := range v {
			piece := b[t%BOARD_MAX][t/BOARD_MAX].GetTopPiece()
			if piece == nil {
				//配置されていない場合適当なP
				ss = append(ss, 0)
			} else {
				ss = append(ss, b[t%BOARD_MAX][t/BOARD_MAX].GetTopPiece().GetUser())
			}
		}
		if isSame(ss) {
			return &ss[0]
		}
	}
	return nil
}

var winBoard = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func isSame(p []P) bool {
	s := p[0]
	for _, v := range p {
		if s != v {
			return false
		}
	}
	return true
}
