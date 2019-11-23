package logic

func (p *Piece) GetUser() P {
	return p.user
}

func (p *Piece) SetUser(u P) bool {
	p.user = u
	return true
}

func (p *Piece) GetSize() Size {
	return p.size
}

func (p *Piece) SetSize(size Size) bool {
	p.size = size
	return true
}

func (p *Piece) GetId() uint {
	return p.id
}

func NewPiece(id uint, user P, size Size) *Piece {
	p := new(Piece)
	p.user = user
	p.size = size
	p.id = id
	return p
}
