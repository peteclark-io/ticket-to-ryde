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
		"5": &Activity{
			ID:         "5",
			Name:       "Bench",
			APCost:     1,
			DistrictID: "S/W",
			Time:       anytimeActivity,
			Rewards:    []*Reward{restReward},
		},
		"6": &Activity{
			ID:         "6",
			Name:       "Petersham Bowlo",
			APCost:     3,
			DistrictID: "S/W",
			Time:       dayTimeActivity,
			Rewards:    []*Reward{foodReward, alcoholReward},
		},
		"7": &Activity{
			ID:         "7",
			Name:       "Jazz Bar",
			APCost:     2,
			DistrictID: "S/W",
			Time:       nightTimeActivity,
			Rewards:    []*Reward{foodReward, alcoholReward},
		},
	},
	Coordinates: map[string]*Coordinate{
		"1": &Coordinate{ActivityID: "1", X: -900, Y: -590},  // Emma's House
		"2": &Coordinate{ActivityID: "2", X: -1000, Y: -510}, // St Peter's Station
		"3": &Coordinate{ActivityID: "3", X: -800, Y: -500},  // Union Hotel
		"4": &Coordinate{ActivityID: "4", X: -975, Y: -460},  // Stinking Bishops
		"5": &Coordinate{ActivityID: "5", X: -790, Y: -460},  // Bench
		"6": &Coordinate{ActivityID: "6", X: -1020, Y: -360}, // Petersham Bowlo
		"7": &Coordinate{ActivityID: "7", X: -800, Y: -280},  // Jazz Bar
	},
	Connections: []Connection{
		{
			Origin:      "1",
			Destination: "2",
			Distance:    2,
		},
		{
			Origin:      "1",
			Destination: "3",
			Distance:    2,
		},
		{
			Origin:      "3",
			Destination: "5",
			Distance:    1,
		},
		{
			Origin:      "2",
			Destination: "4",
			Distance:    2,
		},
		{
			Origin:      "4",
			Destination: "6",
			Distance:    4,
		},
		{
			Origin:      "6",
			Destination: "7",
			Distance:    4,
		},
		{
			Origin:      "5",
			Destination: "7",
			Distance:    3,
		},
	},
}
