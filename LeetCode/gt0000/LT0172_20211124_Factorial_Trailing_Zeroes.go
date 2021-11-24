// package sdq
package main

import (
    "fmt"
)




// c := 0
// for n >= 5 {
//     n = n/5
//     c += n
// }
// return c
// 。。。


// return n == 0 ? 0 : n / 5 + trailingZeroes(n / 5);





// 我感觉是这种快，毕竟 *5， 但是第一次7ms， 第二次0ms。
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2.1 MB, less than 74.65% of Go online submissions for Factorial Trailing Zeroes.
// Runtime: 7 ms, faster than 26.76% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2.2 MB, less than 39.44% of Go online submissions for Factorial Trailing Zeroes.
func lt0172b(n int) int {
    ans := 0
    for i := 5; i <= n; i *= 5 {
        ans += n / i
    }
    return ans
}


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Factorial Trailing Zeroes.
// Memory Usage: 2.1 MB, less than 74.65% of Go online submissions for Factorial Trailing Zeroes.
// 10 5 
// 25 - 6   5 10 15 20 25               25 = 5×5 = 10×10
// 24 - 4   5*2 10 15*12 20
func trailingZeroes(n int) int {
    ans := 0
    for i := 5; i <= n; i += 5 {
        t2 := i
        for t2 % 5 == 0 {
            ans++
            t2 /= 5
        }
    }
    return ans
}


func main_LT0172_20211124() {
// func main() {

    fmt.Println("ans:")


}
