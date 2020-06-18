package chess_board

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var teal color.Color = color.RGBA{0, 200, 200, 255}
var braun color.Color = color.RGBA{150, 75, 0, 255}
var black color.Color = color.RGBA{0, 0, 0, 255}

type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

func ChessBoard() {
	file, err := os.Create("chess_board.png")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.Point{X: 0, Y: 0}, draw.Src)

	// Рисуем шахматную доску
	// Рисуем чёрные квадраты
	mask := image.NewRGBA(image.Rect(0, 0, 0, 0))
	for x := 90; x < 340; x += 80 {
		for y := 90; y < 340; y += 80 {
			mask = image.NewRGBA(image.Rect(x, y, x+40, y+40))
			draw.Draw(mask, mask.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
			draw.DrawMask(img, img.Bounds(), mask, image.ZP, &image.Uniform{black}, image.ZP, draw.Src)

			mask = image.NewRGBA(image.Rect(x+40, y+40, x+80, y+80))
			draw.Draw(mask, mask.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
			draw.DrawMask(img, img.Bounds(), mask, image.ZP, &image.Uniform{black}, image.ZP, draw.Src)
		}

	}

	// Рисуем коричневые квадраты
	for x := 130; x < 380; x += 80 {
		for y := 90; y < 340; y += 80 {

			mask = image.NewRGBA(image.Rect(x, y, x+40, y+40))
			draw.Draw(mask, mask.Bounds(), &image.Uniform{braun}, image.ZP, draw.Src)
			draw.DrawMask(img, img.Bounds(), mask, image.ZP, &image.Uniform{black}, image.ZP, draw.Src)

			mask = image.NewRGBA(image.Rect(x-40, y+40, x, y+80))
			draw.Draw(mask, mask.Bounds(), &image.Uniform{braun}, image.ZP, draw.Src)
			draw.DrawMask(img, img.Bounds(), mask, image.ZP, &image.Uniform{black}, image.ZP, draw.Src)
		}

	}

	png.Encode(file, img)

}
