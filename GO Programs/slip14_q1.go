package main
import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("Original Slice:", slice)
    slice = append(slice, 6)
    fmt.Println("After Append:", slice)
    rem := 2
    slice = append(slice[:rem], slice[rem+1:]...)
    fmt.Println("After Remove:", slice)
    copys := make([]int, len(slice))
    copy(copys, slice)
    fmt.Println("Copied Slice:", copys)
}