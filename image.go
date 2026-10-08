package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func LoadImgage(path string) (*image.NRGBA, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Error opning file")
	}

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("Error in decoding image")
	}

	rgba := img.(*image.NRGBA)

	color := rgba.NRGBAAt(0, 0)
	fmt.Println(color.R)
	fmt.Println(color.B)
	fmt.Println(color.G)
	fmt.Println(color.A)

	return rgba, nil

}

func ImgCapacity(*image.NRGBA)(size byte){

	return 0
}

