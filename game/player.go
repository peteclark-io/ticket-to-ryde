package game

type Player struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Perks []*Perk `json:"perks"`
}

type PlayerPosition struct {
	PlayerID             string `json:"player_id"`
	LastActivity         string `json:"last_activity"`
	Moved                int    `json:"moved"`
	MovedTowardsActivity string `json:"moved_towards_activity"`
	District             string `json:"district"`
}

func NewPlayer(id string, name string) *Player {
	return &Player{ID: id, Name: name}
}
