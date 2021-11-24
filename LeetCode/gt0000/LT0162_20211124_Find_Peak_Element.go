// package sdq
package main

import (
    "fmt"
)


// 二分。。。
// while (l < h) {
//     mid = (l + h) / 2;
//     if (num[mid] > num[mid + 1])
//         h = mid;
//     else if (num[mid] < num[mid + 1])
//         l = mid + 1;
// }
// 用的是 mid 和 mid+1 对比。  由于2侧都是 -INF， 所以 只要 有一个 上坡，那么就必然有山尖。


// Runtime: 3 ms, faster than 25.14% of Go online submissions for Find Peak Element.
// Memory Usage: 2.9 MB, less than 15.41% of Go online submissions for Find Peak Element.
func findPeakElement(nums []int) int {
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i + 1] {
            return i
        }
    }
    return len(nums) - 1

    // b2 := true
    // for i := 0; i < len(nums) - 1; i++ {
    //     if nums[i] > nums[i + 1] {
    //         if b2 {
    //             return i
    //         }
    //         b2 = false
    //     } else {
    //         b2 = true
    //     }
    // }
    // return len(nums) - 1
}


func main_LT0162_20211124() {
// func main() {

    fmt.Println("ans:")


}
