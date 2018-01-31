package scenes

import (
	"context"
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/draw"
	"github.com/peteclark-io/ticket-to-ryde/fps"
	"github.com/peteclark-io/ticket-to-ryde/game"
	"github.com/peteclark-io/ticket-to-ryde/sprites"
	"github.com/peteclark-io/ticket-to-ryde/vars"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/colornames"
)

func TurnScene(g *game.Game, turn *game.Turn) Scene {
	return func(ctx context.Context, win *pixelgl.Window) {
		screen := pixel.R(-1024, -600, 1024, 600)

		camPos := pixel.V(-800, -500)
		canvas := pixelgl.NewCanvas(screen)
		figure := sprites.PlainLegoFigure()
		bayMap := sprites.BayMap()

		zoom := defaultZoom()

		for !win.Closed() {
			win.Clear(colornames.Forestgreen)

			canvas.Clear(pixel.Alpha(0))
			bayMap.Draw(win, pixel.IM.Scaled(pixel.ZV, win.Bounds().H()/bayMap.Frame().H()).Moved(win.Bounds().Center()))

			if turn.Complete() {
				break
			}

			position := g.GetPosition(turn.PlayerID)
			activity := g.Board.GetActivity(position.LastActivity)
			player := g.GetPlayer(turn.PlayerID)

			basicTxt := text.New(pixel.V(50, dimensions.WindowHeight-20), vars.DefaultAtlas)

			fmt.Fprintf(basicTxt, "%s you're in %s - you have %vAP remaining", player.Name, activity.Name, turn.RemainingAP)
			basicTxt.Draw(win, pixel.IM)

			choices := g.Board.GetChoicesForPosition(position)
			choices = turn.FilterChoices(choices)

			choicePositions := make([]*draw.ChoiceSelection, 0)
			for i, c := range choices {
				r := draw.DrawChoice(nil, win, g.Board, c, i)
				choicePositions = append(choicePositions, r)
			}

			draw.DrawScores(dimensions.DefaultScreen, win, g)
			coords := draw.DrawMap(canvas, g.Board)

			if win.JustPressed(pixelgl.MouseButtonLeft) {
				for _, p := range choicePositions {
					if p.Rect.Contains(win.MousePosition()) {
						log.Infof("Selected %s", p.Choice.Type)
						g.Select(position, turn, p.Choice)
					}
				}

				for _, c := range coords {
					if c.Rect.Contains(win.MousePosition()) {
						log.Infof("Selected %s", g.Board.GetActivity(c.Coord.ActivityID))
					}
				}
			}

			speed := 8.0
			if win.Pressed(pixelgl.KeyUp) {
				camPos.Y = camPos.Y + speed
			}

			if win.Pressed(pixelgl.KeyDown) {
				camPos.Y = camPos.Y - speed
			}

			if win.Pressed(pixelgl.KeyRight) {
				camPos.X = camPos.X + speed
			}

			if win.Pressed(pixelgl.KeyLeft) {
				camPos.X = camPos.X - speed
			}

			if zoom.lastScroll != win.MouseScroll() {
				zoom.camZoom *= math.Pow(zoom.zoomSpeed, win.MouseScroll().Y)
				zoom.lastMousePosition = win.MousePosition()
				zoom.lastScroll = win.MouseScroll()
			}

			cam := pixel.IM.Scaled(zoom.getLastMousePosition(pixel.IM.Moved(camPos.Scaled(-1))), zoom.camZoom).Moved(camPos.Scaled(-1))
			canvas.SetMatrix(cam)

			for _, c := range coords {
				if c.Rect.Contains(cam.Unproject(win.MousePosition())) {
					draw.DrawCoordinateLabel(canvas, g.Board, c.Coord)
					break
				}
			}

			figure.Draw(canvas, pixel.IM.
				Moved(g.Board.Coordinates[position.LastActivity].ToVector()))

			canvas.Draw(win, pixel.IM.
				Moved(canvas.Bounds().Center()))

			win.Update()
			fps.MeasureFPS(win)
		}
	}
}
