package scenes

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

type IntroScene struct {
}

func (i IntroScene) RunScene(ctx context.Context, win *pixelgl.Window) {
	win.Clear(colornames.Black)

	titleTxt := text.New(dimensions.Center(400).Add(pixel.V(0, 100)), vars.DefaultAtlas)

	fmt.Fprintf(titleTxt, strings.ToUpper("Ticket to Ryde"))
	titleTxt.Draw(win, pixel.IM.Scaled(titleTxt.Orig, 4))

	pressEnterTxt := text.New(dimensions.Center(100), vars.DefaultAtlas)
	fmt.Fprintf(pressEnterTxt, "Press Enter")
	pressEnterTxt.Draw(win, pixel.IM)

	for !win.Closed() {
		if ctx.Err() != nil {
			return
		}

		if win.JustPressed(pixelgl.KeyEnter) {
			break
		}

		win.Update()
	}

}
