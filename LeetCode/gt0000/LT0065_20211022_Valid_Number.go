// package sdq
package main

import (
    "fmt"
)


// https://leetcode.com/problems/valid-number/discuss/23728/A-simple-solution-in-Python-based-on-DFA
// DFA deterministic finite automation 确定有限自动机
// 有个状态转换图





// boolean number = false, decimal = false, epsilon = false, sign = false;       
// for (int i = 0; i < s.length(); i++) {
//     char c = s.charAt(i);
//     if (Character.isDigit(c)) {
//         number = true;
//     }
//     else if (c == 'e' || c == 'E') {
//         if (!number || epsilon) return false;
//         epsilon = true;
//         number = false;
//     }
//     else if (c == '.') {
//         if (epsilon || decimal) return false;
//         decimal = true;
//     }
//     else if (c == '+' || c == '-') {
//         if (i > 0 && Character.toLowerCase(s.charAt(i-1)) != 'e') return false;
//     }
//     else return false;
// }
// return number;

// var states = []map[string]int {
//     {"digit": 1, "sign": 2, "dot": 3},
//     {"digit": 1, "dot": 4, "exp": 5},
//     {"digit": 1, "dot": 3},
//     {"digit": 4},
//     {"digit": 4, "exp": 5},
//     {"sign": 6, "digit": 7},
//     {"digit": 7},
//     {"digit": 7},
// }

// var endStates = map[int]struct{} {
//     1: struct{}{},
//     4: struct{}{},
//     7: struct{}{},
// }

// func isNumber(s string) bool {
//     curr := 0
//     var transition string
//     for i := range s {
//         switch {
//             case s[i] >= '0' && s[i] <= '9':
//             transition = "digit"
//             case s[i] == '.':
//             transition = "dot"
//             case s[i] == '+' || s[i] == '-':
//             transition = "sign"
//             case s[i] == 'e' || s[i] == 'E':
//             transition = "exp"
//             default:
//             return false
//         }
//         if next, ok := states[curr][transition]; ok {
//             curr = next
//         } else {
//             return false
//         }
//     }
//     _, ok := endStates[curr]
//     return ok
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Number.
// Memory Usage: 2.3 MB, less than 48.84% of Go online submissions for Valid Number.
// 按e划分
// +- 只在第一位
// .前/后 必须有值， 最多一个
// e后面不能有.
// 1.1000 也是 valid吧
func isNumber(s string) bool {
    idx := -1
    for i := 0; i < len(s); i++ {
        if s[i] == 'e' || s[i] == 'E' {
            idx = i
            break
            // if idx != -1 {
            //     return false
            // }
            // idx = i
        }
    }
    s1 := s
    s2 := "1"
    if idx != -1 {
        s1 = s[0 : idx]
        s2 = s[idx + 1 : len(s)]
        if len(s2) == 0 {       // 1e
            return false
        }
        if len(s1) == 0 {       // e3
            return false
        } else if len(s1) == 1 {            // +e3

        }
    }
    
    fun_isInt := func(str string) bool {
        if len(str) == 0 {          // 1e+
            return false
        }
        i := 0
        // if str[i] == '-' || str[i] == '+' {             //  +1.+1    外面判断 +-。 
        //     i++
        // }
        for ; i < len(str); i++ {
            if str[i] > '9' || str[i] < '0' {
                return false
            }
        }
        return true
    }
    if s2[0] == '-' || s2[0] == '+' {
        s2 = s2[1 : len(s2)]
    }
    if !fun_isInt(s2) {
        return false
    }

    idx = -1
    for i := 0; i < len(s1); i++ {              // .
        if s1[i] == '.' {
            idx = i
            break
        }
    }
    s1 = s1
    s2 = "1"
    if idx != -1 {              // 1.   ok
        s2 = s1[idx + 1 : len(s1)]
        s1 = s1[0 : idx]
        if len(s1) == 0 && len(s2) == 0 {       // .
            return false
        }
    }
    if len(s2) != 0 {               // 1. ok      1.+1  bad
        if !fun_isInt(s2) {
            return false
        }
    }
    if len(s1) != 0 {           // .123  ok
        if s1[0] == '+' || s1[0] == '-' {           // +.123 ok
            s1 = s1[1 : len(s1)]
        }
        if len(s1) != 0 {           // +.123
            if !fun_isInt(s1) {
                return false
            }
        } else {
            if len(s2) == 0 {           // +.e3
                return false
            }
            if idx == -1 {          // +e3
                return false
            }
        }
    }
    return true
}





func main_LT0065_20211022() {
// func main() {

    fmt.Println("ans:")

    arr := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
                "0", ".1"}
    
    for i := 0; i < len(arr); i++ {
        ans := isNumber(arr[i])
        if !ans {
            fmt.Println(arr[i])
        }
    }

    fmt.Println("==============================")

    arr = []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53","e",".","+e3","+.e3"}

    for i := 0; i < len(arr); i++ {
        ans := isNumber(arr[i])
        if ans {
            fmt.Println(arr[i])
        }
    }

}
