package ximgy

import (
	"image"
	// PNG decoding
	_ "image/png"
	// JPEG decoding
	_ "image/jpeg"
	"os"

	// BMP decoding
	_ "golang.org/x/image/bmp"
	// TIFF decoding
	_ "golang.org/x/image/tiff"
	// WEBP decoding
	_ "golang.org/x/image/webp"
)

// Open opens, reads, decodes and constructs an Image
func Open(filename string) (*Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	data, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	dataRGBA, ok := data.(*image.RGBA)
	if !ok {
		dataRGBA = MakeRGBA(data)
	}

	bounds := dataRGBA.Bounds()
	empty := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	img := Image{dataRGBA, empty, bounds.Size()}

	return &img, nil
}

// MakeEmpty constructs and returns an empty Image
func MakeEmpty(r image.Rectangle) *Image {
	source := image.NewRGBA(r)
	output := image.NewRGBA(r)
	size := r.Bounds().Size()
	img := Image{source, output, size}

	return &img
}
