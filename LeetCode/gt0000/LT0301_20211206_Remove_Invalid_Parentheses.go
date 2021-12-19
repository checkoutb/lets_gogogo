// package sdq
package main

import (
    "fmt"
)





// gg
// 应该是 只要 > 0 ,就删除


// 感觉是 正向一遍，反向一遍。记录 ( - ) 的值
func removeInvalidParentheses(s string) []string {
    for len(s) > 0 && s[0] == ')' {
        s = s[1 : ]
    }
    for len(s) > 0 && s[len(s) - 1] == '(' {
        s = s[ : len(s) - 1]
    }
    // fmt.Println(". . . ", s)
    ans := []string{}
    dfsa301(s, 0, &ans)
    if len(ans) == 0 {
        ans = append(ans, "")
    }
    return ans
}

func dfsa301(s string, idx int, ans *[]string) {
    fmt.Println(s, idx)
    if idx == len(s) {
        dfsaaa301(s, len(s) - 1, ans)
        return
    }
    cnt := 0
    i := idx
    for ; i < len(s); i++ {
        if s[i] == '(' {
            cnt++
        } else if s[i] == ')' {
            cnt--
            if cnt < 0 {        // -1
                for j := idx; j <= i; j++ {
                    if s[j] == ')' {
                        dfsa301(s[0 : j] + s[j + 1 : ], i, ans)
                        for j + 1 <= i && s[j + 1] == ')' {
                            j++
                        }
                    }
                }
                break

                // b2 := true
                // for j := i; j >= idx; j++ {
                //     if s[j] == ')' {
                //         if b2 && cnt < 0 {
                //             dfsa301(s[0 : j] + s[j + 1 : ], i, ans)
                //             b2 = false
                //         }
                //         cnt++
                //     } else if s[j] == '(' {
                //         b2 = true
                //         cnt--
                //     }
                // }
                // break;
            }
        }
    }
    if i == len(s) && cnt >= 0 {
        dfsaaa301(s, len(s) - 1, ans)
    }
}

func dfsaaa301(s string, idx int, ans *[]string) {
    fmt.Println(" ... ", s, idx)
    if idx < 0 {
        *ans = append(*ans, s)
        return
    }
    cnt := 0
    i := idx
    for ; i >= 0; i-- {
        if s[i] == ')' {
            cnt++
        } else if s[i] == '(' {
            cnt--
            if cnt < 0 {
                for j := idx; j >= i; j-- {
                    if s[j] == '(' {
                        dfsaaa301(s[0 : j] + s[j + 1 : ], i - 1, ans)
                        for j - 1 >= i && s[j - 1] == '(' {
                            j--
                        }
                    }
                }
                break
            }
        }
    }
    if i < 0 && cnt == 0 {
        dfsaaa301(s, -1, ans)
    }
}

func main_LT0301_20211206() {
// func main() {

    fmt.Println("ans:")

    // s := "()())()"
    // s := "(a)())()"
    // s := ")("
    // s := "))"
    // s := "x("
    // s := "(()"
    s := "(r(()()("




    arr := removeInvalidParentheses(s)
    fmt.Println(arr, len(arr))

}
