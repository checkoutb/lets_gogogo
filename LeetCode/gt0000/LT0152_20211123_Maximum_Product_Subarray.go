// package sdq
package main

import (
    "fmt"
)


// int globalMAx = arr[0], p = arr[0], q = arr[0];
// int crr_p = arr[0], crr_q = arr[0];
// for (int i = 1; i < arr.length; i++) {
//     crr_p = p * arr[i];
//     crr_q = q * arr[i];
//     p = Math.max(Math.max(crr_q, crr_p), arr[i]);
//     q = Math.min(Math.min(crr_q, crr_p), arr[i]);
//     globalMAx = Math.max(globalMAx, p);
// }
// return globalMAx;
// md。。。


// ii, ia := nums[0], nums[0]
// mi, ma := ii, ia
// for _, v := range nums[1:] {
//     ii, ia = minmax(v, ii*v, ia*v)
//     if ii < mi { mi = ii }
//     if ia > ma { ma = ia }
// }
// return ma


    // // store the result that is the max we have found so far
    // int r = A[0];
    // // imax/imin stores the max/min product of
    // // subarray that ends with the current number A[i]
    // for (int i = 1, imax = r, imin = r; i < n; i++) {
    //     // multiplied by a negative makes big number smaller, small number bigger
    //     // so we redefine the extremums by swapping them
    //     if (A[i] < 0)
    //         swap(imax, imin);
    //     // max/min product for the current number is either the current number itself
    //     // or the max/min by the previous number times the current one
    //     imax = max(A[i], imax * A[i]);
    //     imin = min(A[i], imin * A[i]);
    //     // the newly computed max value is a candidate for our global result
    //     r = max(r, imax);
    // }
    // return r;


    // int n = A.size(), res = A[0], l = 0, r = 0;
    // for (int i = 0; i < n; i++) {
    //     l =  (l ? l : 1) * A[i];
    //     r =  (r ? r : 1) * A[n - 1 - i];
    //     res = max(res, max(l, r));
    // }
    // return res;
// 如果没有0,那么最终结果，必然 包含 第一个，或者 包含最后一个。  所以 能直接整乘。




// Runtime: 5 ms, faster than 21.34% of Go online submissions for Maximum Product Subarray.
// Memory Usage: 2.9 MB, less than 11.17% of Go online submissions for Maximum Product Subarray.
// 10^4 个   [-10, 10] 并且 任何product <= INT_MAX
func maxProduct(nums []int) int {
    sz1 := len(nums)
    arr := make([]int, sz1 + 1)         // last is 0
    // arr[0] = nums[0]
    // for i := 1; i < sz1; i++ {
    //     // arr[i] = arr[i - 1] * nums[i]
    //     if nums[i] == 0 {
    //         arr[i] = 0
    //     } else {

    //     }
    // }

    pre := 1
    for i := 0; i < sz1; i++ {
        if nums[i] == 0 {
            pre = 1
        } else {
            arr[i] = pre * nums[i]
            pre = arr[i]
        }
    }
    // // 感觉O(n^2) 会tle。 或许可以记录 max  min。 毕竟除法 都会变小， 而且还有0可以跳过.
    // arr2 := make([][]int, 2)
    // arr2[0] = make([]int, sz1)      // [当前下标，最后] max    不是最后，是 最后或0
    // arr2[1] = make([]int, sz1)      // [当前，最后] min
    // arr2[0][sz1 - 1], arr2[1][sz1 - 1] = arr[sz1 - 1], arr[sz1 - 1]
    // // ... 0 .
    // mn, mx := arr[sz1 - 1], arr[sz1 - 1]
    // for i := sz1 - 1; i >= 0; i-- {
    //     if arr[i] == 0 {   // or nums[i] == 0
    //         if i > 0 {
    //             mn, mx = arr[i - 1], arr[i - 1]
    //         }
    //     } else {

    //     }
    // }
    
    // fmt.Println(arr)

    // bububu, 这里只有整数， 所以 必然是 非0 的最后一个 最大 或 最小。  所以就是两头找， 各自找到 一正一负 就可以了。
    // bu, 不是第一个。  比如  1,-6,-1  第一个负数是-6, 但是 最小负数是 。。。不，这里是找 arr的。 而不是 num的。 所以是 第一个，最后一个。
    pn,pp,sn,sp := 0,0,0,0      // prefix.negate ....  suffix.positive
    ans := nums[0]

    for i := 0; i <= sz1; i++ {
        if arr[i] == 0 {            // last 0 触发
            if sn != 0 && (sn / pn) > ans {
                ans = sn / pn
            }
            // if pp != 0 && (sp / pp) > ans {      // ... 正数 不需要 除法。。。 不需要 pp
            //     ans = sp / pp
            // }
            if sp != 0 && sp > ans {
                ans = sp
            }
            if i != sz1 && 0 > ans {
                ans = 0
            }
            pn,pp,sn,sp = 0,0,0,0
        } else if arr[i] < 0 {
            if pn == 0 {
                pn = arr[i]
            } else {            // 不然 [-1] 就会得出  -1/-1 == 1
                sn = arr[i]
            }
        } else {
            if pp == 0 {
                pp = arr[i]
            }
            sp = arr[i]
        }
    }
    return ans

    // pre := 1
    // for i := 0; i < sz1; i++ {
    //     if nums[i] == 0 {
    //         pre = 1

    //     } else {
    //         arr[i] = pre * nums[i]
    //         pre = arr[i]
    //         if pn == 0 && pre < 0 {
    //             pn = pre
    //         }
    //         if pp == 0 && pre > 0 {
    //             pp = pre
    //         }
    //         if pre < 0 {
    //             sn = pre
    //         }
    //         if pre > 0 {
    //             sp = pre
    //         }
    //     }
    // }
}


func main_LT0152_20211123() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{2,3,-2,4}
    arr := []int{-2,0,-1,-2}
    // arr := []int{-2}

    ans := maxProduct(arr)

    fmt.Println(ans)

}
