// package sdq
package main

import (
    "fmt"
)

// D D


// int dp[51] = {};
// for (int i = s.size() - 1; i >= 0; --i) {
//     dp[i] = 1 + dp[i + 1];
//     for (const auto &w: dict)
//         if (s.compare(i, w.size(), w) == 0)          // ...
//             dp[i] = min(dp[i], dp[i + w.size()]);
// }
// return dp[0];




// Runtime26 ms
// Beats
// 100%
// Sorry, there are not enough accepted submissions to show data
// Memory7.7 MB
// Beats
// 100%
func minExtraChar(s string, dictionary []string) int {
    
    sz2 := len(dictionary)
    sz1 := len(s)
    vi := make([]int, sz1)

    process1 := func(idx int) {
        if idx == 0 {
            vi[idx] = 1
        } else {
            vi[idx] = vi[idx - 1] + 1
        }
        v2 := make([]int, sz2)
        for i, ss := range dictionary {
            v2[i] = len(ss) - 1
        }
        chg := true
        offset := 0
        for chg && idx >= offset {
            chg = false
            for i := range v2 {
                if (v2[i] != -1) {
                    chg = true
                    if (s[idx - offset] == dictionary[i][v2[i]]) {
                        v2[i]--
                        if v2[i] == -1 {
                            if idx - offset == 0 {
                                vi[idx] = 0
                                return
                            }
                            if vi[idx - offset - 1] < vi[idx] {
                                vi[idx] = vi[idx - offset - 1]
                            }
                        }
                    } else {
                        v2[i] = -1
                    }
                }
            }
            offset++
        }
    }

    for i := range s {
        process1(i)
    }
    return vi[sz1 - 1]
}


func main_LT2707_20230528() {
// func main() {

    s := "le1etscode"
    arr := []string{"leet","code","leetcode"}

    fmt.Println("ans:", minExtraChar(s, arr))


}
