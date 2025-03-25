package main
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    //defer file.Close()
    fmt.Println("File opened successfully for reading.")
}