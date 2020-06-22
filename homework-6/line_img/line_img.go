package line_img

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red color.Color = color.RGBA{200, 30, 30, 255}
var black color.Color = color.RGBA{0, 0, 0, 255}

type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}


func LineImg() {
	file, err := os.Create("line_img.png")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)

	// Рисуем линии
	for x := 20; x < 480; x++ {
		y := 50
		img.Set(x, y, black)
	}
	for y := 20; y < 480; y++ {
		x := 50
		img.Set(x, y, black)
	}
	for x := 20; x < 480; x++ {
		y := 450
		img.Set(x, y, black)
	}
	for y := 20; y < 480; y++ {
		x := 450
		img.Set(x, y, black)
	}
	for x := 50; x < 450; x++ {
		y := x
		img.Set(x, y, red)
	}
	for x := 50; x < 450; x++ {
		y := 500 - x
		img.Set(x, y, red)
	}

	png.Encode(file, img)

}
