// package sdq
package main

import (
    "fmt"
    "sort"
)


// gg 
// discuss : merge sort， 在merge的时候 。。。


// error。。。 后续下于，不是 全部小于。。
func countSmaller(nums []int) []int {
    sz1 := len(nums)
    ans := make([]int, sz1)
    copy(ans, nums)
    sort.Ints(nums)
    map2 := map[int]int{}
    for i := 0; i < sz1; i++ {
        map2[nums[i]] = i
        for i + 1 < sz1 && nums[i] == nums[i + 1] {
            i++
        }
    }
    for i := 0; i < sz1; i++ {
        ans[i] = map2[ans[i]]
    }
    return ans
}

// tle
func countSmaller2(nums []int) []int {
    sz1 := len(nums)
    ans := make([]int, sz1)
    for i := sz1 - 2; i >= 0; i-- {
        j := i + 1
        for j < sz1 && nums[j] < nums[j - 1] {
            nums[j], nums[j - 1] = nums[j - 1], nums[j]
            j++
        }
        ans[i] = j - i - 1
        // for j := i + 1; j < sz1; j++ {
        //     if nums[j] < nums[j - 1] {
        //         nums[j], nums[j - 1] = nums[j - 1], nums[j]
        //     } else {
        //         ans[i] = (j - i - 1)
        //     }
        // }
        // for j := i + 1; j < sz1 && nums[j] < nums[j - 1]; j++ {
        //     // if nums[j] < nums[j - 1] {
        //         nums[j], nums[j - 1] = nums[j - 1], nums[j]
        //     // } else {
        //     //     break
        //     // }
        // }
    }
    return ans
}

func main_LT0315_20211218() {
// func main() {

    fmt.Println("ans:")


}
