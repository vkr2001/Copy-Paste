package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
        delay()
    }
}

func delay() {
    rand.Seed(time.Now().UnixNano())
    ms := rand.Intn(250)
    time.Sleep(time.Duration(ms) * time.Millisecond)
}
