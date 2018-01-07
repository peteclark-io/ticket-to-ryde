package scenes

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/draw"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/colornames"
)

type GameScene struct {
}

func (i GameScene) RunScene(win *pixelgl.Window) {
	p := game.NewPlayer("1", "Pedro")
	d := game.NewPlayer("2", "Ducky")
	g := game.NewGame(game.BasicBoard, p, d)

	round := g.Start()
	turn := round.NextTurn()

	for !win.Closed() {
		win.Clear(colornames.Black)

		if round.Complete() {
			round = g.Start()
		}

		if turn.Complete() {
			turn = round.NextTurn()
		}

		position := g.GetPosition(turn.PlayerID)
		activity := g.Board.GetActivity(position.LastActivity)
		player := g.GetPlayer(turn.PlayerID)

		basicTxt := text.New(pixel.V(50, 750), vars.DefaultAtlas)

		fmt.Fprintf(basicTxt, "%s you're in %s - you have %vAP remaining", player.Name, activity.Name, turn.RemainingAP)
		basicTxt.Draw(win, pixel.IM)

		choices := g.Board.GetChoicesForPosition(position)
		choices = turn.FilterChoices(choices)

		choicePositions := make([]*draw.ChoiceSelection, 0)
		for i, c := range choices {
			r := draw.DrawChoice(win, c, i)
			choicePositions = append(choicePositions, r)
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			for _, p := range choicePositions {
				if p.Rect.Contains(win.MousePosition()) {
					log.Infof("Selected %s", p.Choice.Type)
					g.Select(position, turn, p.Choice)
				}
			}
		}
		win.Update()
	}
}
