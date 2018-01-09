package scenes

import (
	"context"

	"github.com/faiface/pixel/pixelgl"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"golang.org/x/image/colornames"
)

// GameScene does game
func GameScene() Scene {
	return func(ctx context.Context, win *pixelgl.Window) {
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

			win.Update()
		}
	}
}
