package game

const (
	dayTime   = "day"
	nightTime = "night"
)

type Round struct {
	Time  string
	Turns []*Turn
	index int
}

func newRound(players []*Player, time string) *Round {
	r := &Round{Time: time, index: -1}
	for _, p := range players {
		t := newTurn(p)
		r.Turns = append(r.Turns, t)
	}
	return r
}

func (r *Round) nextTurn() *Turn {
	r.index++
	return r.Turns[r.index]
}
