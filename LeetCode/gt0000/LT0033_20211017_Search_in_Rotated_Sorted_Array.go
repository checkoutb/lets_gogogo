// package sdq
package main

import (
    "fmt"
)



// [12, 13, 14, 15, 16, 17, 18, 19, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] 可以当作：
//     [12, 13, 14, 15, 16, 17, 18, 19, inf, inf, inf, inf, inf, inf, inf, inf, inf, inf, inf
//     [-inf, -inf, -inf, -inf, -inf, -inf, -inf, -inf, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]
// while (lo < hi) {
//     int mid = (lo + hi) / 2;
//     double num = (nums[mid] < nums[0]) == (target < nums[0])
//                ? nums[mid]
//                : target < nums[0] ? -INFINITY : INFINITY;         
//     if (num < target)
//         lo = mid + 1;
//     else if (num > target)
//         hi = mid;
//     else
//         return mid;
// }
// cpp的， 第一次见到这个 INFINITY


// while(lo<hi){
//     int mid=(lo+hi)/2;
//     if(A[mid]>A[hi]) lo=mid+1;
//     else hi=mid;
// }
// 计算出 最小值。 然后
// int mid=(lo+hi)/2;
// int realmid=(mid+rot)%n;
// if(A[realmid]==target)return realmid;
// if(A[realmid]<target)lo=mid+1;
// else hi=mid-1;


// 好烦。。。
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
// Memory Usage: 2.6 MB, less than 38.66% of Go online submissions for Search in Rotated Sorted Array.
func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        t1, t2, t3 := nums[l], nums[mid], nums[r]

        // fmt.Println(l, mid, r, t2, t3)

        if target == t2 {
            return mid
        } else if target > t2 {
            if t1 < t3 {
                l = mid + 1
            } else {
                if t2 < t1 {
                    if target > t3 {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                } else {
                    l = mid + 1
                }
            }
        } else {
            if t1 < t3 {
                r = mid - 1
            } else {
                if t2 < t1 {
                    r = mid - 1
                } else {
                    if target > t3 {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
            }
        }
    }
    return -1 
}


func main_LT0033_20211017() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{4,5,6,7,0,1,2}
    // // tar := 0
    // tar := 3

    arr := []int{1}
    tar := 0

    ans := search(arr, tar)

    fmt.Println(ans)

}
