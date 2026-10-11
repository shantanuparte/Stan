package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("WELCOME TO STAN")

	file_path := flag.String("path", "", "The image to hide") //pointer return karlay
	hide_path := flag.String("hide", "", "The file to hide in image")
	output_path := flag.String("output","output.png","The name of the output file")
	flag.Parse()

	img, err := LoadImgage(*file_path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	img_size := ImgCapacity(img)

	hide_file_stream, err := LoadHideFile(*hide_path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return

	}

	hide_capacity, err := HideCapacity(*hide_path)
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

	output, err := os.Create(*output_path)
	if output != nil {
		return
	}
    

	Encoding(img, hide_file_stream)
	

}




