package main

import (
	"encoding/binary"
	"fmt"
	"image"
)

func Encoding(img *image.NRGBA, data []byte) {

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	str := []byte("STAN")

	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(len(data)))

	final_slice := make([]byte, 0, len(str)+len(size)+len(data))
	final_slice = append(final_slice, str...)
	final_slice = append(final_slice, size...)
	final_slice = append(final_slice, data...)

	data_index := 0
	bit_index := 0

	for y := 0; y < height; y++ {

		for x := 0; x < width; x++ {

			pixel := img.NRGBAAt(x, y)

			rgba_channel := []*uint8{&pixel.R, &pixel.G, &pixel.B}

			for _, cha := range rgba_channel {
				if data_index >= len(final_slice) {
					img.SetNRGBA(x, y, pixel)
					return
				}

				bit := (final_slice[data_index] >> (6 - bit_index) & 0b11)

				idk := *cha &^ 0b11
				*cha = idk | bit
				bit_index += 2

				if bit_index == 8 {
					bit_index = 0
					data_index++
				}

			}

			img.SetNRGBA(x, y, pixel)

		}

	}

}

func Decoder(img *image.NRGBA) ([]byte, error) {

	bounds := img.Bounds()

	height, width := bounds.Dy(), bounds.Dx()

	extracted := make([]byte, 0, height*width*3)
	var bite byte
	count := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			pixel := img.NRGBAAt(x, y)

			channel := []uint8{pixel.R, pixel.G, pixel.B}

			for _, value := range channel {
				bit := value & 0b11
				bite = (bite << 2) | bit
				count += 2
				if count == 8 {
					extracted = append(extracted, bite)
					count = 0
					bite = 0
				}

			}

		}
	}

	if len(extracted) < 8 {
		return nil, fmt.Errorf("Bro it's too low lenght lol")
	}

	if string(extracted[0:4]) != "STAN" {
		return nil, fmt.Errorf("Not stan")
	}

	size := binary.BigEndian.Uint32(extracted[4:8])

	if uint32(len(extracted)) < 8+size {
		return nil, fmt.Errorf("payload is incomplete")
	}

	return extracted[8:8+size], nil

}
