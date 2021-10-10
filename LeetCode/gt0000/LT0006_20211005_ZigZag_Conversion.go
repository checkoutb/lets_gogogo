// package sdq
package main

import (
    "fmt"
    "bytes"
)

 
// var sb strings.Builder
// sb.WriteByte(s[i])

// ss := make([]uint8, len(s), len(s))
// return string(ss)

// split := strings.Split(s, "")
// slice := []string{}
// for i := 0; i < numRows; i++ {
//     slice = append(slice, "")
// }
// return strings.Join(slice,"")


// numRows 个sb， 然后开始 遍历 s， 如果 下标 % (numRows) == 0，开始向下， 如果 == numRows 向上。 down = !down.  
// 向下，行数++，向上，行数--， 行数就是 sb[行数] 。 最后再 sb想连接。


// j += numRows + numRows - 2
// j 是同一行的下一个值。



// Runtime: 8 ms, faster than 78.68% of Go online submissions for ZigZag Conversion.
// Memory Usage: 4.2 MB, less than 92.28% of Go online submissions for ZigZag Conversion.
func convert(s string, numRows int) string {
    if (numRows == 1) {
        return s
    }
    ans := ""
    sz1 := len(s)
    len2 := numRows * 2 - 2             // step
    var sb bytes.Buffer

    for i := 0; i < sz1; i += len2 {
        sb.WriteString(s[i : i + 1])            // zen me xie ?
    }

    for i := 1; i < numRows - 1; i++ {
        t2 := i
        t3 := len2 - i
        for t2 < sz1 {
            sb.WriteString(s[t2 : t2 + 1])
            if t3 < sz1 {
                sb.WriteString(s[t3 : t3 + 1])
            }
            t2 += len2
            t3 += len2
        }
    }

    // if (numRows != 1) {
        for i := numRows - 1; i < sz1; i += len2 {
            sb.WriteString(s[i : i + 1])
        }
    // }

    ans = sb.String()
    return ans
}


func main_LT0006_20211005() {
// func main() {

    fmt.Println("ans:")

    // s := "PAYPALISHIRING"
    // k := 3

    s := "PAYPALISHIRING"
    k := 2

    // s := "a"
    // k := 1

    fmt.Println(convert(s, k))

}
