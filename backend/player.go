package backend

type Player struct {
	Actor
}

func NewPlayer() (player *Player) {
	player = new(Player)
	player.Type = TypePlayer
	return player
}

func (player *Player) ChooseAction(g *GameState) (action Action) {
	// TODO
	return nil
}
