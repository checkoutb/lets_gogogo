// package sdq
package main

import (
    "fmt"
)



// Runtime12 ms
// Beats
// 100%
// Sorry, there are not enough accepted submissions to show data
// Memory2.9 MB
// Beats
// 100%


// -1 -1 -1
// 0 -1 ...
// [-9] ---> -9 ...
func maxStrength(nums []int) int64 {
    // if len(nums) == 1 && nums[0] < 0 {
    //     return 0
    // }
    
    var cntn0 = 0       // not 0
    // var b0 = false
    var ans int64 = 1
    var mxNeg = -10
    for _, n := range nums {
        if n > 0 {
            ans *= int64(n)
            cntn0++
        } else if n < 0 {
            ans *= int64(n)
            if n > mxNeg {
                mxNeg = n
            }
            cntn0++
        } else {
            //b0 = true
        }
    }
    if (cntn0 == 0) {
        ans = 0
    }
    if (ans < 0) {
        if len(nums) != 1 {
            if cntn0 == 1 {
                ans = int64(0)
            } else {
                ans /= int64(mxNeg)
            }
        }

        // if (cntn0 == 1) {
        //     ans = int64(0)
        // } else {
        //     ans /= int64(mxNeg)
        // }
    }
    return ans

    // map2 := map[int]int{}
    // for _, n := range nums {
    //     map2[n]++
    // }


    // var ans int64 = 0
    // var lst int = 0
    // for _, n := range nums {
    //     if n > 0 {
    //         if (ans == 0) {
    //             ans = int64(1)
    //         }
    //         ans *= int64(n)
    //     } else {
    //         if (lst < 0) {

    //         }
    //     }
    // }
    // return ans

    // sort.Ints(nums)
    // lst := 0
    // for _, n := range nums {

    // }
}


func main_LT2708_20230528() {
// func main() {

    fmt.Println("ans:")


}
