// package sdq
package main

import (
    "fmt"
)

// D D

// sort.Ints(nums)
// sol, left, right := 0, 0, len(nums) - 1
// for left < right {
//     if nums[left] + nums[right] < target {
//         sol += right - left
//         left++
//     } else {
//         right--
//     }
// }
// 
// 6




// Runtime6 ms
// Beats
// 34.72%
// Memory2.7 MB
// Beats
// 56.74%

// sort + binary search
func countPairs(nums []int, target int) int {
    var ans int = 0;
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] < target {
                ans++;
            }
        }
    }
    return ans;
}


func main_LT2824_20231010() {
// func main() {

    fmt.Println("ans:")


}
