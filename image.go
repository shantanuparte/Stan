package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func LoadImgage(path string) (*image.NRGBA, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Error opning file")
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("Error in decoding image")
	}

	rgba := image.NewNRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	color := rgba.NRGBAAt(0, 0)
	fmt.Println(color.R)
	fmt.Println(color.B)
	fmt.Println(color.G)
	fmt.Println(color.A)

	return rgba, nil

}

func ImgCapacity(img *image.NRGBA) int {

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	value := (width * height * 3 * 2) / 8

	return value

}

func LoadHideFile(path string) ([]byte, error) {
	data_stream, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Error in opning hiding file")
	}

	return data_stream, nil

}
func HideCapacity(path string) (int, error) {

	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("Error in opning file")
	}

	return int(info.Size()), nil

}

func CheckIfItFits(img_bytes int, hide_file_bytes int) bool {
	if hide_file_bytes > img_bytes {
		return false
	}

	return true
}
