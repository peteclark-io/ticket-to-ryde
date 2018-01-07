package game

var startingPosition = &PlayerPosition{LastActivity: "1", Moved: 0, District: "S/W"}

var BasicBoard = &Board{
	Activities: map[string]*Activity{
		"1": &Activity{
			ID:         "1",
			Name:       "Emma's House",
			APCost:     3,
			DistrictID: "S/W",
			Rewards:    []*Reward{foodReward, alcoholReward},
		},
		"2": &Activity{
			ID:         "2",
			APCost:     4,
			Name:       "Manly Beach",
			DistrictID: "N/E",
			Rewards:    []*Reward{foodReward, alcoholReward, experienceReward},
		},
	},
	Connections: map[string][]Connection{
		"1": {
			{
				ActivityA: "1",
				ActivityB: "2",
				Distance:  3,
			},
		},
		"2": {
			{
				ActivityA: "2",
				ActivityB: "1",
				Distance:  3,
			},
		},
	},
}
