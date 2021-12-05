// package sdq
package main

import (
    "fmt"
)



// long long int lsb=x&(~(x-1));


// rightone := res & (^res + 1)
// res1 := 0
// for _, v := range(nums) {
//     if rightone & v == 0 {
//         res1 ^= v
//     }
// }
// return []int{res1, res1 ^ res}


// Runtime: 4 ms, faster than 97.07% of Go online submissions for Single Number III.
// Memory Usage: 4 MB, less than 100.00% of Go online submissions for Single Number III.
// 记得！
func singleNumber(nums []int) []int {
    t2 := 0
    for _, v := range nums {
        t2 ^= v
    }
    // .. 怎么一步获得 分界值 忘记了。。。
    for i := 0; i < 32; i++ {
        if t2 & (1 << i) != 0 {
            t2 &= (1 << i)
            break
        }
    }
    a, b := 0, 0
    for _, v := range nums {
        if v & t2 != 0 {
            a ^= v
        } else {
            b ^= v
        }
    }
    return []int{a, b}
}

func main_LT0260_20211129() {
// func main() {

    fmt.Println("ans:")


}
