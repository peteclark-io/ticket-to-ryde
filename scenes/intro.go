package scenes

import (
	"context"
	"fmt"
	"math"
	"strings"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/fps"
	"github.com/peteclark-io/ticket-to-ryde/sprites"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

func IntroScene() Scene {
	return func(ctx context.Context, win *pixelgl.Window) {
		win.Clear(colornames.Black)

		logo := sprites.Logo()

		titleTxt := text.New(dimensions.DefaultScreen.CenterInScreen(400).Add(pixel.V(0, 100)), vars.DefaultAtlas)

		fmt.Fprintf(titleTxt, strings.ToUpper("Ticket to Ryde"))
		//titleTxt.Draw(win, pixel.IM.Scaled(titleTxt.Orig, 4))

		pressEnterTxt := text.New(pixel.ZV, vars.DefaultAtlas)
		pressEnterTxt.Color = colornames.Black
		fmt.Fprintf(pressEnterTxt, "Press Enter")

		logo.Draw(win, pixel.IM.Scaled(pixel.ZV,
			math.Min(
				win.Bounds().W()/logo.Frame().W(),
				win.Bounds().H()/logo.Frame().H(),
			),
		).Moved(win.Bounds().Center()))

		pressEnterTxt.Draw(win, pixel.IM.Scaled(pixel.ZV, 3.0).Moved(dimensions.DefaultScreen.CenterWidth(100, 100)))

		for !win.Closed() {
			if ctx.Err() != nil {
				return
			}

			if win.JustPressed(pixelgl.KeyEnter) {
				break
			}

			win.Update()
			fps.MeasureFPS(win)
		}

	}
}
