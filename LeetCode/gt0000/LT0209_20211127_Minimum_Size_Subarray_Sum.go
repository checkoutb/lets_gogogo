// package sdq
package main

import (
    "fmt"
)

// int i = 0, n = A.length, res = n + 1;
// for (int j = 0; j < n; ++j) {
//     s -= A[j];
//     while (s <= 0) {
//         res = Math.min(res, j - i + 1);
//         s += A[i++];
//     }
// }
// return res % (n + 1);
// s 就是 target


// int l = 0, r = 0, n = nums.size(), sum = 0, len = INT_MAX;
// while (r < n) {
//     sum += nums[r++];
//     while (sum >= s) {
//         len = min(len, r - l);
//         sum -= nums[l++];
//     }
// }


// 先 sum += [i]  然后再 for sum >= tar
// Runtime: 15 ms, faster than 13.75% of Go online submissions for Minimum Size Subarray Sum.
// Memory Usage: 4.2 MB, less than 18.22% of Go online submissions for Minimum Size Subarray Sum.
func minSubArrayLen(target int, nums []int) int {
    sum2, st := 0, 0
    ans := len(nums) + 10
    for i := 0; i < len(nums); i++ {
        // if sum2 >= target {
            for sum2 >= target {
                if i - st < ans {
                    ans = i - st
                    // fmt.Println(st, i, ans)
                }
                sum2 -= nums[st]
                st++
            }
        // }
        sum2 += nums[i]
    }
    for sum2 >= target {
        if len(nums) - st < ans {
            ans = len(nums) - st
        }
        sum2 -= nums[st]
        st++
    }
    if ans > len(nums) {
        return 0
    }
    return ans
}


// func main_LT0209_20211127() {
func main() {

    fmt.Println("ans:")

    tar := 15
    arr := []int{5,1,3,5,10,7,4,9,2,8}

    ans := minSubArrayLen(tar, arr)

    fmt.Println(ans)

}
