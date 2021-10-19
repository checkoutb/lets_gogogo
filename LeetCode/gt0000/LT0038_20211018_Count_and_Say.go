// package sdq
package main

import (
    "fmt"
)

// if n == 1 { return "1" }
// prev := countAndSay(n-1)
// var sb strings.Builder
// count := 0
// for i := range prev {
//     count++
//     if i == len(prev)-1 || prev[i+1] != prev[i] {
//         sb.WriteString(strconv.Itoa(count))
//         sb.WriteByte(prev[i])
//         count = 0;
//     }
// }
// return sb.String()


// sb.WriteString(fmt.Sprintf("%d%c", total, c))




// Runtime: 0 ms, faster than 100.00% of Go online submissions for Count and Say.
// Memory Usage: 2.9 MB, less than 68.57% of Go online submissions for Count and Say.
func countAndSay(n int) string {
    s := "1"
    for n > 1 {
        n--
        arr := make([]byte, 0)
        // if len(s) > 1 {
            lst := s[0]
            cnt := 1
            for i := 1; i < len(s); i++ {
                if (s[i] != lst) {
                    arr = append(arr, byte(cnt + '0'), lst)
                    cnt = 1
                    lst = s[i]
                } else {
                    cnt++
                }
            }
        // } else {
            arr = append(arr, byte(cnt + '0'), lst)
        // }
        s = string(arr)
        // fmt.Println(s)
    }
    return s
}

// func dfsa38(s string, n int) {
//     if 
// }


func main_LT0038_20211018() {
// func main() {

    fmt.Println("ans:")

    n := 4
    // n := 1

    ans := countAndSay(n)

    fmt.Println(ans)

}
