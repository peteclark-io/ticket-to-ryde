package game

import "errors"

type Turn struct {
	PlayerID               string
	AP                     int
	RemainingAP            int
	HasPerformedCardAction bool
}

func newTurn(p *Player) *Turn {
	return &Turn{PlayerID: p.ID, AP: 4, RemainingAP: 4}
}

func (t *Turn) Complete() bool {
	return t.RemainingAP == 0
}

func (t *Turn) Spend(cost int) (int, error) {
	if t.RemainingAP < cost {
		return t.RemainingAP, errors.New("You don't have enough AP for that")
	}

	t.RemainingAP = t.RemainingAP - cost
	return t.RemainingAP, nil
}

func (t *Turn) FilterChoices(choices []*Choice) []*Choice {
	results := make([]*Choice, 0)
	for _, c := range choices {
		if (c.Cost > t.RemainingAP) && c.Type != MovementChoiceType {
			continue
		}
		if t.HasPerformedCardAction && (c.Type == PlayCardChoiceType || c.Type == PickupCardChoiceType) {
			continue
		}

		results = append(results, c)
	}
	return results
}
