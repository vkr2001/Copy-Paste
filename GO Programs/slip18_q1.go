package main
import "fmt"

func multi(num int) {
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}

func main() {
    var n int
    fmt.Print("Enter the number to print multiplication table: ")
    fmt.Scanln(&n)
    multi(n)
}