package scenes

import (
	"context"

	"github.com/faiface/pixel/pixelgl"
	"github.com/peteclark-io/ticket-to-ryde/fps"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/transitions"
	"golang.org/x/image/colornames"
)

// GameScene does game
func GameScene() Scene {
	return func(ctx context.Context, win *pixelgl.Window) {
		p := game.NewPlayer("1", "Pedro")
		d := game.NewPlayer("2", "Fernando Luiz Roza")
		g := game.NewGame(game.BasicBoard, p, d)

		round := g.Start()
		for !win.Closed() {
			win.Clear(colornames.Black)

			if round.Complete() {
				round = g.NextRound(round)

				if round.Time == "day" {
					transitions.Sunrise(ctx, win)
				} else {
					transitions.NightTime(ctx, win)
				}
			}

			RoundScene(g, round)(ctx, win)

			win.Update()
			fps.MeasureFPS(win)
		}
	}
}
