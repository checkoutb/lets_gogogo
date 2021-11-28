// package sdq
package main

import (
    "fmt"
)


// class Solution:
//     def rob(self, nums):
//         def rob(nums):
//             now = prev = 0
//             for n in nums:
//                 now, prev = max(now, prev + n), now
//             return now
//         return max(rob(nums[len(nums) != 1:]), rob(nums[:-1]))



// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
// Memory Usage: 2.1 MB, less than 70.19% of Go online submissions for House Robber II.
// 计算2次？ 第一次 不抢[0],可以抢[sz1-1]， 第二次抢[0]，不抢[sz1-1]？
func rob(nums []int) int {
    sz1 := len(nums)
    if sz1 == 1 {
        return nums[0]
    }
    ans := 0
    rob, notr := nums[0], 0         // rob [0]
    for i := 1; i < sz1 - 1; i++ {      // not [sz1-1]
        t2 := notr + nums[i]
        notr = rob
        if t2 > rob {
            rob = t2
        }
    }
    ans = rob

    rob, notr = nums[1], 0     // not rob [0]
    for i := 2; i < sz1; i++ {
        t2 := notr + nums[i]
        notr = rob
        if t2 > rob {
            rob = t2
        }
    }
    if rob > ans {
        ans = rob
    }
    return ans
}


func main_LT0213_20211127() {
// func main() {

    fmt.Println("ans:")


}
