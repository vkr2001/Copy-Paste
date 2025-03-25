package main
import "fmt"

func checkEvenOdd(num int, evenChan chan<- int, oddChan chan<- int) {
    if num%2 == 0 {
        evenChan <- num
    } else {
        oddChan <- num
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evenChan := make(chan int)
    oddChan := make(chan int)

    for _, num := range nums {
        go checkEvenOdd(num, evenChan, oddChan)
    }

    go func() {
        for {
            select {
            case evenNum := <-evenChan:
                fmt.Println("Even number:", evenNum)
            case oddNum := <-oddChan:
                fmt.Println("Odd number:", oddNum)
            }
        }
    }()
    var input string
    fmt.Scanln(&input)
}