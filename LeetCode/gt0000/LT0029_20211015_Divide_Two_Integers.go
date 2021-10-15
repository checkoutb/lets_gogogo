// package sdq
package main

import (
    "fmt"
    "math"
)

// if(dividend==1<<31 && divisor==-1) return Integer.MAX_VALUE;

// while(dividend-divisor>=0){
//     int count = 0;
//     while(dividend-(divisor<<1<<count)>=0){
//         count++;
//     }
//     result +=1<<count;
//     dividend-=divisor<<count;
// }


// if (A == INT_MIN && B == -1) return INT_MAX;
// int a = abs(A), b = abs(B), res = 0;
// for (int x = 31; x >= 0; x--)
//     if ((signed)((unsigned)a >> x) - b >= 0)
//         res += 1 << x, a -= b << x;
// return (A > 0) == (B > 0) ? res : -res;
// O(32)



// ... 题目要求 32bit的。。。

// Runtime: 10 ms, faster than 46.92% of Go online submissions for Divide Two Integers.
// Memory Usage: 2.8 MB, less than 24.62% of Go online submissions for Divide Two Integers.

// no * / %
// 无限累加？
func divide(dividend int, divisor int) int {
    ans := 0
    flag := 1
    if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
        flag = -1
    }
    if dividend < 0 {
        dividend = -dividend
    }
    if divisor < 0 {
        divisor = -divisor
    }
    // map2 := map[int]int
    // cnt := 1
    // for divisor < dividend {            // !int32  int64!
    //     map2[divisor] = cnt
    //     divisor += divisor
    //     cnt += cnt
    // }

    arrval := []int{}
    arrcnt := []int{}
    cnt := 1
    for divisor <= dividend {
        arrval = append(arrval, divisor)
        arrcnt = append(arrcnt, cnt)
        divisor += divisor
        cnt += cnt
    }

    lst := len(arrval) - 1

    for lst >= 0 {
        for dividend >= arrval[lst] {
            dividend -= arrval[lst]
            ans += arrcnt[lst]
        }
        lst--
    }
    ans = flag * ans
    if ans <= math.MaxInt32 && ans >= math.MinInt32 {
        return ans
    }
    if ans < 0 {
        return math.MinInt32
    } else {
        return math.MaxInt32
    }
}


func main_LT0029_20211015() {
// func main() {

    fmt.Println("ans:")

    fmt.Println(math.MaxInt32)

    // a := 10
    // b := 3

    // ans := divide(a, b)

    // arr := []int{7,-3}
    arr := []int{1,1}
    
    ans := divide(arr[0], arr[1])

    fmt.Println(ans)

}
