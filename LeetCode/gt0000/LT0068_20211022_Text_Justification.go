// package sdq
package main

import (
    "fmt"
)






// Runtime: 4 ms, faster than 10.64% of Go online submissions for Text Justification.
// Memory Usage: 2.3 MB, less than 28.72% of Go online submissions for Text Justification.
func fullJustify(words []string, maxWidth int) []string {
    idx, sz1, mx1 := 0, len(words), maxWidth + 1
    ans := []string{}
    for idx < sz1 {
        i2 := idx
        sdq := 0        // size
        for idx < sz1 && sdq + len(words[idx]) + 1 <= mx1 {
            sdq += len(words[idx]) + 1
            idx++
        }
        sss := words[i2]
        i2++
        if idx == sz1 || i2 == idx {            // i2 == idx, just one word in this line, need "     " as suffix
            for ; i2 < idx; i2++ {
                sss += " " + words[i2]
            }
            for len(sss) < maxWidth {
                sss += " "
            }
        } else {
            t2 := (mx1 - sdq + idx - i2) / (idx - i2)
            t3 := (mx1 - sdq + idx - i2) % (idx - i2)
            // fmt.Println(t2, ", ", t3, ", ", (mx1 - sdq + idx - i2), ", ", (idx - i2))
            semp := ""
            for i := 0; i < t2; i++ {
                semp += " "
            }
            semp2 := semp + " "
            cnt := 0
            for ; i2 < idx; i2++ {
                if cnt < t3 {
                    sss += semp2 + words[i2]
                    cnt++
                } else {
                    sss += semp + words[i2]
                }
            }
        }
        ans = append(ans, sss)
    }
    return ans
}


func main_LT0068_20211022() {
// func main() {

    fmt.Println("ans:")

    // arr := []string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}
    // mx := 20

    // arr := []string{"What","must","be","acknowledgment","shall","be"}
    arr := []string{"This", "is", "an", "example", "of", "text", "justification."}
    mx := 16

    ans := fullJustify(arr, mx)

    // fmt.Println(ans)
    for _, s := range ans {
        fmt.Println("-", s, "-", len(s))
    }

}
