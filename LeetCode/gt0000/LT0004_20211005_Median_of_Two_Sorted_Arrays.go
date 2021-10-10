// package sdq
package main

import (
    "fmt"
)





// gg




func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var ans float64 = 0.0
    sz1, sz2 := len(nums1), len(nums2)
    st, en := -1000001, 1000001
    for st < en {
        t2 := (st + en) / 2
        s1, e1, s2, e2 := 0, sz1, 0, sz2
        for (s1 < e1) {
            mid := (s1 + e1) / 2
            if (nums1[mid] < t2) {
                s1 = mid + 1
            } else if (nums1[mid] == t2) {
                e1 = mid
                // s1, e1 = mid, mid
            } else {
                e1 = mid
            }
        }
        for s2 < e2 {
            mid := (s2 + e2) / 2
            if (nums2[mid] < t2) {
                s2 = mid + 1
            } else if nums2[mid] == t2 {
                e2 = mid
                // s2, e2 = mid, mid
            } else {
                e2 = mid
            }
        }
        if (s1 + s2 + 2) > (sz1 + sz2) / 2 {

        }
    }
    return ans
}


// // 好像不是 log(n+m) 了, 大了几倍。
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     var ans float64 = 0.0
//     st, en := -1000001, 1000001

//     for st < en {
//         mid = (st + en) / 2

//     }

//     return ans
// }



// 不应该二分 下标，而是应该二分值。

// log -> 二分。
// 1 2 3
// 11 12 13 14 15 16 17
// 2 14   1 2 x x x x 14 15 16 17
// 3 12   1 2 3 x [12 13] 14 15 16 17

// 1 3 5 7 9 11
// 2 4 6 8 10 12
// 5 6
// 1 3 5     6 8 10 12
// 好像不行。。

// 看来要 二分 然后再 二分搜索。 那就是 log*log 。。。 logN * logM == log(N+M) ....
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
    var ans float64 = 0.0
    sz1, sz2 := len(nums1), len(nums2)
    if sz1 == 0 || sz2 == 0 {
        if sz2 != 0 {
            nums1 = nums2
        }
        // nums1 = append(nums1, nums2...)
        t2 := len(nums1) / 2
        if len(nums1) % 2 == 0 {
            return float64 (nums1[t2] + nums1[t2 - 1]) / 2
        } else {
            return float64 (nums1[t2])
        }
    }
    st1, en1, st2, en2 := 0, sz1, 0, 0
    for st1 < en1 {
        mid := (st1 + en1) / 2
        t2 := nums1[mid]
        st2, en2 = 0, sz2
        for st2 < en2 {
            m2 := (st2 + en2) / 2
            t3 := nums2[m2]
            if t3 < t2 {
                st2 = m2 + 1
            } else if t2 == t3 {
                st2, en2 = m2, m2
            } else {
                en2 = m2
            }
        }
        if st2 == sz2 {
            st2 = sz2 - 1
        } else if nums2[st2] > t2 && st2 > 0 {
            st2--
        }
        if (mid + st2) > ((sz1 + sz2) / 2) {
            en1 = mid
        } else if (mid + st2) == (sz1 + sz2) / 2 {
            st1, en1 = mid, mid
        } else {
            st1 = mid + 1
        }
    }
    if (sz1 == st1) {
        st1 -= 1
    }
    fmt.Println(st1)
    fmt.Println(st2)
    if ((sz1 + sz2) % 2 == 0) {
        ans = ((float64)(nums1[st1] + nums2[st2])) / 2
    } else {
        if ((st1 + st2 + 2 - 1) == (sz1 + sz2) % 2) {
            if (nums1[st1] > nums2[st2]) {
                ans = float64 (nums1[st1])
            } else {
                ans = float64 (nums2[st2])
            }
        } else {
            if (nums1[st1] > nums2[st2]) {
                ans = float64 (nums2[st2])
            } else {
                ans = float64 (nums1[st1])
            }
        }
    }
    return ans
}



func main_LT0004_20211005() {
// func main() {

    // var arr1 []int = []int{1,2}
    // var arr2 []int = []int{3,4}

    // arr1 := []int{1,3}
    // arr2 := []int{2,4}

    // arr1 := []int{1,3}
    // arr2 := []int{2}

    // var arr1 []int = []int{0,0}
    // arr2 := []int{1,1}

    // arr1 := []int{}
    // arr2 := []int{1}

    // arr1 := []int{2}
    // arr2 := []int{}

    // arr1 := []int{}
    // arr2 := []int{2,3}

    arr1 := []int{2,2,4,4}
    arr2 := []int{2,2,4,4}

    fmt.Println("ans:")

    ans := findMedianSortedArrays(arr1, arr2)

    fmt.Println(ans)

}
