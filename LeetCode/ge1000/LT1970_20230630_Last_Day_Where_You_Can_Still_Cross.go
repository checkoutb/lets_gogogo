// package sdq
package main

import (
    "fmt"
)


// D D

// 看details，都是小于300ms的，估计运气问题。 不知道这道题目多少是tle。

// 看detail，都是 uf，但没看懂。

// discuss uf 6


// Runtime1965 ms
// Beats
// 33.33%
// Memory12 MB
// Beats
// 33.33%


// 0 land, 1 water
// day 0, all are land
// one day, one cell be water
// top row -> bottom row, by land.      Binary Search ?
// if vvi[0][0] cannot go to any vvi[sz1 - 1][x] in day 3, in the day > 3, not need to check it.
// so for each vvi[0][x], use binary search to check which day is the last day that i can goto bottom from vvi[0][x] ?
// for each vvi[0][x] { binary search } or binary search { for each vvi[0][x] }
// first is better ?
func latestDayToCross(row int, col int, cells [][]int) int {
    var vvi [][]int = make([][]int, row);
    for i := range vvi {
        vvi[i] = make([]int, col);
    }
    for i, vi := range cells {
        vvi[vi[0] - 1][vi[1] - 1] = i;
    }

    // fmt.Println(vvi);

    var vvb[][]bool = make([][]bool, row);      // vst
    for i := range vvb {
        vvb[i] = make([]bool, col);
    }

    // var arr []int = []int{1,0,-1,0,1};
    // dfsa1 := func(i int, j int, lmt int) bool { // < lmt,  >= lmt is land
    //     if i < 0 || j < 0 || i == row || j == col || vvb[i][j] || vvi[i][j] < lmt {
    //         return false;
    //     }
    //     vvb [i][j] = true;
    //     for k := 0; k < 4; k++ {
    //         if dfsa1(i + arr[k], j + arr[k + 1], lmt) {
    //             return true;
    //         }
    //     }
    //     return false;
    // }

    filla1 := func() {
        for i := range vvb {
            for j := range vvb[i] {
                vvb[i][j] = false;
            }
        }
    }
    ans := 0;
    for j := range vvi[0] {
        if vvi[0][j] < ans {
            continue;
        }
        st := ans + 1;
        en := len(cells);       // should be vvi[0][j] (maybe + 1)
        for st <= en {
            md := (st + en) / 2;
            // fmt.Println(" .. ", st, en, md);
            filla1();       // ...fang for wai mian le...
            if dfsa1(vvi, vvb, 0, j, md, row, col) {
                // fmt.Println("???", md)
                if md > ans {
                    ans = md;
                }
                st = md + 1;
            } else {
                en = md - 1;
            }
        }
    }
    return ans;
}
var arr []int = []int{1,0,-1,0,1};
func dfsa1(vvi [][]int, vvb [][]bool, i int, j int, lmt int, row, col int) bool { // < lmt,  >= lmt is land

    // fmt.Println(&vvi[0][0], &vvb[0][0]);     // == ref

    if i < 0 || j < 0 || i == row || j == col || vvb[i][j] || vvi[i][j] < lmt {
        return false;
    }
    // fmt.Println(i, j, lmt, vvb[i][j], vvi[i][j]);
    if i == row - 1 {
        // fmt.Println("------!!!");
        return true;
    }

    vvb [i][j] = true;
    for k := 0; k < 4; k++ {
        if dfsa1(vvi, vvb, i + arr[k], j + arr[k + 1], lmt, row, col) {
            return true;
        }
    }
    return false;
}


// func main_LT1970_20230630() {
func main() {

    r := 3;
    c := 3;
    vvi := [][]int{{1,2},{2,1},{3,3},{2,2},{1,1},{1,3},{2,3},{3,2},{3,1}};

    fmt.Println("ans:", latestDayToCross(r, c, vvi));


}
