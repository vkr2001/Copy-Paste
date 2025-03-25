package main
import "fmt"

func sqrnum(num int, sqChan, cuChan chan int) {
	sqSum := 0
	cuSum := 0

	for num > 0 {
		dig := num % 10
		num /= 10
		square := dig * dig
		cube := dig * dig * dig
		sqSum += square
		cuSum += cube
	}

	sqChan <- sqSum
	cuChan <- cuSum
}

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)
	sqChan := make(chan int)
	cuChan := make(chan int)

	go sqrnum(num, sqChan, cuChan)

	sqr := <-sqChan
	cu := <-cuChan

	fmt.Printf("Sum of squares of digits: %d\n", sqr)
	fmt.Printf("Sum of cubes of digits: %d\n", cu)
}