package main

import (
	"image"
)

func Encoding(img *image.NRGBA, data []byte) {

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	data_index := 0
	bit_index := 0

	for y := 0; y < height; y++ {

		for x := 0; x < width; x++ {

			pixel := img.NRGBAAt(x, y)
			rgba_channel := []*uint8{&pixel.R, &pixel.G, &pixel.B}

			for _, cha := range rgba_channel {
				//chan is actuall channell value
				bit := (data[data_index] >> (6 - bit_index) & 0b11)

				idk := *cha &^ 0b11
				*cha = idk | bit
				bit_index += 2

				if bit_index == 8 {
					bit_index = 0
					data_index++
				}

			}

		}

		if data_index > len(data) {
			break
		}
	}

}
