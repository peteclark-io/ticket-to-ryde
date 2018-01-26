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
	"github.com/peteclark-io/ticket-to-ryde/vars"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/colornames"
)

func TurnScene(g *game.Game, turn *game.Turn) Scene {
	return func(ctx context.Context, win *pixelgl.Window) {
		screen := pixel.R(-512, -300, 512, 300)
		// d := dimensions.D{Rect: screen}

		camPos := pixel.ZV
		canvas := pixelgl.NewCanvas(screen)

		for !win.Closed() {
			win.Clear(colornames.Black)
			canvas.Clear(pixel.Alpha(0))

			if turn.Complete() {
				break
			}

			matrix := pixel.IM.Scaled(pixel.ZV,
				math.Min(
					win.Bounds().W()/canvas.Bounds().W(),
					win.Bounds().H()/canvas.Bounds().H(),
				),
			).Scaled(pixel.ZV, 2).Moved(win.Bounds().Center())

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
					if c.Rect.Contains(matrix.Unproject(win.MousePosition())) {
						log.Infof("Selected %s", g.Board.GetActivity(c.Coord.ActivityID))
					}
				}
			}

			if win.Pressed(pixelgl.KeyDown) {
				camPos.Y--
			}

			if win.Pressed(pixelgl.KeyUp) {
				camPos.Y++
			}

			if win.Pressed(pixelgl.KeyLeft) {
				camPos.X--
			}

			if win.Pressed(pixelgl.KeyRight) {
				camPos.X++
			}

			canvas.SetMatrix(pixel.IM.Moved(camPos.Scaled(-2)))

			canvas.Draw(win, matrix)

			win.Update()
			fps.MeasureFPS(win)
		}
	}
}
