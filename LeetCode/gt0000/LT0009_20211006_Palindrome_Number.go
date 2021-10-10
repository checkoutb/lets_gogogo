// package sdq
package main

import (
    "fmt"
)


// digits := []int{}
// for x != 0 {
//     digits = append(digits, x%10)
//     x = x/10
// }

// if x == y || x / 10 == y {
//     return true
// }

// Runtime: 34 ms, faster than 17.37% of Go online submissions for Palindrome Number.
// Memory Usage: 5.1 MB, less than 65.10% of Go online submissions for Palindrome Number.
func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    t2 := 0
    for t2 < x {
        t3 := x % 10
        x /= 10
        if (t2 == x) {
            return true
        }
        t2 *= 10
        t2 += t3
        if (t2 == x) {
            return true
        }
    }
    return t2 == 0
}


func main_LT0009_20211006() {
// func main() {

    fmt.Println("ans:")

    // x := 10
    // x := 121
    x := 0

    fmt.Println(isPalindrome(x))

}
