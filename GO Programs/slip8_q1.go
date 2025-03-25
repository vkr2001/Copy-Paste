package main

import "fmt"

type book struct {
    bid     int
    bname   string
    bauthor string
    price   float32
}

func main() {
    var b [10]book
    var n int
    fmt.Print("Enter the number of books: ")
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Print("\nEnter book ID: ")
        fmt.Scan(&b[i].bid)
        fmt.Print("Enter book name: ")
        fmt.Scan(&b[i].bname)
        fmt.Print("Enter book author: ")
        fmt.Scan(&b[i].bauthor)
        fmt.Print("Enter book price: ")
        fmt.Scan(&b[i].price)
    }
    fmt.Println("\n\nBook Information")
    fmt.Println("\nBOOK ID  TITLE       AUTHOR      PRICE")
    fmt.Println("---------------------------------------")
    for i := 0; i < n; i++ {
        fmt.Println(b[i].bid, "\t\t", b[i].bname, "\t\t", b[i].bauthor, "\t\t", b[i].price)
    }
}
