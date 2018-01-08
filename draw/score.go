package draw

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

func DrawScores(win *pixelgl.Window, g *game.Game) {
	scoresTxt := text.New(pixel.V(900, 750), vars.DefaultAtlas)

	for _, s := range g.Scores {
		name := g.GetPlayer(s.PlayerID).Name
		score := s.TotalScore()

		fmt.Fprintf(scoresTxt, "%s: %v HP\n", name, score)
	}

	scoresTxt.Draw(win, pixel.IM)
}
