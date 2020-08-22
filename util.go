package ximgy

import (
	"image"
	"image/draw"
)

// MakeRGBA returns an RGBA version of the source image
func MakeRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	output := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(output, output.Bounds(), src, bounds.Min, draw.Src)
	return output
}

// ClampCoords ensures x and y fit within size
func ClampCoords(x, y int, size image.Point, underflow bool) (int, int) {
	if underflow {
		for x < 0 {
			x += size.X
		}
		for y < 0 {
			y += size.Y
		}
		for x >= size.X {
			x -= size.X
		}
		for y >= size.Y {
			y -= size.Y
		}
	} else {
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
		if x >= size.X {
			x = size.X - 1
		}
		if y >= size.Y {
			y = size.Y - 1
		}
	}
	return x, y
}
