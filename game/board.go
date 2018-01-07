package game

import "errors"

type Board struct {
	Activities  map[string]*Activity
	Connections map[string][]Connection
}

type Connection struct {
	Distance  int
	ActivityA string
	ActivityB string
}

func (b *Board) GetActivity(id string) *Activity {
	return b.Activities[id]
}

func (b *Board) Move(position *PlayerPosition, choice *Choice, turn *Turn) *PlayerPosition {
	if !(choice.Type == MovementChoiceType || choice.Type == RideBusChoiceType) {
		panic("you're attempting to move with an invalid choice type")
	}

	activity := b.Activities[choice.ActivityID]
	if turn.RemainingAP >= choice.Cost {
		position.District = activity.DistrictID
		position.LastActivity = activity.ID
		position.Moved = 0
		position.MovedTowardsActivity = ""

		turn.Spend(choice.Cost)
		return position
	}

	position.Moved = position.Moved + turn.RemainingAP
	position.MovedTowardsActivity = choice.ActivityID

	turn.Spend(turn.RemainingAP)
	return position
}

func (b *Board) GetChoicesForPosition(p *PlayerPosition) []*Choice {
	choices := []*Choice{pickupCardChoice, playCardChoice, finishTurnChoice}

	if p.Moved != 0 {
		choices = append(choices, newChoice(MovementChoiceType, p.Moved, p.LastActivity)) // go back to last activity
		connections := b.Connections[p.LastActivity]
		for _, conn := range connections {
			if conn.ActivityB == p.MovedTowardsActivity {
				neededMovement := conn.Distance - p.Moved
				choices = append(choices, newChoice(MovementChoiceType, neededMovement, conn.ActivityB)) // move to next activity
			}
		}
		return choices
	}

	activity, ok := b.Activities[p.LastActivity]
	if !ok {
		panic(errors.New("last activity not found, something fucked up"))
	}

	choices = append(choices, newChoice(ActivityChoiceType, activity.APCost, activity.ID))
	if activity.HasBusStop { // TODO: Fix this hardcoded ride bus number!
		choices = append(choices, newChoice(RideBusChoiceType, 3, activity.ID))
	}

	connections := b.Connections[activity.ID]
	for _, conn := range connections {
		choices = append(choices, newChoice(MovementChoiceType, conn.Distance, conn.ActivityB))
	}

	return choices
}
