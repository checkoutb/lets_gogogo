// package sdq
package main

import (
    "fmt"
)


// patches, i, lookingFor := 0, 0, 1
// for lookingFor <= n {
//     if i < len(nums) && nums[i] <= lookingFor {
//         lookingFor += nums[i]
//         i++
//     } else {
//         lookingFor += lookingFor
//         patches++
//     }
// }


// patch := 0
// for i, x := 0, 1; x <= n; {
//     if i < len(nums) && x >= nums[i] {
//         x += nums[i]
//         i++
//     } else {
//         x += x
//         patch++
//     }
// }


// func minPatches(nums []int, n int) (patch int) {
// 	for i, x := 0, 1; x <= n; {
// 		if i < len(nums) && x >= nums[i] {
// 			x += nums[i]
// 			i++
// 		} else {
// 			x += x
// 			patch++
// 		}
// 	}
// 	return
// }



// Runtime: 8 ms, faster than 20.00% of Go online submissions for Patching Array.
// Memory Usage: 3.5 MB, less than 10.00% of Go online submissions for Patching Array.
// 肯定是add。  从1开始，++， 如果不行，就 add 不行的那个值。
func minPatches(nums []int, n int) int {
    idx := 0
    nxt := 1
    ans := 0
    for nxt <= n {
        if idx < len(nums) {
            if nums[idx] <= nxt {
                nxt += nums[idx]
                idx++
            } else {
                ans++           // add nxt
                nxt += nxt
            }
        } else {
            ans++
            nxt += nxt
        }
    }
    return ans
}

func main_LT0330_20220113() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{1,3}
    // n := 6

    // arr := []int{1,5,10}
    // n := 20

    arr := []int{1,2,2}
    n := 5

    ans := minPatches(arr, n)

    fmt.Println(ans)

}
