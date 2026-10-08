package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("WELCOME TO STAN")

	file_path := flag.String("path", "", "The file to hide in image") //pointer return karlay
	flag.Parse()

	if *file_path == "" {
		fmt.Println("Error: -path is empty darling fill it")
		flag.Usage()
		return
	}

	file_meta_data, err := os.Stat(*file_path)
	
	if err != nil {
		fmt.Printf("Error opning file: %v\n",err)
		return
	}


	fmt.Println("Name: " + file_meta_data.Name())
	fmt.Println(file_meta_data.Size())

	



}
