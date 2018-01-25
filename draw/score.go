package draw

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

func DrawScores(d dimensions.D, tgt pixel.Target, g *game.Game) {
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

	scoresTxt := text.New(d.TopRightCorner(float64(maxLength)*5, 10.0), vars.DefaultAtlas)
	fmt.Fprint(scoresTxt, txt)

	scoresTxt.Draw(tgt, pixel.IM)
}
