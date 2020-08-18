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
