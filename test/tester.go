package main

import (
	"fmt"

	"github.com/ximfect/ximgy"
)

func main() {
	img, err := ximgy.Open("EffzbrUXsAMKkkp.png")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("image loaded", img.Size)
	}
}
