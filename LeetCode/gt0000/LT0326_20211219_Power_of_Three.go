// package sdq
package main

import (
    "fmt"
)


// return ( n>0 &&  1162261467%n==0);

// return n > 0 == 3**19 % n


// Runtime: 39 ms, faster than 26.83% of Go online submissions for Power of Three.
// Memory Usage: 6.5 MB, less than 54.88% of Go online submissions for Power of Three.
func isPowerOfThree(n int) bool {
    if n <= 0 {
        return false
    }
    for n % 3 == 0 {
        n /= 3
    }
    return n == 1
}


func main_LT0326_20211219() {
// func main() {

    fmt.Println("ans:")


}
