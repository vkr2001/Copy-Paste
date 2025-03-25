package main
import "fmt"

func isPalindrome(n int) int {
    var rev, rem, org int
    rev = 0
    org = n
    for(n != 0) {
        rem = n % 10
        rev = rev * 10 + rem
        n /= 10
    }
    if (org == rev) {
        return 1
    } else {
        return 0
    }
}

func main(){
    var n, r int
    fmt.Print("Enter the number: ")
    fmt.Scanln(&n)
    r = isPalindrome(n)
    if (r == 1) {
        fmt.Printf("%d is a Palidrome number.\n", n)
    } else {
        fmt.Printf("%d is not a Palindrome number.\n", n)
    }
 }