// package sdq
package main

import (
    "fmt"
)



// Runtime0 ms
// Beats
// 100%
// Sorry, there are not enough accepted submissions to show data
// Memory2.6 MB
// Beats
// 100%

// .. {1,2},{1,2} t=1 ...
// math.
func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    tx, ty := sx - fx, sy - fy
    if tx < 0 {
        tx = -tx
    }
    if ty < 0 {
        ty = -ty
    }
    if tx == 0 && ty == 0 && t == 1 {
        return false;
    }
    if tx > ty {
        tx = tx + ty - ty
    } else  {
        tx = tx + ty - tx   // ... tx = max(tx, ty)
    }
    return tx <= t
}


func main_LT2849_20231010() {
// func main() {

    fmt.Println("ans:")


}
