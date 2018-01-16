package draw

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/peteclark-io/ticket-to-ryde/game"
)

func DrawMap(tgt pixel.Target, board *game.Board) {
	imd := imdraw.New(nil)

	for _, coord := range board.Coordinates {
		drawCoordinate(imd, coord)
	}

	imd.Draw(tgt)
}

func drawCoordinate(imd *imdraw.IMDraw, coord *game.Coordinate) {
	imd.Push(pixel.V(coord.X, coord.Y))
	imd.Circle(4, 2)
}
