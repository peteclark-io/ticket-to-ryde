package draw

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
	"golang.org/x/image/colornames"
)

var Board *game.Board

const height = 70.0
const width = 180.0
const margin = 10.0

func init() {
	Board = game.BasicBoard
}

type ChoiceSelection struct {
	Rect   pixel.Rect
	Choice *game.Choice
}

func DrawChoice(win *pixelgl.Window, c *game.Choice, index int) *ChoiceSelection {
	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.EndShape = imdraw.SharpEndShape

	yMin := (700 - height) // - (height * float64(index)) + margin
	yMax := 700.0          //- (height * float64(index))

	xMin := 50 + (width*float64(index) + margin)
	xMax := 50 + width + (width * float64(index))

	r := pixel.R(xMin, yMin, xMax, yMax)

	imd.Push(pixel.V(xMin, yMin), pixel.V(xMin, yMax), pixel.V(xMax, yMax), pixel.V(xMax, yMin), pixel.V(xMin, yMin))
	imd.Line(3)

	basicTxt := text.New(pixel.V(xMin+10, yMax-20), vars.DefaultAtlas)

	switch c.Type {
	case game.MovementChoiceType:
		fmt.Fprintf(basicTxt, "Go to %s", Board.Activities[c.ActivityID].Name)
	case game.ActivityChoiceType:
		fmt.Fprintf(basicTxt, "Go inside %s", Board.Activities[c.ActivityID].Name)
	case game.PickupCardChoiceType:
		fmt.Fprintf(basicTxt, "Pick up a card?")
	case game.PlayCardChoiceType:
		fmt.Fprintf(basicTxt, "Play a card?")
	case game.FinishTurnChoiceType:
		fmt.Fprintf(basicTxt, "Finish turn?")
	}

	if c.Cost != 0 {
		fmt.Fprintf(basicTxt, "\n\nCosts: %vAP", c.Cost)
	}

	basicTxt.Draw(win, pixel.IM)
	imd.Draw(win)

	return &ChoiceSelection{Rect: r, Choice: c}
}
