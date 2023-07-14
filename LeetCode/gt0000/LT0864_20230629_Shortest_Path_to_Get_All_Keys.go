// package sdq
package main

import (
    "fmt"
    // "unicode"
)






// Runtime11 ms
// Beats
// 59.9%
// Memory7.5 MB
// Beats
// 36.36%

// This is the Daily Coding Challenge for June 29th 2023. 
// There are 14 incomplete challenges for June 2023 and you have 1 ticket left for this month. 
// Are you sure you want to use a ticket to make up this submission?
// - of cource not, thank you

// 硬写写出来了，但是leetcode访问不了。????, 还好我一点儿也不急，因为这个是昨天的daily challenge。


// status, i, j, step
type Tuple struct {
    sta int
    i int
    j int
    stp int
}

// topic: bfs, bit op
func shortestPathAllKeys(grid []string) int {
    sz1 := len(grid);
    sz2 := len(grid[0]);

    // arr := make([]Tuple);
    arr := []Tuple{};
    var cnt int = 0;
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == '@' {
                arr = append(arr, Tuple{0, i, j, 0});
            // } else if unicode.IsUpper(grid[i][j]) {
            } else if grid[i][j] >= 'A' && grid[i][j] <= 'Z' {
                cnt++;
            }
        }
    }

    vb := make([][][]bool, sz1);
    var tot = 1 << cnt;
    for i := range vb {
        vb[i] = make([][]bool, sz2);
        for j := range vb[i] {
            vb[i][j] = make([]bool, tot);
        }
    }
    tot--;

    // return status == tup.sta
    fun1 := func(tup Tuple, i int, j int) int {
        if i < 0 || j < 0 || i == sz1 || j == sz2 || grid[i][j] == '#' {
            return -1;
        }
        // if unicode.IsUpper(grid[i][j]) && ((tup.sta & (1 << (grid[i][j] - 'A'))) == 0) {
        if grid[i][j] >= 'A' && grid[i][j] <= 'Z' && ((tup.sta & (1 << (grid[i][j] - 'A'))) == 0) {
            return -1;
        }
        st2 := 0;
        if grid[i][j] == '.' {
            st2 = tup.sta;
        }
        st2 = tup.sta | (1 << (grid[i][j] - 'a'));
        if vb[i][j][st2] {
            return -1;
        }
        vb[i][j][st2] = true;

        // fmt.Println("... ", st2)

        return st2;
    }

    for len(arr) > 0 {
        // var arr2 []Tuple = make([]Tuple);
        var arr2 []Tuple = []Tuple{};
        for _, v := range arr {
            if v.sta == tot{
                return v.stp;
            }
            
            // fmt.Println(v.sta, v.i, v.j, v.stp);

            st2 := fun1(v, v.i + 1, v.j);
            if st2 >= 0 {
                arr2 = append(arr2, Tuple{st2, v.i + 1, v.j, v.stp + 1});
            }
            st2 = fun1(v, v.i - 1, v.j);
            if st2 >= 0 {
                arr2 = append(arr2, Tuple{st2, v.i - 1, v.j, v.stp + 1});
            }
            st2 = fun1(v, v.i, v.j - 1);
            if st2 >= 0 {
                arr2 = append(arr2, Tuple{st2, v.i, v.j - 1, v.stp + 1});
            }
            st2 = fun1(v, v.i, v.j + 1);
            if st2 >= 0 {
                arr2 = append(arr2, Tuple{st2, v.i, v.j + 1, v.stp + 1});
            }
        }
        arr = arr2;
    }
    return -1;
}


func main_LT0864_20230629() {
// func main() {

    // vs := []string{"@.a..","###.#","b.A.B"};
    // vs := []string{"@..aA","..B#.","....b"};
    vs := []string{"@Aa"};

    fmt.Println("ans:", shortestPathAllKeys(vs));


}
