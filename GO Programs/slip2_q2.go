package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := os.Args[1]
	file, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", file.Name())
	fmt.Println("Size (in bytes):", file.Size())
	fmt.Println("Permissions:", file.Mode())
	fmt.Println("Last Modified:", file.ModTime())
}
