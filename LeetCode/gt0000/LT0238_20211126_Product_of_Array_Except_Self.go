// package sdq
package main

import (
    "fmt"
)


// vector<int> result(n, 1);
// for(int i=0; i<n; i++){
//     result[i]*=left;
//     result[n-1-i]*=right;
//     left*=nums[i];
//     right*=nums[n-1-i];
// }
// 应该可以的， 不用result，直接操作 nums， 就是 空间O(1)了




// O(1)的空间。。除非让我 divide。

// Runtime: 28 ms, faster than 78.97% of Go online submissions for Product of Array Except Self.
// Memory Usage: 7.8 MB, less than 81.72% of Go online submissions for Product of Array Except Self.
// Runtime: 42 ms, faster than 23.10% of Go online submissions for Product of Array Except Self.
// Memory Usage: 7.3 MB, less than 96.03% of Go online submissions for Product of Array Except Self.
// O(n) 且不使用divide
// 正向，反向 就可以了。
func productExceptSelf(nums []int) []int {
    sz1 := len(nums)
    ans, pre := make([]int, sz1), 1
    for i := sz1 - 1; i > 0; i-- {
        ans[i] = pre * nums[i]
        pre = ans[i]
        if pre == 0 {
            break
        }
    }
    // fmt.Println(ans)
    pre = 1
    for i := 0; i < sz1 - 1; i++ {
        ans[i] = pre * ans[i + 1]
        pre = pre * nums[i]
        if pre == 0 {
            i++
            for i < sz1 - 1 {
                ans[i] = 0
                i++
            }
        }
    }
    ans[sz1 - 1] = pre
    return ans

    // i := 1          // ans[0] is 0
    // pre = nums[0]
    // for i < sz1 && ans[i] == 0 {
    //     pre *= nums[i]
    //     i++
    // }
    // i--     // not 0, or last
    // for i < sz1 {
    //     ans
    //     i++
    // }

    // ans, arr2 := make([]int, sz1)

    // arr1, arr2 := make([]int, sz1), make([]int, sz2)
    // pre1, pre2 := 1, 1
    // for i := 0; i < sz1; i++ {
    //     arr1[i] = pre1 * nums[i]
    //     pre1 = arr1[i]
    //     arr2[sz1 - 1 - i] = pre2 * nums[sz1 - 1 - i]
    //     pre2 = arr2[sz1 - 1 - i]
    // }
    // for i := 0; i < sz1; i++ {

    // }
}

func main_LT0238_20211126() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{1,2,3,4}
    arr := []int{1,2,3,0,1,-1}

    ans := productExceptSelf(arr)

    fmt.Println(ans)

}
