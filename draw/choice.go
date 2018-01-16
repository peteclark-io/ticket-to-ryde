package draw

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
	"golang.org/x/image/colornames"
)

type ChoiceSelection struct {
	Rect   pixel.Rect
	Choice *game.Choice
}

type ChoiceConfig struct {
	Height  float64
	Width   float64
	XMargin float64
	YMargin float64
}

func NewChoiceConfig(height, width, xMargin, yMargin float64) *ChoiceConfig {
	return &ChoiceConfig{Height: height, Width: width, XMargin: xMargin, YMargin: yMargin}
}

func NewDefaultChoiceConfig() *ChoiceConfig {
	return &ChoiceConfig{Height: 70.0, Width: 180.0, XMargin: 10.0, YMargin: 60.0}
}

func DrawChoice(conf *ChoiceConfig, win *pixelgl.Window, board *game.Board, c *game.Choice, index int) *ChoiceSelection {
	if conf == nil {
		conf = NewDefaultChoiceConfig()
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.White
	imd.EndShape = imdraw.SharpEndShape

	yMin := (dimensions.WindowHeight - conf.Height - conf.YMargin) // - (height * float64(index)) + margin
	yMax := dimensions.WindowHeight - conf.YMargin                 //- (height * float64(index))

	xMin := 50 + (conf.Width*float64(index) + conf.XMargin)
	if index == 0 {
		xMin = 50 + (conf.Width * float64(index))
	}

	xMax := 50 + conf.Width + (conf.Width * float64(index))

	r := pixel.R(xMin, yMin, xMax, yMax)

	imd.Push(pixel.V(xMin, yMin), pixel.V(xMin, yMax), pixel.V(xMax, yMax), pixel.V(xMax, yMin), pixel.V(xMin, yMin))
	imd.Line(3)

	basicTxt := text.New(pixel.V(xMin+10, yMax-20), vars.DefaultAtlas)

	switch c.Type {
	case game.MovementChoiceType:
		fmt.Fprintf(basicTxt, "Go to %s", board.Activities[c.ActivityID].Name)
	case game.ActivityChoiceType:
		fmt.Fprintf(basicTxt, "Go inside %s", board.Activities[c.ActivityID].Name)
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
