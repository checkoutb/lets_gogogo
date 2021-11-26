// package sdq
package main

import (
    "fmt"
)


// func isPowerOfTwo(n int) bool {
//     if n < 1 {
//         return false
//     }
//     return n & (n - 1) == 0
// }
// ... 移除最右侧的1....


// Runtime: 6 ms, faster than 6.85% of Go online submissions for Power of Two.
// Memory Usage: 2.3 MB, less than 8.56% of Go online submissions for Power of Two.
// 首位不能是1   int32首位1 是负数。   这里int == int64
// 不用loop 就 直接 32个 bit计算啊。  就是要计算有多少个1. 并且符号位不能是 1
func isPowerOfTwo(n int) bool {
    cnt1 := 0
    for i := 0; i < 31; i++ {
        if n & 1 == 1 {
            cnt1++
            if cnt1 > 1 {
                return false
            }
        }
        n >>= 1
    }
    return cnt1 == 1 && n == 0
}

func main_LT0231_20211126() {
// func main() {

    fmt.Println("ans:")


}
