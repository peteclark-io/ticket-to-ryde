package draw

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/vars"
)

func DrawMap(tgt pixel.Target, board *game.Board) {
	imd := imdraw.New(nil)

	for _, coord := range board.Coordinates {
		drawCoordinate(imd, coord)
		drawLabel(tgt, board.GetActivity(coord.ActivityID).Name, coord)
	}

	for _, conn := range board.Connections {
		drawConnections(tgt, imd, conn.Distance, board.Coordinates[conn.Origin], board.Coordinates[conn.Destination])
	}

	imd.Draw(tgt)
}

func drawConnections(tgt pixel.Target, imd *imdraw.IMDraw, distance int, a, b *game.Coordinate) {
	aV := pixel.V(a.X, a.Y)
	bV := pixel.V(b.X, b.Y)

	// sum := aV.Add(bV).Map(func(arg1 float64) float64 {
	// 	return arg1 / 2
	// })

	basicTxt := text.New(aV.Normal(), vars.DefaultAtlas)
	fmt.Fprint(basicTxt, distance)
	basicTxt.Draw(tgt, pixel.IM)

	imd.Push(aV, bV)
	imd.Line(2)
}

func drawCoordinate(imd *imdraw.IMDraw, coord *game.Coordinate) {
	imd.Push(pixel.V(coord.X, coord.Y))
	imd.Circle(10, 2)
}

func drawLabel(tgt pixel.Target, label string, coord *game.Coordinate) {
	basicTxt := text.New(pixel.V(coord.X+14, coord.Y), vars.DefaultAtlas)
	fmt.Fprint(basicTxt, label)

	basicTxt.Draw(tgt, pixel.IM)
}
