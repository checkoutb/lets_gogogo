// package sdq
package main

import (
    "fmt"
    "math"
)


// 先 while(s[i] == ' ') i++
// 然后 if(s[i] == '+'/'-')
// 最后 while(s[i] is digit)


// sign = 1 - 2 * (str[i++] == '-'); 
// c的，根据 + / - 来决定 是 1 还是 -1


// Runtime: 3 ms, faster than 52.21% of Go online submissions for String to Integer (atoi).
// Memory Usage: 2.4 MB, less than 27.26% of Go online submissions for String to Integer (atoi).
func myAtoi(s string) int {
    var mn int = math.MinInt32
    var mx int = math.MaxInt32
    ans := 0
    t2 := 1
    b2 := false
    for i := 0; i < len(s); i++ {
        if (s[i] == ' ') {
            if (b2) {
                return ans
            }
            continue
        }
        if s[i] == '+' || s[i] == '-' {
            if (b2) {
                return ans
            }
            if s[i] == '-' {
                t2 = -1
            }
            if i + 1 == len(s) || s[i + 1] < '0' || s[i + 1] > '9' {
                return 0
            }
            continue
        }
        if s[i] < '0' || s[i] > '9' {
            if b2 {
                return ans
            } else {
                return 0
            }
        }
        b2 = true
        ans *= 10
        ans += t2 * int (s[i] - '0')
        if (ans < mn) {
            return mn
        }
        if ans > mx {
            return mx
        }
    }
    return ans
}

func main_LT0008_20211006() {
// func main() {

    fmt.Println("ans:")

    // s := "    -42"
    // s := "4193 with wo"
    // s := "asdasd 123132"
    // s := "-91283472332"
    s := "-5-"

    fmt.Println(myAtoi(s))
}
