// package sdq
package main

import (
    "fmt"
)

// int a = 1, b = 1;
// while (n--)
//     a = (b += a) - a;
// return a;

// def climb_stairs(n)
//     a = b = 1
//     n.times { a, b = b, a+b }
//     a
// end




// Runtime: 3 ms, faster than 16.46% of Go online submissions for Climbing Stairs.
// Memory Usage: 2.1 MB, less than 8.79% of Go online submissions for Climbing Stairs.
// fibonacci ? 就是 自己是 前一步 + 前前一步 的和....  只需要2/3个 var 来dp
func climbStairs(n int) int {
    arr := make([]int, n + 1)
    for i := 0; i <= n; i++ {
        if i <= 2 {
            arr[i] = i
        } else {
            arr[i] = arr[i - 1] + arr[i - 2]
        }
    }
    return arr[n]
}

func main_LT0070_20211022() {
// func main() {

    fmt.Println("ans:")


}
