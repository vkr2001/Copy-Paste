package main
import "fmt"

func addSub(a, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}

func main() {
    var n1, n2 int
    fmt.Print("Enter the two numbers to add and subtract:")
    fmt.Scanln(&n1, &n2)
    sum, diff := addSub(n1, n2)
    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Difference: %d\n", diff)
}