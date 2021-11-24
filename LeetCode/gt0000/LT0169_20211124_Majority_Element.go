// package sdq
package main

import (
    "fmt"
    "sort"
)


// public int majorityElement(int[] nums) {
//     int res = nums[0], count = 0;
//     for (int i = 0; i < nums.length; i++) {
//         if (count == 0) {
//             res = nums[i];
//         } 
//         if (res == nums[i]) {
//             count++;
//         } else {
//             count--;
//         }
//     }
//     return res;
// }
// .....意会了。。 大约就是  不同的时候--， 由于某个数(A)的个数>一半，所以 它(A)能活下来， 其他的数(B) 都被 和它(B)不想等的数--，导致count变0,从而 从res上 被踢下来了。


// Runtime: 16 ms, faster than 91.51% of Go online submissions for Majority Element.
// Memory Usage: 6.1 MB, less than 60.30% of Go online submissions for Majority Element.
// 只会sort， 然后返回中间的数。
// 线性时间，固定空间。
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums) / 2]
}


func main_LT0169_20211124() {
// func main() {

    fmt.Println("ans:")


}
