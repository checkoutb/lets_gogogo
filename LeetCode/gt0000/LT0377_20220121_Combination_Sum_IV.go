// package sdq
package main

import (
    "fmt"
    "sort"
)





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum IV.
// Memory Usage: 2 MB, less than 69.57% of Go online submissions for Combination Sum IV.
// dfs, dp
func combinationSum4(nums []int, target int) int {
    arr := make([]int, target + 1)
    sort.Ints(nums)
    arr[0] = 1
    for i := 0; i < target; i++ {
        if arr[i] != 0 {
            for _, v := range nums {
                t2 := i + v
                if t2 > target {
                    break
                }
                arr[t2] += arr[i]
            }
        }
    }
    return arr[target]
}


func main_LT0377_20220121() {
// func main() {

    fmt.Println("ans:")


}
