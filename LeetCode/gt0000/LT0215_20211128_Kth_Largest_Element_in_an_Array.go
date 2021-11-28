// package sdq
package main

import (
    "fmt"
    "sort"
)


// heap.... priority queue

// sort.Ints(nums)
// return nums[len(nums) - k]

// partial_sort  nth_element

// Quick Select


// Runtime: 12 ms, faster than 33.72% of Go online submissions for Kth Largest Element in an Array.
// Memory Usage: 3.8 MB, less than 40.46% of Go online submissions for Kth Largest Element in an Array.
// sort下，就可以了
// 不过 估计考的是 快速排序 。。。。。 not 。。。 not again。。。
// [3,2,3,1,2,4,5,5,6]
// 4   需要4，  就是 6554，， 排序后的 元素， 不需要 distinct。
func findKthLargest(nums []int, k int) int {
    return dfsa215(nums, len(nums) - k)


    // sort.Ints(nums)
    
    // for i := len(nums) - 1; i > 0; i-- {
    //     if nums[i] != nums[i - 1] {
    //         k--
    //         if k == 0 {
    //             return nums[i]
    //         }
    //     }
    // }
    // return nums[0]

    // for i := len(nums) - 2; i >= 0; i-- {
    //     if nums[i] != nums[i + 1] {
    //         k--
    //         if k == 0 {
    //             return nums[i + 1]
    //         }
    //     }
    // }
    // return nums[0]
}

func dfsa215(nums []int, k int) int {
    sz1 := len(nums)
    if sz1 == 1 {
        return nums[0]
    }
    t2 := nums[sz1 - 1]
    l, r := 0, sz1 - 2
    for l <= r {
        if nums[l] < t2 {
            l++
        } else {        // >= t2
            nums[l], nums[r] = nums[r], nums[l]
            r--
        }
    }
    // l 是第一个 >= t2的元素下标  s是最后一个<t2的 . 如果有的话。
    nums[l], nums[sz1 - 1] = nums[sz1 - 1], nums[l]         // [l] == t2

    if l == k {
        // 。。。。。。。 第二大 ， 可以有多个相同第一的 第二大。。。。。             反了。。 就是快速排序的。
        return nums[l]
    } else if l > k {
        return dfsa215(nums[0 : l], k)
    } else {
        return dfsa215(nums[l + 1 : ], k - l - 1)
    }
}


func main_LT0215_20211128() {
// func main() {

    fmt.Println("ans:")


}
