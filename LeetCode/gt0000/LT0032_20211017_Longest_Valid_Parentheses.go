// package sdq
package main

import (
    "fmt"
)


// 碰到 ) ， 往前搜索 一个 (


func longestValidParentheses(s string) int {
    sz1, ans := len(s), 0

    return ans
}



// 好像可以， 直接硬算， 碰到 负数 就刷新。     error... ()(()
func longestValidParentheses3(s string) int {
    sz1 := len(s)
    ans := 0
    // arr := make([]bool, sz1)
    diff := 0           // (:+, )->-
    st := 0
    for i := 0; i < sz1; i++ {
        if s[i] == '(' {
            diff++
        } else {
            diff--
        }
        if (diff < 0) {
            if i - st > ans {
                ans = i - st
            }
            st = i + 1
            diff = 0
        }
    }
    if sz1 - st - diff > ans {
        ans = sz1 - st - diff
    }
    return ans
}



// tle...
// 匹配就是 (.cnt == ).cnt  && 任何prefix-str (.cnt >= ).cnt
// dp么？ 但是3*10^3 平方下。。  可以一维数组， 第一步找 ()，   后续就是 () + ()  或者  ( + () + ) ... 好像不能一维。。不会。。
func longestValidParentheses2(s string) int {
    sz1 := len(s)
    ans := 0

    dp := make([][]bool, sz1)
    for i := range dp {
        dp[i] = make([]bool, sz1)
    }
    
    for i := 0; i < sz1 - 1; i++ {
        // dp[i][i] = true
        if s[i] == '(' && s[i + 1] == ')' {
            dp[i][i + 1] = true
            ans = 2
        }
    }

    // ans2 := -1
    // 如果 某次sz2 没有 新的产生，可以直接退出。           ... error  ")(((((()())()()))()(()))("  22
    for sz2 := 4; sz2 <= sz1; sz2 += 2 {
        // if (ans2 == ans) {
        //     break
        // }
        // ans2 = ans
        for i := 0; i + sz2 <= sz1; i++ {
            if dp[i + 1][i + sz2 - 2] && s[i] == '(' && s[i + sz2 - 1] == ')' {
                dp[i][i + sz2 - 1] = true
                ans = sz2
                continue
            }
            for j := i + 1; j < i + sz2 - 1; j += 2 {
                if dp[i][j] && dp[j + 1][i + sz2 - 1] {
                    dp[i][i + sz2 - 1] = true
                    ans = sz2
                    continue
                }
            }
        }
    }
    return ans
}


func main_LT0032_20211017() {
// func main() {

    fmt.Println("ans:")

    // s := ")(((((()())()()))()(()))("
    // s := ""
    // s := ")()())"
    // s := "(()"
    s := "()(()"

    ans := longestValidParentheses(s)

    fmt.Println(ans)
}
