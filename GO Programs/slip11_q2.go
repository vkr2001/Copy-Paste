package main
import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 8
	ch <- 5
	ch <- 11
	ch <- 17
	ch <- 23
	fmt.Println("Capacity:", cap(ch)) // Prints: Capacity: 3
	fmt.Println("Length:", len(ch))   // Prints: Length: 3
	val := <-ch
	fmt.Println("Modified length:", len(ch)) // Prints: Modified length: 2
	fmt.Println("Read value:", val) // Prints: Read value: 1
}