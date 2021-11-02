// package sdq
package main

import (
    "fmt"
)


// while(i >=0 && j>=0)
// {
//     if(A[i] > B[j])
//         A[k--] = A[i--];
//     else
//         A[k--] = B[j--];
// }
// while(j>=0)
//     A[k--] = B[j--];
// 确实,   A是有序的. 所以 如果 i > 0  就 不需要 swap


// for k := m + n - 1; k >= 0; k-- {
//     if j < 0 || (i >= 0 && nums1[i] > nums2[j]) {
//         nums1[k] = nums1[i]
//         i--
//     } else {
//         nums1[k] = nums2[j]
//         j--
//     }
// }


// while (j >= 0) {
//     nums1[tar--] = i >= 0 && nums1[i] > nums2[j] ? nums1[i--] : nums2[j--];
// }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
// Memory Usage: 2.4 MB, less than 87.20% of Go online submissions for Merge Sorted Array.
// 我好像记得还有  merge N个 有序数组,  merge N个 有序链表...
// ... 不记得 return void
// 那就是反序了.
func merge88(nums1 []int, m int, nums2 []int, n int)  {
    idx := len(nums1) - 1
    m--
    n--
    for m >= 0 || n >= 0 {
        t1, t2 := -2100000000, -2100000000
        if m >= 0 {
            t1 = nums1[m]
        }
        if n >= 0 {
            t2 = nums2[n]
        }
        if t1 > t2 {
            nums1[idx] = t1
            m--
        } else {
            nums1[idx] = t2
            n--
        }
        idx--
    }
}


func main_LT0088_20211027() {
// func main() {

    fmt.Println("ans:")


}
