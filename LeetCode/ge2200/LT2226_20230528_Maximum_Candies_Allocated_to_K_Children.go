// package sdq
package main

import (
    "fmt"
)

// D
// start, end:=int64(0), sum/k


// Runtime249 ms
// Beats
// 33.33%
// Memory10.8 MB
// Beats
// 20%

func maximumCandies(candies []int, k int64) int {

    var st, en int = 1, maxa1(candies)
    var ans int = 0
    for st <= en {
        md := (st + en) / 2;

        fmt.Println(st, en, md)

        var t2 int64 = 0
        for _, n := range candies {
            t2 += int64(n / md)         // st == 1, 防止 md == 0
            if t2 >= k {
                break
            }
        }
        if (t2 >= k) {
            ans, st = md, md + 1
        } else {
            en = md - 1
        }
    }
    return ans
}

func maxa1(arr []int) int {
    ans := 0
    for _, n := range arr {
        if n > ans {
            ans = n
        }
    }
    return ans;
}

func suma1(arr []int) int64 {
    var ans int64 = 0
    for _, n := range arr {
        ans += int64(n)
    }
    return ans
}


func main_LT2226_20230528() {
// func main() {

    arr := []int{5,8,6}
    k := int64(4)

    fmt.Println("ans:", maximumCandies(arr, k))


}
