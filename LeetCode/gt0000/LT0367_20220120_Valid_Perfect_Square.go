// package sdq
package main

import (
    "fmt"
)



// long long lo = 1, hi = num;

// 牛顿方法 求平方跟。


// sqrt(2147483647)   <    46341
func isPerfectSquare(num int) bool {
    l, r := 0, 46340
    for l <= r {
        mid := (l + r) / 2
        if mid * mid == num {
            return true
        } else if mid * mid > num {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false
}

func main_LT0367_20220120() {
// func main() {

    fmt.Println("ans:")


}
