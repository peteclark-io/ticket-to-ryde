package game

var startingPosition = &PlayerPosition{LastActivity: "1", Moved: 0, District: "S/W"}

var BasicBoard = &Board{
	Activities: map[string]*Activity{
		"1": &Activity{
			ID:         "1",
			Name:       "Emma's House",
			APCost:     3,
			DistrictID: "S/W",
			Time:       anytimeActivity,
			Rewards:    []*Reward{foodReward, alcoholReward, friendshipReward},
		},
		"2": &Activity{
			ID:         "2",
			Name:       "St. Peter's Station",
			APCost:     2,
			DistrictID: "S/W",
			Time:       dayTimeActivity,
			Rewards:    []*Reward{},
		},
		"3": &Activity{
			ID:         "3",
			Name:       "Union Hotel",
			APCost:     2,
			DistrictID: "S/W",
			Time:       nightTimeActivity,
			Rewards:    []*Reward{foodReward, alcoholReward},
		},
		"4": &Activity{
			ID:         "4",
			Name:       "Stinking Bishops",
			APCost:     2,
			DistrictID: "S/W",
			Time:       anytimeActivity,
			Rewards:    []*Reward{foodReward, alcoholReward},
		},
	},
	Coordinates: map[string]*Coordinate{
		"1": &Coordinate{ActivityID: "1", X: 0, Y: 0},
		"2": &Coordinate{ActivityID: "2", X: -240, Y: 40},
		"3": &Coordinate{ActivityID: "3", X: -90, Y: -70},
		"4": &Coordinate{ActivityID: "4", X: -380, Y: -250},
	},
	Connections: []Connection{
		{
			Origin:      "1",
			Destination: "2",
			Distance:    1,
		},
		{
			Origin:      "1",
			Destination: "3",
			Distance:    2,
		},
		{
			Origin:      "2",
			Destination: "4",
			Distance:    3,
		},
	},
}
