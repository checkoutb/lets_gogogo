// package sdq
package main

import (
    "fmt"
)


// for low < high {
//     mid := (low + high)/2
//     if nums[mid] < target {
//         low = mid+1
//     } else {
//         high = mid
//     }
// }
//         find ans[0]
// high = len(nums) - 1
// for low < high {
//     mid := (low + high)/2 + 1
//     if nums[mid] <= target {
//         low = mid
//     } else {
//         high = mid - 1
//     }
// }
//          find ans[1]

// lower_bound(x), lower_bound(x+1)/upper_bound(x)



// Runtime: 4 ms, faster than 98.70% of Go online submissions for Find First and Last Position of Element in Sorted Array.
// Memory Usage: 4 MB, less than 43.45% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
    ans := []int{-1, -1}
    if len(nums) == 0 {
        return ans
    }
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        // fmt.Println(l, r, mid)
        t2 := nums[mid]
        if target == t2 {
            l, r = mid, mid - 1
        } else if target < t2 {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    if l < len(nums) && nums[l] == target {
        for r = l; r >= 0 && nums[r] == target; r-- {}
        ans[0] = r + 1
        for r = l; r < len(nums) && nums[r] == target; r++ {}
        ans[1] = r - 1
    }
    return ans
}


// func main_LT0034_20211017() {
func main() {

    fmt.Println("ans:")
    arr := []int{1}
    // arr := []int{5,7,7,8,8,10}
    // tar := 8
    tar := 6

    ans := searchRange(arr, tar)

    fmt.Println(ans[0], " - ", ans[1])
}
