package game

type Game struct {
	Players   []*Player
	Board     *Board
	Positions []*PlayerPosition
	Scores    []*Score
}

func NewGame(board *Board, players ...*Player) *Game {
	g := &Game{Players: players, Board: board}
	for _, player := range players {
		g.Positions = append(g.Positions, &PlayerPosition{
			PlayerID:     player.ID,
			Moved:        startingPosition.Moved,
			District:     startingPosition.District,
			LastActivity: startingPosition.LastActivity,
		})

		g.Scores = append(g.Scores, &Score{PlayerID: player.ID})
	}

	return g
}

func (g *Game) Select(position *PlayerPosition, turn *Turn, choice *Choice) {
	switch choice.Type {
	case MovementChoiceType:
		g.Board.Move(position, choice, turn)

	case ActivityChoiceType:
		activity := g.Board.GetActivity(choice.ActivityID)
		rewards := activity.perform(turn)
		g.UpdateScore(turn.PlayerID, rewards)

	case PickupCardChoiceType:
		turn.HasPerformedCardAction = true
		turn.Spend(choice.Cost)
	case PlayCardChoiceType:
		turn.HasPerformedCardAction = true
		turn.Spend(choice.Cost)
	case FinishTurnChoiceType:
		turn.Spend(turn.RemainingAP)
	}
}

func (g *Game) UpdateScore(player string, rewards []*Reward) {
	for _, score := range g.Scores {
		if score.PlayerID == player {
			score.addRewards(rewards)
			score.total()
			return
		}
	}
}

func (g *Game) GetPlayer(player string) *Player {
	for _, p := range g.Players {
		if p.ID == player {
			return p
		}
	}
	panic("couldn't find the fucking player!!!")
}

func (g *Game) GetPosition(player string) *PlayerPosition {
	for _, p := range g.Positions {
		if p.PlayerID == player {
			return p
		}
	}
	panic("couldn't find the fucking player!!!")
}

func (g *Game) Start() *Round {
	return newRound(g.Players, dayTime)
}
