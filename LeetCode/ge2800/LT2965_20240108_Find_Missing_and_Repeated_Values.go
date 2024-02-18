package sdq
// package main

import (
    "fmt"
)





// Runtime19 ms
// Beats
// 48.26%
// Memory6.4 MB
// Beats
// 71.14%

func findMissingAndRepeatedValues(grid [][]int) []int {
    // sum2 := func(arr []int) int {
    //     var ans = 0
    //     for _, i := range arr {
    //         ans += i
    //     }
    //     return ans
    // }
    // var t2 = 0;
    // for _, arr := range grid {
    //     t2 += sum2(arr);
    // }
    
    var sz1 = len(grid)
    var vb = make([]int, sz1 * sz1 + 1)
    
    for i := range grid {
        for _, num := range grid[i] {
            vb[num] += 1;
        }
    }

    // for i := range
    idx := 1
    ans := make([]int, 2);
    for idx <= sz1 * sz1 {
        if vb[idx] == 2 {
            ans[0] = idx
        } else if vb[idx] == 0 {
            ans[1] = idx;
        }
        idx += 1
    }
    return ans
}


func main_LT2965_20240108() {
// func main() {

    fmt.Println("ans:")

    vvi := [][]int{{1,3},{2,2}};

    vv2 := findMissingAndRepeatedValues(vvi);

    fmt.Println(vv2);
}
