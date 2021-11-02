// package sdq
package main

import (
    "fmt"
)



// if nums[left] < nums[mid]  {
//     if target < nums[mid] && target >= nums[left] {
//         right = mid - 1
//     } else {
//         left = mid + 1
//     }
// } else if nums[left] == nums[mid] {
//     left++
// } else {
//     if target > nums[mid] && target <= nums[right] {
//         left = mid + 1
//     } else {
//         right = mid - 1
//     }
// }



// 1) everytime check if targe == nums[mid], if so, we find it.
// 2) otherwise, we check if the first half is in order (i.e. nums[left]<=nums[mid]) 
//   and if so, go to step 3), otherwise, the second half is in order,   go to step 4)
// 3) check if target in the range of [left, mid-1] (i.e. nums[left]<=target < nums[mid]), if so, do search in the first half, i.e. right = mid-1; otherwise, search in the second half left = mid+1;
// 4)  check if target in the range of [mid+1, right] (i.e. nums[mid]<target <= nums[right]), if so, do search in the second half, i.e. left = mid+1; otherwise search in the first half right = mid-1;
// 
// while(left<=right)
// {
//     mid = (left + right) >> 1;
//     if(nums[mid] == target) return true;
//     // the only difference from the first one, trickly case, just updat left and right
//     if( (nums[left] == nums[mid]) && (nums[right] == nums[mid]) ) {++left; --right;}
//     else if(nums[left] <= nums[mid])
//     {
//         if( (nums[left]<=target) && (nums[mid] > target) ) right = mid-1;
//         else left = mid + 1; 
//     }
//     else
//     {
//         if((nums[mid] < target) &&  (nums[right] >= target) ) left = mid+1;
//         else right = mid-1;
//     }
// }
// return false;



// 10/26/2021 11:11	Accepted	11 ms	3.4 MB	golang
// 04/21/2021 09:02	Wrong Answer	N/A	N/A	python3
// 02/01/2021 17:47	Wrong Answer	N/A	N/A	cpp
// 03/03/2018 17:03	Wrong Answer	N/A	N/A	java
// Runtime: 11 ms, faster than 16.30% of Go online submissions for Search in Rotated Sorted Array II.
// Memory Usage: 3.4 MB, less than 44.57% of Go online submissions for Search in Rotated Sorted Array II.
func search81(nums []int, target int) bool {
    sz1 := len(nums)
    l, r := 0, sz1 - 1
    for l <= r {
        mid := l + (r - l) >> 1
        t1, t2, t3 := nums[l], nums[mid], nums[r]
        if t2 == target {
            return true
        } else if t2 > target {
            if t1 < t2 {
                if t1 == target {
                    return true
                } else if t1 > target {
                    l = mid + 1
                } else {
                    r = mid - 1
                }
            } else if t1 == t2 {
                // 1 1 1 1 1 1 1 1 2 1 1   or    1 1 2 1 1 1 1 1 1 1 
                l++             // ........  for nums[l] == xx {l++} ?  not slow  it's o(n)
            } else {
                r = mid - 1
            }
        } else {        // t2 < target
            if t3 > t2 {
                if t3 == target {
                    return true
                } else if t3 > target {
                    l = mid + 1
                } else {
                    r = mid - 1
                }
            } else if t2 == t3 {
                r--
            } else {
                l = mid + 1
            }
        }
    }
    return false
}

func main_LT0081_20211025() {
// func main() {

    fmt.Println("ans:")

    arr := []int{2,5,6,0,0,1,2}
    tar := 3

    ans := search81(arr, tar)

    fmt.Println(ans)

}
