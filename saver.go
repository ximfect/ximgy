package ximgy

import (
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// GetFilenameFormat returns the format from the filename
func GetFilenameFormat(filename string) string {
	filenameSplit := strings.Split(filename, ".")
	filenameExt := filenameSplit[len(filenameSplit)-1]

	switch filenameExt {
	case "png":
		return "png"
	case "jpeg":
		return "jpeg"
	case "jpg":
		return "jpeg"
	case "jfif":
		return "jpeg"
	}

	return "..."
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
	}

	return nil
}
