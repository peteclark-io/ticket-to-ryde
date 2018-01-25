package dimensions

import "github.com/faiface/pixel"

const (
	WindowWidth  = 1024.0
	WindowHeight = 600.0
)

var DefaultScreen = D{pixel.R(0, 0, WindowWidth, WindowHeight)}

type D struct {
	pixel.Rect
}

func (d D) CenterWidth(itemWidth int, y int) pixel.Vec {
	x := (d.Max.X / 2.0) - (float64(itemWidth) / 2.0)
	return pixel.V(x, float64(y))
}

func (d D) CenterHeight(x int) pixel.Vec {
	y := (d.Max.Y / 2.0)
	return pixel.V(float64(x), y)
}

func (d D) CenterInScreen(itemWidth int) pixel.Vec {
	x := (d.Max.X / 2.0) - (float64(itemWidth) / 2.0)
	y := (d.Max.Y / 2.0)
	return pixel.V(x, y)
}

func (d D) TopRightCorner(itemWidth float64, margin float64) pixel.Vec {
	return pixel.V(d.Max.X-itemWidth-margin, d.Max.Y-20)
}
