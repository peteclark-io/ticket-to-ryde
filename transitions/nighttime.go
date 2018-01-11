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

func NightTime(ctx context.Context, win *pixelgl.Window) {
	win.Clear(colornames.Black)

	titleTxt := text.New(dimensions.Center(11*30).Add(pixel.V(0, 100)), vars.DefaultAtlas)

	fmt.Fprintf(titleTxt, strings.ToUpper("Night Time!"))
	titleTxt.Draw(win, pixel.IM.Scaled(titleTxt.Orig, 4))

	dt := time.Now()

	for !win.Closed() {
		if time.Now().After(dt.Add(time.Second * 3)) {
			break
		}

		win.Update()
	}
}
