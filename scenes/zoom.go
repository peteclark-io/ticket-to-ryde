package scenes

import "github.com/faiface/pixel"

type zoom struct {
	camZoom           float64
	zoomSpeed         float64
	lastScroll        pixel.Vec
	lastMousePosition pixel.Vec
}

func defaultZoom() *zoom {
	return &zoom{camZoom: 1.0, zoomSpeed: 0.5, lastScroll: pixel.ZV, lastMousePosition: pixel.ZV}
}

func (z *zoom) getLastMousePosition(cam pixel.Matrix) pixel.Vec {
	return cam.Unproject(z.lastMousePosition)
}
