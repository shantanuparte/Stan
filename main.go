package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
)

func main() {
	fmt.Println("WELCOME TO STAN")

	mode := flag.String("mode", "encode", "Encode or decode?")
	file_path := flag.String("path", "", "The image to hide") //pointer return karlay
	hide_path := flag.String("hide", "", "The file to hide in image")
	output_path := flag.String("output", "output.png", "The name of the output file")

	flag.Parse()

	switch *mode {
	case "encode":
		if *hide_path == "" || *file_path == "" {
			fmt.Println("You should fill the \"path\" and \"hide\" value")
			return
		}
		handleEncodee(*hide_path, *file_path, *output_path)

	case "decode":
		if *file_path == "" {
			fmt.Println("The file")
		}
		hanldeDeocde(*file_path, *output_path)

	}

}

func hanldeDeocde(file_path, output_path string) {

	img, err := LoadImgage(file_path)
	if err != nil {
		fmt.Printf("Error loading image: %v\n",err)
		return
	}

	fmt.Printf("Decoding image from %s...\n", file_path)

	extracted , err := Decoder(img)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile(output_path, extracted, 0644)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		return
	}

	fmt.Printf("Successfully decoded! File saved to %s\n", output_path)

}

func handleEncodee(file_path, hide_path, output_path string) {
	img, err := LoadImgage(file_path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	img_size := ImgCapacity(img)

	hide_file_stream, err := LoadHideFile(hide_path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return

	}

	hide_capacity, err := HideCapacity(hide_path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	can := CheckIfItFits(img_size, hide_capacity)
	if can != true {
		fmt.Println("It Doesn't fit in the image")
		return
	}

	fmt.Println("This works")

	fmt.Println(hide_file_stream)
	fmt.Println(string(hide_file_stream))

	output, err := os.Create(output_path)
	if err != nil {
		return
	}
	defer output.Close()

	Encoding(img, hide_file_stream)

	if err = png.Encode(output, img); err != nil {
		fmt.Println("Error in image encoding")
		return
	}
}
