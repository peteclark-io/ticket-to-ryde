package game

const (
	dayTimeActivity   = dayTime
	nightTimeActivity = nightTime
	anytimeActivity   = "day/night"
)

type Activity struct {
	ID         string
	Name       string
	DistrictID string
	Time       string
	APCost     int
	Rewards    []Reward
	HasBusStop bool
}
