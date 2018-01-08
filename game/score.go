package game

type Score struct {
	Total          int    `json:"total"`
	PlayerID       string `json:"player_id"`
	FoodRewards    int    `json:"food_rewards"`
	AlcoholRewards int    `json:"alcohol_rewards"`
	Experiences    int    `json:"experiences"`
}

func (s *Score) TotalScore() int {
	val := s.FoodRewards * 5
	val += s.AlcoholRewards * 7
	val += s.Experiences * 10
	s.Total = val
	return s.Total
}

func (s *Score) addRewards(rewards []*Reward) {
	for _, reward := range rewards {
		switch reward.Type {
		case foodRewardType:
			s.FoodRewards++
		case alcoholRewardType:
			s.AlcoholRewards++
		case experienceRewardType:
			s.Experiences++
		}
	}
}
