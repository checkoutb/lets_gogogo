// package sdq
package main

import (
    "fmt"
)





// Runtime: 144 ms, faster than 81.82% of Go online submissions for Rotate Function.
// Memory Usage: 8.3 MB, less than 72.73% of Go online submissions for Rotate Function.
// 0*[0] + 1*[1] + 2*[2] + 3*[3]    // -[1] -[2] -[3] + 3[0]        -[1] -[2] -[3] -[0] + 4[0]
// 0*[1] + 1*[2] + 2*[3] + 3[0]   // -[2] -[3] -[0] + 3[1]          -all + 4[1]
// 0[2]     1[3]    2[0]    3[1]
// 顺时针，逆时针 无所谓。
func maxRotateFunction(nums []int) int {
    ans := 0
    sz1 := len(nums)
    sum2 := 0
    sumt2 := 0
    for i, v := range nums {
        sum2 += v
        sumt2 += i * v
    }
    ans = sumt2
    for i := 0; i < sz1 - 1; i++ {
        sumt2 += sz1 * nums[i] - sum2
        if sumt2 > ans {
            ans = sumt2
        }
    }
    return ans
}


func main_LT0396_20220125() {
// func main() {

    fmt.Println("ans:")


}
