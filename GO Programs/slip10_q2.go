package main
import "fmt"

func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
    var n int
    fmt.Print("Enter the number of fibonacci series: ")
    fmt.Scan(&n)
    fmt.Printf("Fibonacci Series of %d numbers:\n", n)
	c := make(chan int)
	go fibonacci(n, c) 
	for num := range c {
		fmt.Println(num)
	}
}