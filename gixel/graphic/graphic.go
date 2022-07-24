package graphic

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GxlGraphic struct {
	img *ebiten.Image
}

func MakeGraphic(w, h int, c color.Color) *GxlGraphic {
	img := ebiten.NewImage(w, h)
	img.Fill(c)

	return &GxlGraphic{
		img: img,
	}
}

func LoadGraphic(path string) *GxlGraphic {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
		// TODO: Error handling, default value?
	}

	return &GxlGraphic{
		img: img,
	}
}

func (g *GxlGraphic) GetImage() *ebiten.Image {
	return g.img
}