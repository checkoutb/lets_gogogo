// package sdq
package main

import (
    "fmt"
)

// for right > left {
//     right &= right - 1
// }
// return right
// ....
// Brian Kernighan 算法


// m, n := left, right
// offset := 0
// for m != n {
//     m >>= 1
//     n >>= 1
//     offset++
// }
// return m << offset


// for i := 30; i >= 0; i--{
//     if (left & (2 << i)) != (right & (2 << i)){
//         return left & (0-(2 << (i+1)))
//     }
// }
// return left



// Runtime: 60 ms, faster than 23.25% of Go online submissions for Bitwise AND of Numbers Range.
// Memory Usage: 6 MB, less than 23.89% of Go online submissions for Bitwise AND of Numbers Range.
// 101
// 111
// 公共前缀 . 
func rangeBitwiseAnd(left int, right int) int {
    arr1 := make([]int, 32)
    arr2 := make([]int, 32)
    for i := 0; (1 << i) <= left; i++ {
        if left & (1 << i) != 0 {
            arr1[i] = 1
        }
    }
    for i := 0; (1 << i) <= right; i++ {
        if right & (1 << i) != 0 {
            arr2[i] = 1
        }
    }
    ans := 0
    for i := 31; i >= 0; i-- {
        if arr1[i] != arr2[i] {
            for i >= 0 {
                arr1[i], arr2[i] = 0, 0
                i--
            }
        } else {
            if arr1[i] == 1 {
                ans |= (1 << i)
            }
        }
    }
    return ans
}


func main_LT0201_20211123() {
// func main() {

    fmt.Println("ans:")


}
