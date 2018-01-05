package game

import "errors"

type Turn struct {
	PlayerID    string
	AP          int
	RemainingAP int
}

func newTurn(p *Player) *Turn {
	return &Turn{PlayerID: p.ID, AP: 5, RemainingAP: 5}
}

func (t *Turn) spend(ap int) (int, error) {
	if t.RemainingAP < ap {
		return t.RemainingAP, errors.New("You don't have enough AP for that")
	}

	t.RemainingAP = t.RemainingAP - ap
	return t.RemainingAP, nil
}
