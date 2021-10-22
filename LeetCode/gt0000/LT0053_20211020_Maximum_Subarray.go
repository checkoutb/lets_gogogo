// package sdq
package main

import (
    "fmt"
    // "math"
)


// int current = nums[0];
// int max = nums[0];
// for (int i = 1; i < nums.length; i++) {
//     current = Math.max(nums[i], current + nums[i]);
//     max = Math.max(max, current);
// }

// resultSum, currentSum := nums[0], nums[0]
// for i := 1; i < len(nums); i++ {
//     currentSum = max(currentSum + nums[i], nums[i])
//     resultSum = max(resultSum, currentSum)
// }

// for _, num := range nums {
//     if curr_count < 0 {
//         curr_count = 0
//     }
//     curr_count += num
//     max_count = max(max_count, curr_count)
// }


// Runtime: 135 ms, faster than 15.10% of Go online submissions for Maximum Subarray.
// Memory Usage: 9.4 MB, less than 43.84% of Go online submissions for Maximum Subarray.
// 感觉有个有名字的算法 可以用。。但是。。。
// 最多直接跳过 负数。
// 感觉是 一路往前走，只要是 正数 就继续走， 直到 变成负数，就 重新开始。
// [-1] ...
func maxSubArray(nums []int) int {
    ans, tmp := -1111111, 0

    for _, v := range nums {
        tmp += v
        if tmp > ans {
            ans = tmp
        }
        if tmp < 0 {
            tmp = 0
        }
    }

    return ans
}


func main_LT0053_20211020() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{-2,1,-3,4,-1,2,1,-5,4}
    // arr := []int{1}
    arr := []int{5,4,-1,7,8}

    ans := maxSubArray(arr)

    fmt.Println(ans)

}
