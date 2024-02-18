package sdq

// package main

import (
	"fmt"
)

// D D

// maxPos := min(steps/2, arrLen-1)
// dp := make([]int, maxPos+1)

// Runtime5 ms
// Beats
// 77.78%
// Memory7.1 MB
// Beats
// 55.56%

// step 500
// arrLen 10^6
// left, right, stay
// begin 0, end 0.  no outside
// first [0] [1]
// second [0] [1] [2] ...
func numWays(steps int, arrLen int) int {
    var MOD int = 1e9 + 7
    fmn := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    vi := make([]int, fmn(steps + 1, arrLen))
    vi[0] = 1

    for steps > 0 {
        steps--
        v2 := make([]int, len(vi))
        for i, cnt := range vi {
            if i > 0 {
                v2[i - 1] = (v2[i - 1] + cnt) % MOD
            }
            if i + 1 < len(vi) {
                v2[i + 1] = (v2[i + 1] + cnt) % MOD
            }
            v2[i] = (v2[i] + cnt) % MOD
        }
        vi = v2
    }
    return vi[0]
}


func main_LT1269_20231015() {
// func main() {

    fmt.Println("ans:")


}
