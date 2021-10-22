// package sdq
package main

import (
    "fmt"
)


// func mySqrt(x int) (r int) {             // 返回 (r int)
// 	r = x
// 	for r*r > x {
// 		r = (r + x/r) / 2
// 	}
// 	return
// }

// 牛顿

// long r = x;
// while (r*r > x)
//     r = (r + x/r) / 2;
// return r;


// Runtime: 4 ms, faster than 57.55% of Go online submissions for Sqrt(x).
// Memory Usage: 2.2 MB, less than 48.11% of Go online submissions for Sqrt(x).
// x < INT_MAX,   t2 * t2 > INT_MAX ...   but int64...    or < 0   or < last Product
func mySqrt(x int) int {
    t2 := 1
    for t2 * t2 <= x {
        t2++
    }
    return t2 - 1
}


func main_LT0069_20211022() {
// func main() {

    fmt.Println("ans:")


}
