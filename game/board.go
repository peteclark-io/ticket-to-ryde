package game

import "errors"

type Board struct {
	Activities              map[string]*Activity
	Coordinates             map[string]*Coordinate
	Connections             []Connection
	ConnectionsByActivityID map[string][]Connection
}

type Connection struct {
	Distance    int
	Origin      string
	Destination string
}

type Coordinate struct {
	ActivityID string  `json:"activity_id"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
}

func (b *Board) init() {
	b.ConnectionsByActivityID = make(map[string][]Connection)
	for _, conn := range b.Connections {
		orig := b.ConnectionsByActivityID[conn.Origin]
		b.ConnectionsByActivityID[conn.Origin] = append(orig, conn)

		dest := b.ConnectionsByActivityID[conn.Destination]
		b.ConnectionsByActivityID[conn.Destination] = append(dest, conn.invert())
	}
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
		connections := b.ConnectionsByActivityID[p.LastActivity]
		for _, conn := range connections {
			if conn.Destination == p.MovedTowardsActivity {
				neededMovement := conn.Distance - p.Moved
				choices = append(choices, newChoice(MovementChoiceType, neededMovement, conn.Destination)) // move to next activity
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

	connections := b.ConnectionsByActivityID[activity.ID]
	for _, conn := range connections {
		choices = append(choices, newChoice(MovementChoiceType, conn.Distance, conn.Destination))
	}

	return choices
}

func (c Connection) invert() Connection {
	return Connection{Origin: c.Destination, Destination: c.Origin, Distance: c.Distance}
}
