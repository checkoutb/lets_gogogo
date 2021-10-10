// package sdq
package main

import (
    "fmt"
    // "strconv"
)


// const MIN_INT = math.MinInt32
// const MAX_INT = math.MaxInt32



// Runtime: 6 ms, faster than 9.15% of Go online submissions for Reverse Integer.
// Memory Usage: 2.4 MB, less than 16.80% of Go online submissions for Reverse Integer.
// go int 看os的，64位的。。
// 负数范围 比正数 多1。
// 只能变大，不能变小。
func reverse(x int) int {
    var x2 int32 = int32(x)
    var ans int32 = 0
    var a2 int32 = 1
    if x2 < 0 {
        x2 = -x2
        a2 = -1
    }
    if x2 < 0 {         // INT_MIN
        return 0
    }
    for x2 > 0 {
        // fmt.Println(x2)
        // fmt.Println(ans)
        t2 := x2 % 10
        t3 := ans
        x2 /= 10
        ans *= 10
        ans += t2
        // // fmt.Println("  " + strconv.Itoa(ans))
        // fmt.Println("   " + fmt.Sprint(ans))
        // fmt.Println(t3)
        // if ans < t3 {
        if ans / 10 != t3 {
            return 0
        }
    }
    return int(ans * a2)
}


func main_LT0007_20211006() {
// func main() {

    fmt.Println("ans:")

    // x := 123
    // x := -321
    // x := 120
    // x := ^(int(^uint(0) >> 1))
    var x int = int(^(int32(^uint32(0) >> 1))) + 3

    fmt.Println(reverse(x))

}
