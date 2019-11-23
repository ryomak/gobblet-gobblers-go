package logic

func SqureInit() *Square {
	ps := [3]*Piece{}
	square := new(Square)
	square.top = -1
	square.pieces = ps
	return square
}

func (s *Square) GetTop() int {
	return s.top
}

func (s *Square) GetPieces() [3]*Piece {
	return s.pieces
}

func (s *Square) GetTopPiece() *Piece {
	if s.GetTop() == -1 {
		return nil
	}
	return s.pieces[s.top]
}

func (s *Square) SetTopPiece(p *Piece) bool {
	if ok := s.IncTop(); !ok {
		return false
	}
	s.pieces[s.GetTop()] = p
	return true
}

func (s *Square) RemoveTopPiece() (*Piece, bool) {
	p := s.pieces[s.GetTop()]
	if ok := s.DecTop(); !ok {
		return nil, false
	}
	s.pieces[s.GetTop()+1] = nil
	return p, true
}

func (s *Square) SetTop(index int) {
	s.top = index
}

func (s *Square) IncTop() bool {
	if s.top == 2 {
		return false
	}
	s.top++
	return true
}

func (s *Square) DecTop() bool {
	if s.top == -1 {
		return false
	}
	s.top--
	return true
}
