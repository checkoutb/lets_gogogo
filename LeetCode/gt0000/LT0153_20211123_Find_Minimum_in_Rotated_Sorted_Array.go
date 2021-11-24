// package sdq
package main

import (
    "fmt"
)


// while (lo<=hi){
//     int mid=(lo+hi)/2;
//     min=Math.min(min,arr[mid]);
//     if(arr[lo]<=arr[mid]){
//         min=Math.min(min,arr[lo]);
//         lo=mid+1;
//     }else {
//         hi=mid-1;
//     }
// }

// while (l < r) {
//     const m = Math.floor((l + r) / 2);
//     if (nums[m] > nums[r]) l = m + 1;
//     else r = m;
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Minimum in Rotated Sorted Array.
// Memory Usage: 2.5 MB, less than 32.34% of Go online submissions for Find Minimum in Rotated Sorted Array.
func findMin(nums []int) int {
    sz1 := len(nums)
    l, r := 0, sz1 - 1
    for l < r {
        mid := (l + r) / 2
        // fmt.Println(l, mid, r, nums[l], nums[mid], nums[r])
        if nums[mid] > nums[l] {        // 0 1 2   2 4 1
            if nums[l] > nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else if nums[mid] < nums[l] {
            r = mid
        } else {    // l == mid
            if nums[l] < nums[r] {
                r = l
            } else {
                l = r
            }
        }
    }
    return nums[l]
}

func main_LT0153_20211123() {
// func main() {

    fmt.Println("ans:")

    arr := []int{4,5,6,7,0,1,2}
    // arr := []int{2,1}

    ans := findMin(arr)

    fmt.Println(ans)

}
