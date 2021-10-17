// package sdq
package main

import (
    "fmt"
)



// n := len(nums)
// found := false
// for i:=n-1;i>0;i-- {
//     if nums[i] <= nums[i-1] {
//         continue                     // 从后往前看， 非降序 就 继续 往前走。
//     }
//     j := n-1
//     for nums[i-1] >= nums[j] {           // 出现 降序后， 找到第一个 大于 [i-1]的 [j]。  [j] 就是 swap的对象。
//         j--
//         continue
//     }
//     nums[i-1], nums[j] = nums[j], nums[i-1]          // swap
//     sort.Ints(nums[i:])                          // sort 后半部分。
//     found  = true
//     break
// }
// if !found {
//     sort.Ints(nums)
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Next Permutation.
// Memory Usage: 2.5 MB, less than 47.48% of Go online submissions for Next Permutation.
// 找到一个 小于后续中的最大值的， 然后后续的升序。
// 55555  321
func nextPermutation(nums []int)  {
    sz1 := len(nums)
    mx := nums[sz1 -1]
    i := sz1 - 2
    for ; i >= 0; i-- {                 // 好像 越界自动退出的？ 之前写的 i++。。Leetcode上是 下面的 nums[i+1]越界，而不是 死循环。
        if nums[i] < mx {
            break
        } else {
            mx = nums[i]
        }
    }
    if i < 0 {
        reverse31(nums, 0, sz1 - 1)
        return
    }
    //  7 6 2 5 4   76425
    // 7 6 2 5 1   大于2的最小值。          // 后一半本身就是有序的，后一半是降序。   从后往前看，是一路非降序，第一次降序 就是 i
    mn := nums[i + 1]       // [i+1] > [i]
    mnidx := i + 1
    for j := i + 2; j < sz1; j++ {
        if nums[j] <= mn && nums[j] > nums[i] {          // 可以是 找到第一个小于等于[i]的， 前面一个就是 大于[i]的
                                            // <=  保证swap 后 依然有序。
            mn = nums[j]
            mnidx = j
        }
    }
    nums[i], nums[mnidx] = nums[mnidx], nums[i]
    reverse31(nums, i + 1, sz1 - 1)
}

func reverse31(nums []int, st int, en int) {
    for st < en {
        nums[st], nums[en] = nums[en], nums[st]
        st++
        en--
    }
}

func main_LT0031_20211017() {
// func main() {

    fmt.Println("ans:")


}
