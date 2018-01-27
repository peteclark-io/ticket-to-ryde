package draw

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"golang.org/x/image/colornames"
)

func drawActivity(imd *imdraw.IMDraw, coord *game.Coordinate) {
	rect := rectangleAroundVector(coord.ToVector(), 40, 60)
	imd.Color = colornames.Midnightblue
	imd.Push(rect...)
	imd.Polygon(0)

	imd.Color = colornames.White
	imd.Push(rect...)
	imd.Rectangle(2)
}

func rectangleAroundVector(v pixel.Vec, height, width float64) []pixel.Vec {
	var results []pixel.Vec

	results = append(results, v.Add(pixel.V(-width/2, -height/2)))
	results = append(results, v.Add(pixel.V(-width/2, height/2)))
	results = append(results, v.Add(pixel.V(width/2, height/2)))
	results = append(results, v.Add(pixel.V(width/2, -height/2)))
	results = append(results, v.Add(pixel.V(-width/2, -height/2)))

	return results
}
