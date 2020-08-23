package main

import (
	"fmt"
	"image/color"

	"github.com/ximfect/ximgy"
)

func invert(pixel ximgy.Pixel) (color.RGBA, error) {
	return color.RGBA{255 - pixel.R, 255 - pixel.G, 255 - pixel.B, pixel.A}, nil
}

func main() {
	img, err := ximgy.Open("EffzbrUXsAMKkkp.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("image loaded", img.Size)
	img.Iterate(invert)
	fmt.Println("image iterated")

	err = ximgy.Save(img, "out.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("image saved")
}
