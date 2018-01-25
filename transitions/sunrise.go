package transitions

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/vars"
	"golang.org/x/image/colornames"
)

func Sunrise(ctx context.Context, win *pixelgl.Window) {
	win.Clear(colornames.Skyblue)

	titleTxt := text.New(dimensions.DefaultScreen.CenterInScreen(7*30).Add(pixel.V(0, 100)), vars.DefaultAtlas)

	fmt.Fprintf(titleTxt, strings.ToUpper("Dayime!"))
	titleTxt.Draw(win, pixel.IM.Scaled(titleTxt.Orig, 4))

	dt := time.Now()

	for !win.Closed() {
		if time.Now().After(dt.Add(time.Millisecond * 2200)) {
			break
		}

		win.Update()
	}
}
