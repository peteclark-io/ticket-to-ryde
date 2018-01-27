package sprites

import (
	"image"
	_ "image/png" // for png
	"os"

	"github.com/faiface/pixel"
)

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

func BayMap() *pixel.Sprite {
	pic, err := loadPicture("./assets/the-bay.png")
	if err != nil {
		panic(err)
	}

	return pixel.NewSprite(pic, pic.Bounds())
}

func Logo() *pixel.Sprite {
	pic, err := loadPicture("./assets/logo.png")
	if err != nil {
		panic(err)
	}

	return pixel.NewSprite(pic, pic.Bounds())
}

func PlainLegoFigure() *pixel.Sprite {
	pic, err := loadPicture("./assets/plain-lego-man.png")
	if err != nil {
		panic(err)
	}

	return pixel.NewSprite(pic, pic.Bounds())
}
