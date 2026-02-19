package player

type Player struct {
	Name   string
	Score  int
	InGame bool
	Room   string
}

func NewPlayer(name string) *Player {
	return &Player{Name: name}
}

func (p *Player) JoinGame(room string) {
	if p.InGame {
		return
	}
	p.InGame = true
	p.Room = room
}

func (p *Player) LeaveGame() {
	if !p.InGame {
		return
	}
	p.InGame = false
	p.Room = ""
}

func (p *Player) IsInGame() bool {
	return p.InGame
}

func (p *Player) AddScore(points int) {
	p.Score += points
}

func (p *Player) GetScore() int {
	return p.Score
}

func (p *Player) GetName() string {
	return p.Name
}
