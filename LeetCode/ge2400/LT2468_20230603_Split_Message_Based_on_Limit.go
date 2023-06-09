// package sdq
package main

import (
    "fmt"
    "strconv"
)

// D D


// mLen := len(message)

// b := 1
// aLen := sz(1)

// for b*limit < b*(sz(b)+3)+aLen+mLen {
//     if sz(b)*2+3 >= limit {
//         return []string{}
//     }
//     b++
//     aLen = aLen + sz(b)
// }

// rs := make([]string, 0)
// i := 0
// for a := 1; a <= b; a++ {
//     var sb strings.Builder
//     j := limit - (3 + sz(a) + sz(b))
//     sb.WriteString(message[i:min(i+j, mLen)])
//     sb.WriteRune('<')
//     sb.WriteString(strconv.Itoa(a))
//     sb.WriteRune('/')
//     sb.WriteString(strconv.Itoa(b))
//     sb.WriteRune('>')
//     rs = append(rs, sb.String())
//     i = i + j
// }




// Runtime120 ms
// Beats
// 100%
// Memory7.3 MB
// Beats
// 100%

// < 10
// < 100 < 1000
// 无法直接二分吧？ 99段可以，100段不行，应该也是可能的。
func splitMessage(message string, limit int) []string {
    sz1 := len(message)
    ans := []string{}

    count := -1
    for i := 1; i < 7; i++ {
        cnt := 0
        t2 := limit - i - 3
        t3 := 1
        t4 := 1
        for j := 1; j <= i; j++ {
            t4 *= 10
            cnt += (t2 - j) * (t4 - t3)
            t3 = t4
        }
fmt.Println("cnt ", cnt, t4)
        if cnt >= sz1 {
            cnt -= (t2 - i) * (t4 - t4 / 10)
            cnt = sz1 - cnt
            count = t4 / 10 + (cnt - 1 + t2 - i) / (t2 - i)
            // if t4 == 10 {
                count--
            // }
            break
        }
    }
fmt.Println("count ", count)
    if count == -1 {
        return ans
    }
    
    sfx := "/" + strconv.Itoa(count) + ">"
    
    idx := 0
    for i := 1; i <= count; i++ {
        s := "<" + strconv.Itoa(i) + sfx
        idx2 := idx + limit - len(s)
        if idx2 > sz1 {
            idx2 = sz1
        }
        s = message[idx : idx2] + s
        idx = idx2
        ans = append(ans, s)
    }

    return ans
}


func main_LT2468_20230603() {
// func main() {

    fmt.Println("ans:")

    // s := "short message"
    // lmt := 15

    s := "this is really a very awesome message"
    lmt := 9


    fmt.Println(splitMessage(s, lmt))

}
