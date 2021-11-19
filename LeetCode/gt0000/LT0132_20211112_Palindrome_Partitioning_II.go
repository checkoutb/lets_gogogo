// package sdq
package main

import (
    "fmt"
)




// for(int i = 0; i <= n; i++) dp[i] = i - 1;//initialize the value for each dp state.
// for(int i = 2; i <= n; i++){
//     for(int j = i - 1; j >= 0; j--){
//         //if(isPalindr[j][i]){
//         if(s.charAt(i - 1) == s.charAt(j) && (i - 1 - j < 2 || isPalindr[j + 1][i - 1])){
//             isPalindr[j][i] = true;
//             dp[i] = Math.min(dp[i], dp[j] + 1);
//         }
//     }
// }


// Manancher-like
// for(i=1;i<N;i++)
// {
//     for(j=0;(i-j)>=0 && (i+j)<N && s[i-j]== s[i+j]; ++j) // odd-length substrings 
//         minCUTS[i+j+1] = min(minCUTS[i+j+1], 1 + minCUTS[i-j]);
//     for(j=0;(i-j-1)>=0 && (i+j)<N && s[i-j-1]== s[i+j]; ++j) // even-length substrings
//         minCUTS[i+j+1] = min(minCUTS[i+j+1], 1 + minCUTS[i-j-1]);
// }




// 估计是 2个dp 可以合并的。 一次dp就可以了。
// Runtime: 140 ms, faster than 5.56% of Go online submissions for Palindrome Partitioning II.
// Memory Usage: 36.4 MB, less than 5.56% of Go online submissions for Palindrome Partitioning II.
// dp + dp ,  稀疏矩阵
func minCut(s string) int {
    sz1 := len(s)
    dp1 := make([][]int, sz1)
    // fmt.Println(dp1[0])
    tmp := make([]map[int]bool, sz1)
    for i, _ := range tmp {
        tmp[i] = make(map[int]bool)
    }
    for i := sz1 - 1; i >= 0; i-- {
        for j := i + 1; j < sz1; j++ {      // needn't  s[i:i+1]
            if s[i] == s[j] {
                t2, t3 := i + 1, j - 1
                b2 := false         // is palindrome
                if t2 >= t3 {
                    b2 = true
                } else {
                    if _, ok := tmp[t2][t3]; ok {
                        b2 = true
                    }
                }
                if b2 {
                    tmp[i][j] = true
                    dp1[i] = append(dp1[i], j)
                }
            }
        }
    }

    dp2 := make([]int, sz1)         // min count/(cut + 1) in s[idx : ]
    for i, _ := range dp2 {     // make 能设置初始值？
        dp2[i] = (sz1 - i)      // s[idx : ]    not s[idx+1 : ]
    }

    for i := sz1 - 1; i >= 0; i-- {
        // t2 := 
        for j := len(dp1[i]) - 1; j >= 0; j-- {
            t2 := dp1[i][j] + 1
            if t2 == sz1 {
                dp2[i] = 1
                break
            } else {
                if (1 + dp2[t2]) < dp2[i] {
                    dp2[i] = dp2[t2] + 1
                }
            }
        }
        // .. add s[i:i]
        if (i + 1) != sz1 && (1 + dp2[i + 1]) < dp2[i] {
            dp2[i] = dp2[i + 1] + 1
        }

        // for _, j := range dp1[i] {
        //     t2 := j + 1
        //     if t2 == sz1 {

        //     }
        // }
    }
    
    fmt.Println(dp2, dp1)

    return dp2[0] - 1
}

func main_LT0132_20211112() {
// func main() {

    fmt.Println("ans:")

    // s := "aab"
    // s := "a"
    // s := "ab"
    s := "cdd"

    ans := minCut(s)

    fmt.Println(ans)

}
