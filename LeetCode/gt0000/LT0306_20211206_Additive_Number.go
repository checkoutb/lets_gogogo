// package sdq
package main

import (
    "fmt"
)



// for for 计算前2位， 然后 dfs



// 最近几道，疯狂报错。。。
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Additive Number.
// Memory Usage: 2 MB, less than 61.11% of Go online submissions for Additive Number.
// 拆分出数字， f(n) = f(n-1) + f(n-2)
func isAdditiveNumber(s string) bool {
    t2, t3 := 0, 0
    num := make([]int, len(s))
    all0 := true
    for i := 0; i < len(s); i++ {
        num[i] = int(s[i] - '0')
        if num[i] != 0 {
            all0 = false
        }
    }
    if len(num) < 3 {
        return false
    }
    if all0 {
        return true
    }
    sz1 := len(num) / 2
    for i := 0; i < sz1; i++ {
        t2 = 0
        for a := 0; a <= i; a++ {
            t2 *= 10
            t2 += num[a]
        }
        for j := 1; j <= sz1; j++ {
            if i + j >= len(s) {
                break
            }
            t3 = 0
            // if num[i + 1] == '0' {
            //     break
            // }
            for a := i + 1; a <= i + j; a++ {
                t3 *= 10
                t3 += num[a]
            }
            // fmt.Println(" - ", t2, t3)
            k := i + j + 1
            if k >= len(s) {
                break
            }
            if num[k] == 0 {
                continue
            }
            t1 := t2
            t5 := 0
            for ; k < len(num); k++ {
                if t5 < t1 + t3 {
                    t5 *= 10
                    t5 += num[k]
                } else if t5 == t1 + t3 {
                    t1,t3 = t3,t5
                    t5 = num[k]
                    if t5 == 0 {
                        break
                    }
                } else {
                    break
                }
                // fmt.Println(t1, t3, t5)
            }
            if k == len(num) && t5 == t1 + t3 {
                return true
            }
            if num[i + 1] == 0 {
                break
            }
        }
        if num[0] == 0 {
            break
        }
    }
    return false
}

// func dfsa306()


func main_LT0306_20211206() {
// func main() {

    fmt.Println("ans:")

    // s := "112358"
    // s := "199100599"
    // s := "0000000"
    // s := "111122335588143"      // 11 11 22 33 55 88 143
    s := "121474836472147483648"        // 1+21yi=21yi



    fmt.Println(isAdditiveNumber(s))

}
