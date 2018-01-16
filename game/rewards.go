package game

const (
	foodRewardType       = "food"
	alcoholRewardType    = "alcohol"
	experienceRewardType = "experience"
	friendshipRewardType = "friendship"
)

var (
	foodReward       = &Reward{Type: foodRewardType, Score: 5}
	alcoholReward    = &Reward{Type: alcoholRewardType, Score: 7}
	experienceReward = &Reward{Type: experienceRewardType, Score: 10}
	friendshipReward = &Reward{Type: friendshipRewardType, Score: 10}
)

type Reward struct {
	Type  string `json:"type"`
	Score int    `json:"score"`
}
