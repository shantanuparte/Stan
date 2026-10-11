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
	output_decode := flag.String("deout", "", "The file type")

	flag.Parse()

	switch *mode {
	case "encode":
		if *hide_path == "" || *file_path == "" {
			fmt.Println("You should fill the \"path\" and \"hide\" value")
			return
		}
		handleEncodee(*file_path, *hide_path, *output_path)

	case "decode":
		if *file_path == "" {
			fmt.Println("The file must be filled")
			return
		}
		hanldeDeocde(*file_path, *output_decode)

	default:
		fmt.Printf("Invaid mode. use \"encode\" or \"deconde\"")
		return

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

	can := CheckIfItFits(img_size, hide_capacity+8)
	if can != true {
		fmt.Println("It Doesn't fit in the image")
		return
	}

	fmt.Println("Encoding file into image...")

	

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
