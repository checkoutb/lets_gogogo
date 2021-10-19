// package sdq
package main

import (
    "fmt"
)



// jumps := 0
// farthest := 0
// curEnd := 0
// for i := 0; i < len(nums) - 1; i++ {
//     farthest = max(farthest, nums[i] + i)
//     if i == curEnd {
//         jumps++
//         curEnd = farthest
//     }
// }
// return jumps




// Runtime: 20 ms, faster than 54.92% of Go online submissions for Jump Game II.
// Memory Usage: 5.8 MB, less than 68.09% of Go online submissions for Jump Game II.
func jump(nums []int) int {
    sz1 := len(nums)
    if (sz1 == 1) {
        return 0
    }
    ans := 0
    mx := nums[0]
    stp := nums[0]
    for i := 0; i < sz1; i++ {
        if i <= stp {
            if i + nums[i] > mx {
                mx = i + nums[i]
            }
        } else {
            ans++
            stp = mx
            i--
        }
    }
    return ans + 1
}


func main_LT0045_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{2,3,1,1,4}
    // arr := []int{2,3,0,1,4}
    arr := []int{1,1,1,1,1,1}

    ans := jump(arr)

    fmt.Println(ans)

}
