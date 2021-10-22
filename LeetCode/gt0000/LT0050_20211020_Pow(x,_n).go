// package sdq
package main

import (
    "fmt"
)


// dfs
// if math.Mod(float64(n), 2) == 0 {
//     n := helper(x, n/2)
//     return n * n
// }
// return helper(x, n-1) * x





// INT_MIN， 可以直接变成INT_MAX， 还有 比较 上次循环的 ans 和本次ans ，是否可以提前退出。
// cpp 里 -INT_MIN == INT_MIN

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pow(x, n).
// Memory Usage: 2.3 MB, less than 13.27% of Go online submissions for Pow(x, n).
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    b2 := true          // n > 0 ?
    if n < 0 {
        n = -n
        b2 = false
    }
    arr := []float64{x}
    t2 := 1
    for t2 < n {
        lst := arr[len(arr) - 1]
        arr = append(arr, lst * lst)
        t2 *= 2
    }
    // fmt.Println(arr)
    var ans float64 = 1.0
    idx := len(arr) - 1
    for n > 0 {
        // fmt.Println(1<<idx, ", ", float64((int(1) << (idx))), ", ", x)
        if n >= (1 << idx) {         // 1 是 float。。无法shift，  float64 和int 无法比较，无法 四则运算。。 bu, 之前用了 x >= ，现在用 n>= 后 就不需要了， 应该是 自动推导成了 float64,和左操作数一致。
            n -= 1 << idx
            if b2 {
                ans *= arr[idx]
            } else {
                ans /= arr[idx]
            }
            // fmt.Println(ans, ". ", idx)
        } else {
            idx--
        }
    }
    return ans
}


func main_LT0050_20211020() {
// func main() {

    fmt.Println("ans:")

    // x := 2.00000
    // n := 10

    // x := 2.1
    // n := 3

    // x := 2.0
    // n := -2

    x := 1.0
    n := -2100000000

    ans := myPow(x, n)

    fmt.Println(ans)

}
