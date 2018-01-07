package main

import (
	"context"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/peteclark-io/ticket-to-ryde/dimensions"
	"github.com/peteclark-io/ticket-to-ryde/scenes"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Ticket to Ryde",
		Bounds: pixel.R(0, 0, dimensions.WindowWidth, dimensions.WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	i := scenes.IntroScene{}
	i.RunScene(context.TODO(), win)

	g := scenes.GameScene{}
	g.RunScene(win)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func main() {
	pixelgl.Run(run)
}
