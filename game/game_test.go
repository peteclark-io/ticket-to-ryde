package game

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// func TestGame(t *testing.T) {
// 	p := NewPlayer("1", "Pedro")
// 	g := NewGame(BasicBoard, p)
//
// 	r := g.Start()
//
// 	for !r.Complete() {
// 		turn := r.NextTurn()
// 		for !turn.Complete() {
// 			position := g.GetPosition(turn.PlayerID)
// 			printToScreen(position)
//
// 			choices := g.Board.GetChoicesForPosition(position)
// 			printToScreen(choices)
//
// 			choices = turn.FilterChoices(choices)
// 			printToScreen(choices)
//
// 			choice := choices[len(choices)-1]
//
// 			switch choice.Type {
// 			case MovementChoiceType:
// 				pos := g.Board.Move(position, choice, turn)
// 				printToScreen(pos)
//
// 			case ActivityChoiceType:
// 				activity := g.Board.GetActivity(choice.ActivityID)
// 				rewards := activity.perform(turn)
// 				g.UpdateScore(turn.PlayerID, rewards)
//
// 			case PickupCardChoiceType:
// 				//TODO:
// 			case PlayCardChoiceType:
// 				//TODO:
// 			}
//
// 			turn.Spend(choice.Cost)
// 		}
// 	}
//
// 	printToScreen(g.Scores)
// }

func printToScreen(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "  ")
	log.Info(string(b))
}
