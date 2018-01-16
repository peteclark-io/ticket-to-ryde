package dimensions

import "github.com/faiface/pixel"

const (
	WindowWidth  = 1024.0
	WindowHeight = 600.0
)

func CenterWidth(itemWidth int, y int) pixel.Vec {
	x := (WindowWidth / 2.0) - (float64(itemWidth) / 2.0)
	return pixel.V(x, float64(y))
}

func CenterHeight(x int) pixel.Vec {
	y := (WindowHeight / 2.0)
	return pixel.V(float64(x), y)
}

func Center(itemWidth int) pixel.Vec {
	x := (WindowWidth / 2.0) - (float64(itemWidth) / 2.0)
	y := (WindowHeight / 2.0)
	return pixel.V(x, y)
}

func TopRightCorner(itemWidth float64, margin float64) pixel.Vec {
	return pixel.V(WindowWidth-itemWidth-margin, WindowHeight-20)
}
