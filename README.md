# ximgy
*Image handling utility in Go*

# Example
Download with
```sh
go get github.com/ximfect/ximgy
```

Example script
```go
package main

import (
	"fmt"
	"image/color"

	"github.com/ximfect/ximgy"
)

func invert(pixel ximgy.Pixel) color.RGBA {
	return color.RGBA{255 - pixel.R, 255 - pixel.G, 255 - pixel.B, pixel.A}
}

func main() {
	img, err := ximgy.Open("image.png")
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
```

# How to use
You can load an image with `ximgy.Open`:
```go
img, err := ximgy.Open("image.png")
```
By default, PNG and JPEG are supported. `Open` returns a `ximgy.Image` pointer.
You get individual pixels of the image by using `At`:
```go
pixel := img.At(10, 20)
```
And you can replace pixels with `Set`:
```go
newColor := color.RGBA{255, 255, 255, 255}
img.Set(15, 25, newColor)
```
**NOTE:** `At` and `Set` affect different instances of an `image.RGBA`. 
That means: if you use `Set` to change a pixel at x:10/y:10, `At` will not reflect that change.

You can save your image later using `ximgy.Save`, and again -- only PNG and JPEG are supported.
```go
ximgy.Save(img, "my_new_image.png")
```
Another useful function the `ximgy.Image` provides is `Iterate`.
It allows us to run a function on every pixel of the image, and take the function's output as the color of the pixel.
```go
// Error handling is ignored in this example, but you should always handle your errors

func invert(pixel ximgy.Pixel) color.RGBA {
    return color.RGBA{255 - pixel.R, 255 - pixel.G, 255 - pixel.B, pixel.A}
}

func main() {
    img, _ := ximgy.Open("my_image.jpeg")
    img.Iterate(invert)
    ximgy.Save(img, "my_new_image.jpeg")
}
```