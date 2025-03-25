package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello, this is NIRAJ BENDSURE.\nThis is our LAST PRACTICAL."
	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Text file created successfully!")
}
