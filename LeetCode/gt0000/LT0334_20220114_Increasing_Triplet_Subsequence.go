// package sdq
package main

import (
    "fmt"
)



// min := 1000000000000000000
// max := 1000000000000000000
// for _, num := range nums {
//     if num <= min {
//         min = num
//     } else if num <= max {
//         max = num
//     } else {
//         return true
//     }
// }
// 如果 小于 min，那么 更新min。 如果 大于min， 和max 判断， 如果小于max 那么更新max。 如果大于 max 就 true
// 2,3,1,5 
// max 被激活，说明 前面有比 max 小的。 现在有一个数 大于max， 那么 就能确定 true 了。
// 但是真的很难 想到啊。




// Runtime: 122 ms, faster than 21.65% of Go online submissions for Increasing Triplet Subsequence.
// Memory Usage: 18.3 MB, less than 14.43% of Go online submissions for Increasing Triplet Subsequence.
// 增长 × 2
// 找一个数，前面 有比它小的， 后面有比它大的。
// 有点费内存啊。
func increasingTriplet(nums []int) bool {
    sz1 := len(nums)
    barr := make([]bool, sz1)        // 后面是否有 大于它的。
    mx := nums[sz1 - 1]
    for i := sz1 - 2; i >= 0; i-- {
        if mx > nums[i] {
            barr[i] = true
        } else {
            mx = nums[i]
        }
    }
    mn := nums[0]
    for i := 1; i < sz1; i++ {
        if mn < nums[i] && barr[i] {
            return true
        }
        if mn > nums[i] {
            mn = nums[i]
        }
    }
    return false
}

func main_LT0334_20220114() {
// func main() {

    fmt.Println("ans:")


}
