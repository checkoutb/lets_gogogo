// package sdq
package main

import (
    "fmt"
    "sort"
)


// abs := func(a int) int {
//     if a < 0 {
//         return -a
//     }
//     return a
// }


// dang nian yeshi henren a.

// Runtime: 4 ms, faster than 94.69% of Go online submissions for 3Sum Closest.
// Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for 3Sum Closest.
func threeSumClosest(nums []int, target int) int {
    sz1 := len(nums)
    sort.Ints(nums)
    ans := 1111111
    for idx, v := range nums {
        l, r := idx + 1, sz1 - 1
        for l < r {
            t2 := v + nums[l] + nums[r]
            // t3 := abs(t2, target)
            if t2 == target {
                return t2
            } else if t2 < target {
                t3 := target - t2
                if (t3 < abs(ans, target)) {
                    ans = t2
                }
                l++
            } else {
                t3 := t2 - target
                if (t3 < abs(ans, target)) {
                    ans = t2
                }
                r--
            }
        }
    }
    return ans
}

func abs(a int, b int) int {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}

func main_LT0016_20211006() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{-1,2,1,-4}
    arr := []int{0,0,0,0}
    k := 1

    fmt.Println(threeSumClosest(arr, k))

}
