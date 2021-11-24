// package sdq
package main

import (
    "fmt"
)



// while(start != end){
//     int mid = (start+ end) >>> 1;
//     if (nums[mid] > nums[end]){
//         start = mid+1;
//     }else if(nums[mid] < nums[end]  ){
//         end = mid;
//     }
//     else{
//         end--;           // ...
//     }
// }

// for low < high {
//     mid := low + (high-low) >> 1
//     if nums[low] < nums[high] {
//         return nums[low]
//     }
//     if nums[mid] > nums[low] {
//         low = mid + 1
//     } else if nums[mid] < nums[low] {
//         high = mid
//     } else {
//         low++
//     }
// }

// for left < right {
//     mid := (left + right) / 2
//     if nums[mid] < nums[right] {
//         right = mid
//     }else if nums[mid] > nums[right]{
//         left = mid + 1
//     }else {
//         right--
//     }
// }


// Runtime: 4 ms, faster than 85.99% of Go online submissions for Find Minimum in Rotated Sorted Array II.
// Memory Usage: 3.2 MB, less than 48.41% of Go online submissions for Find Minimum in Rotated Sorted Array II.
// 2 2 2 2 2 1 2 2
// 2 1 2 2 2 2 2 2  // 这种，好像只能 遍历啊。
func findMin(nums []int) int {
    sz1 := len(nums)
    l, r := 0, sz1 - 1
    for l < r {
        mid := (l + r) / 2
        t1, t2, t3 := nums[l], nums[mid], nums[r]
        if t1 < t2 {
            if t1 < t3 {
                // r = mid - 1     // just return nums[l]...
                return nums[l]
            } else {    // t1 >= t3     2 2 3 1 2
                l = mid + 1
            }
        } else if t1 > t2 {
            r = mid
        } else {        // t1 == t2
            if t2 < t3 {       // t1 == t2 < t3    2 2 2 2 3
                return nums[l]
            } else if t2 > t3 { // t1 == t2 > t3    3 3 3 1 2
                l = mid + 1
            } else {        // t1==t2==t3
                for i := l; i <= r; i++ {
                    if nums[l] > nums[i] {
                        return nums[i]
                    }
                }
                return nums[l]
            }
        }
    }
    return nums[l]
}


// func main_LT0154_20211123() {
func main() {

    fmt.Println("ans:")

    // arr := []int{1,3,5}
    // arr := []int{2,2,2,0,1}
    // arr := []int{2,2,3,4,5}
    // arr := []int{2,2,1,2,2}
    // arr := []int{2,2,2,1,2}
    arr := []int{2,3,2,2,2}

    ans := findMin(arr)

    fmt.Println(ans)

}
