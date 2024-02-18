// package sdq
package main

import (
    "fmt"
    "sort"
)




// Runtime142 ms
// Beats
// 96.67%
// Memory9.4 MB
// Beats
// 90%

func divideArray(nums []int, k int) [][]int {
    
    sort.Ints(nums)

    for i := 0; i < len(nums); i += 3 {
        if nums[i] + k < nums[i + 2] {
            return [][]int{};
        }
    }
    var ans = make([][]int, len(nums) / 3)
    for i := 0; i < len(nums) / 3; i += 1 {
        ans[i] = make([]int, 3)
    }

    for i := 0; i < len(nums); i += 1 {
        ans[i / 3][i % 3] = nums[i]
    }
    return ans;
}


func main_LT2966_20240108() {
// func main() {

    fmt.Println("ans:")


}
