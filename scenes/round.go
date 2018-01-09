package scenes

import (
	"context"

	"github.com/faiface/pixel/pixelgl"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"golang.org/x/image/colornames"
)

func RoundScene(g *game.Game, round *game.Round) Scene {
	return func(ctx context.Context, win *pixelgl.Window) {
		for !win.Closed() {
			win.Clear(colornames.Black)

			if round.Complete() {
				break
			}

			turn := round.NextTurn()

			TurnScene(g, turn)(ctx, win)

			win.Update()
		}
	}
}
