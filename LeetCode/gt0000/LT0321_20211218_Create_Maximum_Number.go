// package sdq
package main

import (
    "fmt"
)


// gg


// tle。
// 每个数 计算 它后面可以有多少个数， 取 后续多于(k-ans) 个的，且最大的。   多个呢？ 不不不，最多2个 每个数组一个，因为要最开始的。 就dfs下。
func maxNumber(nums1 []int, nums2 []int, k int) []int {
    ans := dfsa321(nums1, nums2, k)
    return ans
}

func dfsa321(arr1, arr2 []int, k int) []int {
    if k == 0 {
        return []int{}
    }
    sz1, sz2 := len(arr1), len(arr2)
    mxi1, mxi2 := 0, 0
    for i := 1; i < sz1; i++ {
        if sz1 - i + sz2 < k {
            break
        }
        if arr1[i] > arr1[mxi1] {
            mxi1 = i
        }
    }
    for i := 1; i < sz2; i++ {
        if sz2 - i + sz1 < k { break }
        if arr2[i] > arr2[mxi2] {
            mxi2 = i
        }
    }

    if sz1 == 0 {
        return append([]int{arr2[mxi2]}, dfsa321(arr1, arr2[mxi2 + 1 : ], k - 1)...)
    }
    if sz2 == 0 {
        return append([]int{arr1[mxi1]}, dfsa321(arr1[mxi1 + 1 : ], arr2, k - 1)...)
    }

    // ans := []int{}
    // if arr1[mxi1] >= arr2[mxi2] {    }
    // if arr1[mxi1] <= arr2[mxi2] {    }
    if arr1[mxi1] > arr2[mxi2] {
        ans := append([]int{arr1[mxi1]}, dfsa321(arr1[mxi1 + 1 : ], arr2, k - 1)...)
        return ans
    } else if arr1[mxi1] < arr2[mxi2] {
        ans := append([]int{arr2[mxi2]}, dfsa321(arr1, arr2[mxi2 + 1 : ], k - 1)...)
        return ans
    } else {
        // need compare
        a1 := dfsa321(arr1[mxi1 + 1 : ], arr2, k - 1)
        a2 := dfsa321(arr1, arr2[mxi2 + 1 : ], k - 1)
        for i := 0; i < len(a1); i++ {
            if a1[i] > a2[i] {
                return append([]int{arr1[mxi1]}, a1...)
            } else if a1[i] < a2[i] {
                return append([]int{arr2[mxi2]}, a2...)
            }
        }
        return append([]int{arr1[mxi1]}, a1...)
    }
}


// error
// 感觉是取最大且先出现，且剩下的数>剩余所需 的数，  如果2个相同呢。
// 或者 全用到 merge 组成一个 m+n的， 然后 删除最小？  应该是这个。不是最小，是删除 最前面出现的 上升。
func maxNumber2(nums1 []int, nums2 []int, k int) []int {
    arr := []int{}
    idx1, idx2 := 0, 0
    sz1, sz2 := len(nums1), len(nums2)
    for len(arr) < sz1 + sz2 {
        if idx1 == sz1 {
            for idx2 < sz2 {
                arr = append(arr, nums2[idx2])
                idx2++
            }
        } else if idx2 == sz2 {
            for idx1 < sz1 {
                arr = append(arr, nums1[idx1])
                idx1++
            }
        } else {
            if nums1[idx1] < nums2[idx2] {
                arr = append(arr, nums2[idx2])
                idx2++
            } else {
                arr = append(arr, nums1[idx1])
                idx1++
            }
        }
    }
    fmt.Println(arr)
    idx := 1
    for len(arr) > k && idx < len(arr) {
        if arr[idx] <= arr[idx - 1] {
            idx++
        } else {
            arr = append(arr[0 : idx - 1], arr[idx : ]...)
            if idx > 1 {
                idx--
            }
        }
    }
    if len(arr) > k {
        arr = arr[0 : k]
    }
    return arr
}

func main_LT0321_20211218() {
// func main() {

    fmt.Println("ans:")

    arr1 := []int{3,4,6,5}
    arr2 := []int{9,1,2,5,8,3}
    k := 5

    ans := maxNumber(arr1, arr2, k)
    
    fmt.Println(ans)

}
