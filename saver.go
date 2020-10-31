package ximgy

import (
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	// "golang.org/x/image/webp"
)

// GetFilenameFormat returns the format from the filename
func GetFilenameFormat(filename string) string {
	filenameSplit := strings.Split(filename, ".")
	filenameExt := filenameSplit[len(filenameSplit)-1]

	switch filenameExt {
	case "png":
		return "png"
	case "jpeg", "jpg", "jpe", "jfif":
		return "jpeg"
	case "bmp", "dib":
		return "bmp"
	case "tiff", "tif":
		return "tiff"
	case "webp":
		// unused; webp encoding currently unsupported
		return "webp"
	default:
		return "..."
	}
}

// Save saves an image under the given filename
func Save(img *Image, filename string) error {
	format := GetFilenameFormat(filename)
	if format == "..." {
		return errors.New("unknown format")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	switch format {
	case "png":
		err = png.Encode(file, img.GetOutput())
		if err != nil {
			return err
		}
	case "jpeg":
		err = jpeg.Encode(file, img.GetOutput(), &jpeg.Options{Quality: 80})
		if err != nil {
			return err
		}
	case "bmp":
		err = bmp.Encode(file, img.GetOutput())
		if err != nil {
			return err
		}
	case "tiff":
		err = tiff.Encode(file, img.GetOutput(), &tiff.Options{})
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	return nil
}
