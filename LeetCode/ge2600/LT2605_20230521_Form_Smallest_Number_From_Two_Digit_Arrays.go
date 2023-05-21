// package sdq
package main

import (
    "fmt"
)

// D D

// sort.Ints(nums2)


// Runtime0 ms
// Beats
// 100%
// Memory2.1 MB
// Beats
// 84.51%
func minNumber(nums1 []int, nums2 []int) int {
    var mn = 100
    for _, num := range nums1 {
        for _, num2 := range nums2 {
            if num == num2 {
                if (mn > num) {
                    mn = num
                }
            }
        }
    }

    if (mn < 100) {
        return mn;
    }

    var mn1 = mina1(nums1)
    var mn2 = mina1(nums2)

    if (mn1 > mn2) {
        return mn2 * 10 + mn1;
    } else {
        return mn1 * 10 + mn2;
    }

}

func mina1(arr []int) int {
    ans := arr[0]
    for _, num := range arr {
        if ans > num {
            ans = num
        }
    }
    return ans
}

// func main_LT2605_20230521() {
func main() {

    arr := []int{4,1,3}
    arr2 := []int{1,7}

    fmt.Println("ans:", minNumber(arr, arr2))


}
