package draw

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

func DrawScores(win *pixelgl.Window, g *game.Game) {
	txt := ""
	maxLength := 0
	for _, s := range g.Scores {
		name := g.GetPlayer(s.PlayerID).Name
		score := s.TotalScore()

		txt += fmt.Sprintf("%s: %v HP\n", name, score)
		if l := len(txt); l > maxLength {
			maxLength = l
		}
	}

	scoresTxt := text.New(dimensions.TopRightCorner(float64(maxLength)*4, 5.0), vars.DefaultAtlas)
	fmt.Fprint(scoresTxt, txt)

	scoresTxt.Draw(win, pixel.IM)
}
