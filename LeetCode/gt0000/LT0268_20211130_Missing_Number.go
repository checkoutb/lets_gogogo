// package sdq
package main

import (
    "fmt"
)

// missing := len(nums)
// for i := 0; i < len(nums); i++ {
//     missing ^= i ^ nums[i]
// }
// return missing

// totalSum := len(nums)
// for idx, i := range nums {
//     totalSum += idx - i
// }
// return totalSum

// for(int i=0; i<nums.length; i++){
//     res ^= i;
//     res ^= nums[i];
// }

// int res = nums.length;
// for (i = 0; i < nums.length; i++) {
//     xor = xor ^ i ^ nums[i];
// }
// return xor ^ i;

// 下标是 [0, sz1] 全部出现的， 值是 [0, sz1] 减去 一个值，  所以就变成了  1个出现一次，其他的都是出现2次。



// Runtime: 12 ms, faster than 98.59% of Go online submissions for Missing Number.
// Memory Usage: 6.3 MB, less than 100.00% of Go online submissions for Missing Number.
// sum
func missingNumber(nums []int) int {
    t2 := 0
    for _, v := range nums {
        t2 += v
    }
    sz1 := len(nums)
    return sz1 * (sz1 + 1) / 2 - t2
}


func main_LT0268_20211130() {
// func main() {

    fmt.Println("ans:")


}
