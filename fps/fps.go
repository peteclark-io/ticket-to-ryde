package fps

import (
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

var frames = 0
var tick = time.Tick(1 * time.Second)

func MeasureFPS(win *pixelgl.Window) {
	frames++
	select {
	case <-tick:
		win.SetTitle(fmt.Sprintf("Ticket To Ryde | FPS: %d", frames))
		frames = 0
	default:
	}
}
