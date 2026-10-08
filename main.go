package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("WELCOME TO STAN")

	file_path := flag.String("path", "", "The file to hide in image") //pointer return karlay

	flag.Parse()
	
	

	img, err := LoadImgage(*file_path)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		return
	}

	

}
