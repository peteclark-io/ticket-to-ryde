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
		//d := game.NewPlayer("2", "Ducky")
		g := game.NewGame(game.BasicBoard, p)

		round := g.Start()
		for !win.Closed() {
			win.Clear(colornames.Black)

			if round.Complete() {
				break
			}

			RoundScene(g, round)(ctx, win)
			win.Update()
		}
	}
}
