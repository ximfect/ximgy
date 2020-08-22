package ximgy

import (
	"image"
	"image/color"
)

// Image represents two copies of an Image, one read-only and one write-only
type Image struct {
	source *image.RGBA
	output *image.RGBA
	Size   image.Point
}

// Pixel represents a pixel used in Interate
type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
	X int
	Y int
}

// At returns a pixel from the source image
func (i *Image) At(x, y int) color.RGBA {
	x, y = ClampCoords(x, y, i.Size, true)
	return i.source.At(x, y).(color.RGBA)
}

// Set modifies a pixel in the output image
func (i *Image) Set(x, y int, v color.RGBA) {
	x, y = ClampCoords(x, y, i.Size, true)
	i.output.SetRGBA(x, y, v)
}

// Iterate calls the given function for every pixel in the source image,
// taking it's output as the new value for the pixel in the output image.
func (i *Image) Iterate(f func(Pixel) color.RGBA) {
	for y := 0; y < i.Size.Y; y++ {
		for x := 0; x < i.Size.X; x++ {
			value := i.At(x, y)
			r32, g32, b32, a32 := value.RGBA()
			r8 := uint8(r32 >> 8)
			g8 := uint8(g32 >> 8)
			b8 := uint8(b32 >> 8)
			a8 := uint8(a32 >> 8)
			pixel := Pixel{r8, g8, b8, a8, x, y}
			out := f(pixel)
			i.Set(x, y, out)
		}
	}
}

// SetSource sets the input image
func (i *Image) SetSource(src *image.RGBA) {
	i.source = src
}

// GetOutput returns the image that was (most likely) modified
func (i *Image) GetOutput() *image.RGBA {
	return i.output
}
