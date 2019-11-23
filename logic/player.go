package logic

type Player struct {
	name   string
	player P
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SetName(n string) {
	p.name = n
}

func (p *Player) GetP() P {
	return p.player
}

func (pl *Player) SetP(p P) {
	pl.player = p
}
