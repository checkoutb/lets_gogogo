// package sdq
package main

import (
    "fmt"
)


// result := make(map[int]int)
// for i := range nums {
//     if _, ok := result[nums[i]]; ok {
//         delete(result, nums[i])
//     } else {
//         result[nums[i]] = 1
//     }
// }
// for i := range result {                  // ????   按方法体来说，这里是 遍历  map.keys .. 但是，难道 不写就自动忽视掉了？？？？
//     return i
// }
// return 0


// Runtime: 16 ms, faster than 89.61% of Go online submissions for Single Number.
// Memory Usage: 6.2 MB, less than 53.26% of Go online submissions for Single Number.
// 等一个 重复3次。 等一个 2个unique
func singleNumber(nums []int) int {
    ans := 0
    for _, v := range nums {
        ans ^= v
    }
    return ans
}

func main_LT0136_20211119() {
// func main() {

    fmt.Println("ans:")


}
