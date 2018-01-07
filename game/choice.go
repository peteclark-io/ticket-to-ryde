package game

const (
	MovementChoiceType   = "move"
	ActivityChoiceType   = "activity"
	PickupCardChoiceType = "pickup-card"
	PlayCardChoiceType   = "play-card"
	RideBusChoiceType    = "ride-bus"
	FinishTurnChoiceType = "finish-turn"
)

var (
	pickupCardChoice = newChoice(PickupCardChoiceType, 1, "")
	playCardChoice   = newChoice(PlayCardChoiceType, 1, "")
	finishTurnChoice = newChoice(FinishTurnChoiceType, 0, "")
)

type Choice struct {
	Type       string `json:"type"`
	Cost       int    `json:"cost"`
	ActivityID string `json:"activity_id,omitempty"`
}

func newChoice(choiceType string, cost int, activityID string) *Choice {
	return &Choice{Type: choiceType, Cost: cost, ActivityID: activityID}
}
