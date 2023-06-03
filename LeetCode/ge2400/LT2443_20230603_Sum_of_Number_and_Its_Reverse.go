// package sdq
package main

import (
    "fmt"
)

// brute force


// g

func sumOfNumberAndReverse(num int) bool {
    vi := []int{}
    for num > 0 {
        vi = append(vi, num % 10)
        num /= 10
    }
    sz1 := len(vi)
    // for i := 0; i < sz1 / 2; i++ {
    //     vi[i], vi[sz1 - 1 - i] = vi[sz1 - 1 - i], vi[i]
    // }

    // [0] is low

    carl := 0
    carh := 0
    for i := 0; i < sz1 / 2; i++ {
        lw, hg := vi[i], vi[sz1 - 1 - i]
        lw -= carl

    }
    if sz1 % 2 == 1 {

    }


}


func main_LT2443_20230603() {
// func main() {

    fmt.Println("ans:")


}
