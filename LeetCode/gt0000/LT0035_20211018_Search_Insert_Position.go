// package sdq
package main

import (
    "fmt"
)


// i := sort.Search(len(nums), func(i int) bool {
//     return nums[i] >= target
// })


// while (low <= high) {
//     int mid = low + (high-low)/2;
//     if (nums[mid] < target)
//         low = mid+1;
//     else
//         high = mid-1;
// }



// Runtime: 4 ms, faster than 80.28% of Go online submissions for Search Insert Position.
// Memory Usage: 3 MB, less than 46.32% of Go online submissions for Search Insert Position.
func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}


func main_LT0035_20211018() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,3,5,6}
    tar := 2

    ans := searchInsert(arr, tar)

    fmt.Println(ans)

}
