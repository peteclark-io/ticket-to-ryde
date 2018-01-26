package draw

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

type CoordSelection struct {
	Rect  pixel.Rect
	Coord *game.Coordinate
}

func DrawMap(tgt pixel.Target, board *game.Board) []CoordSelection {
	imd := imdraw.New(nil)

	var coords []CoordSelection
	for _, coord := range board.Coordinates {
		coords = append(coords, drawCoordinate(imd, coord))
		drawLabel(tgt, board.GetActivity(coord.ActivityID).Name, coord)
	}

	for _, conn := range board.Connections {
		drawConnections(tgt, imd, conn.Distance, board.Coordinates[conn.Origin], board.Coordinates[conn.Destination])
	}

	imd.Draw(tgt)
	return coords
}

func drawConnections(tgt pixel.Target, imd *imdraw.IMDraw, distance int, a, b *game.Coordinate) {
	aV := pixel.V(a.X, a.Y)
	bV := pixel.V(b.X, b.Y)

	sum := aV.Add(bV).Map(func(arg1 float64) float64 {
		return arg1 / 2
	})

	dir := bV.Sub(sum).Unit().Rotated(-math.Pi / 2)

	basicTxt := text.New(sum.Add(dir.Scaled(scaleFactor(dir))), vars.DefaultAtlas)

	fmt.Fprint(basicTxt, fmt.Sprintf("%vAP", distance))
	basicTxt.Draw(tgt, pixel.IM)

	imd.Push(aV, bV)
	imd.Line(2)
}

func scaleFactor(v pixel.Vec) float64 {
	if v.X > 0 && v.Y > 0 {
		return 6.0
	} else if v.X < 0 && v.Y > 0 {
		return 18.0
	} else if v.X < 0 && v.Y < 0 {
		return 10.0
	}
	return 14.0
}

func drawCoordinate(imd *imdraw.IMDraw, coord *game.Coordinate) CoordSelection {
	imd.Push(pixel.V(coord.X, coord.Y))
	imd.Circle(10, 2)
	return CoordSelection{Rect: pixel.R(coord.X-6, coord.Y-6, coord.X+6, coord.Y+6), Coord: coord}
}

func drawLabel(tgt pixel.Target, label string, coord *game.Coordinate) {
	basicTxt := text.New(pixel.V(coord.X+14, coord.Y), vars.DefaultAtlas)
	fmt.Fprint(basicTxt, label)

	basicTxt.Draw(tgt, pixel.IM)
}
