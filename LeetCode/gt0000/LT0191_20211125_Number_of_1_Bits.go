// package sdq
package main

import (
    "fmt"
)



// for num != 0 {
//     num = num & (num-1)
//     count++
// }
// 前几天的那个， 这个方法可以求到 最右侧的1. 。 这个是 移除右侧所有的0.


// ans += num & 1


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
// Memory Usage: 2.1 MB, less than 28.67% of Go online submissions for Number of 1 Bits.
func hammingWeight(num uint32) int {
    ans := 0
    for num != 0 {
        if num & 1 == 1 {
            ans++
        }
        num >>= 1
    }
    return ans
}


func main_LT0191_20211125() {
// func main() {

    fmt.Println("ans:")


}
