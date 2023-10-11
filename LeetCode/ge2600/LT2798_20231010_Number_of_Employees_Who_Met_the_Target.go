// package sdq
package main

import (
    "fmt"
)



// Runtime3 ms
// Beats
// 64%
// Memory2.7 MB
// Beats
// 95.14%
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    var ans = 0
    for _, val := range hours {
        if val >= target {
            ans += 1
        }
    }
    return ans
}


func main_LT2798_20231010() {
// func main() {

    fmt.Println("ans:")


}
