// package sdq
package main

import (
    "fmt"
    "sort"
    "math"
)

// D D

// c := []int{}
// for _, cut := range cuts {
//     c = append(c, cut)
// }
// c = append(c, []int{0, n}...)
// sort.Ints(c)
// dp := make([][]int, len(c))

// for i := range dp {
//     dp[i] = make([]int, len(c))
// }

// for i := len(c) - 1; i >= 0; i-- {
//     for j := i + 1; j < len(c); j++ {
//         for k := i + 1; k < j; k++ {



// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }




// Runtime24 ms
// Beats
// 75%
// Memory4.6 MB
// Beats
// 75%

func minCost(n int, cuts []int) int {
    sz1 := len(cuts);
    sort.Ints(cuts);

    vvi := make([][]int, sz1);
    for i := 0; i < sz1; i++ {
        vvi[i] = make([]int, sz1);
        for j := 0; j < sz1; j++ {
            vvi[i][j] = -1;
        }
    }
    return dfs1(cuts, vvi, 0, sz1 - 1, 0, n);
}

func dfs1(cuts []int, vvi [][]int, st int, en int, stk_st int, stk_en int) int {
    if (st > en) {
        return 0;
    } else if (vvi[st][en] != -1) {
        return vvi[st][en];
    } else if (st == en) {
        vvi[st][en] = stk_en - stk_st;
        return vvi[st][en];
    }
    var t2 int = 0;
    var mn int = math.MaxInt32;
    for i := st; i <= en; i++ {
        t2 = dfs1(cuts, vvi, st, i - 1, stk_st, cuts[i]) + dfs1(cuts, vvi, i + 1, en, cuts[i], stk_en);
        if t2 < mn {
            mn = t2;
        }
    }
    vvi[st][en] = mn + stk_en - stk_st;
    return vvi[st][en];
}


func main_LT1547_20230528() {
// func main() {

    fmt.Println("ans:")

    var n int = 9;
    arr := []int{5,6,1,4,2};

    fmt.Println(minCost(n, arr));


}
