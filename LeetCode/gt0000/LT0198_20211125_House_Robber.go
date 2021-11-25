// package sdq
package main

import (
    "fmt"
)


// for i := 2; i < len(nums); i++ {
//     dp[i] = max(dp[i-2] + nums[i], dp[i-1])
// }

// dfs, 如果本次rob，那么就从 下下个房子开始。  不rob就从下个开始，所以 是 max(nums[i] + rob(nums, i + 2), rob(nums, i + 1))

// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
// Memory Usage: 2 MB, less than 71.57% of Go online submissions for House Robber.
func rob(nums []int) int {
    rob, notRob := 0, 0
    for _, v := range nums {
        t2 := notRob + v
        notRob = rob
        if t2 > rob {
            rob = t2
        }
        // fmt.Println(rob, notRob)
    }
    if rob > notRob {
        return rob
    } else {
        return notRob
    }
}


func main_LT0198_20211125() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,2,3,1}
    // arr := []int{2,7,9,3,1}
    // arr := []int{2,1,1,2}

    fmt.Println(rob(arr))

}
