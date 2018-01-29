package scenes

import "github.com/faiface/pixel"

type zoom struct {
	camZoom           float64
	zoomSpeed         float64
	lastScroll        pixel.Vec
	lastMousePosition pixel.Vec
}

func defaultZoom() *zoom {
	return &zoom{camZoom: 1.0, zoomSpeed: 3.0, lastScroll: pixel.ZV, lastMousePosition: pixel.ZV}
}
