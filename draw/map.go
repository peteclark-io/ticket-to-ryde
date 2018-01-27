package draw

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
	"golang.org/x/image/colornames"
)

type CoordSelection struct {
	Rect  pixel.Rect
	Coord *game.Coordinate
}

func DrawMap(tgt pixel.Target, board *game.Board) []CoordSelection {
	imd := imdraw.New(nil)

	for _, conn := range board.Connections {
		drawConnections(tgt, imd, conn.Distance, board.Coordinates[conn.Origin], board.Coordinates[conn.Destination])
	}

	var coords []CoordSelection
	for _, coord := range board.Coordinates {
		// coords = append(coords, )
		drawActivity(imd, coord)
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
	basicTxt.Color = colornames.White

	fmt.Fprint(basicTxt, fmt.Sprintf("%vAP", distance))
	basicTxt.Draw(tgt, pixel.IM)

	imd.Color = colornames.White
	imd.Push(aV, bV)
	imd.Line(1)
}

func scaleFactor(v pixel.Vec) float64 {
	perc := math.Abs(v.Angle() / math.Pi) // deal with the extreme angles
	if perc > 0.95 || perc < 0.05 {
		return 16.0
	} else if perc > 0.85 || perc < 0.15 {
		return 14.0
	}

	if perc > 0.48 && perc < 0.52 {
		return 5.0
	}

	if v.X > 0 && v.Y > 0 {
		return 6.0
	} else if v.X < 0 && v.Y > 0 {
		return 10.0
	} else if v.X < 0 && v.Y < 0 {
		return 10.0
	}
	return 10.0
}

func drawCoordinate(imd *imdraw.IMDraw, coord *game.Coordinate) CoordSelection {
	imd.Push(pixel.V(coord.X, coord.Y))
	imd.Circle(6, 0)
	offset := 9.5
	return CoordSelection{Rect: pixel.R(coord.X-offset, coord.Y-offset, coord.X+offset, coord.Y+offset), Coord: coord}
}

func DrawCoordinateLabel(tgt pixel.Target, board *game.Board, coord *game.Coordinate) {
	label := board.GetActivity(coord.ActivityID).Name

	basicTxt := text.New(pixel.V(coord.X+14, coord.Y), vars.DefaultAtlas)
	fmt.Fprint(basicTxt, label)

	basicTxt.Draw(tgt, pixel.IM)
}
