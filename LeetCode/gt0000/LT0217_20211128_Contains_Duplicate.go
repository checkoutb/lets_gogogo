// package sdq
package main

import (
    "fmt"
    "sort"
)





// Runtime: 16 ms, faster than 95.39% of Go online submissions for Contains Duplicate.
// Memory Usage: 6.2 MB, less than 89.04% of Go online submissions for Contains Duplicate.
// sort, 前后对比
// map
func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
            return true
        }
    }
    return false
}

func main_LT0217_20211128() {
// func main() {

    fmt.Println("ans:")


}
