package main
import "fmt"

type student struct {
    name, class  string
    age   int
    per float32
}

func (s *student) show() {
    fmt.Printf("Name: %s\n", s.name)
    fmt.Printf("Age: %d\n", s.age)
    fmt.Printf("Class: %s\n", s.class)
    fmt.Printf("Percentage: %.2f\n", s.per)
}

func main() {
    s := &student{
        name:  "NSB",
        age:   21,
        class: "TY",
        per: 88.99,
    }
    s.show()
}
