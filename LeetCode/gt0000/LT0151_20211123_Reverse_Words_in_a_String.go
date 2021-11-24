// package sdq
package main

import (
    "fmt"
)



// var c []string
// var result string
// c = strings.Fields(s)
// for i := len(c) - 1; i >= 0 ; i-- {
//     result += c[i] + " "
// }
// return strings.Trim(result, " ")


// String[] words = s.trim().split(" +");
// Collections.reverse(Arrays.asList(words));
// return String.join(" ", words);
// " +" 是RE。 >=1个空白


// 应该是 j记录i， 内部for中i--
// Runtime: 6 ms, faster than 32.72% of Go online submissions for Reverse Words in a String.
// Memory Usage: 8 MB, less than 18.20% of Go online submissions for Reverse Words in a String.
func reverseWords(s string) string {
    ans := ""
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            j := i - 1
            for j >= 0 && s[j] != ' ' {
                j--
            }
            ans += s[j + 1 : i + 1]
            ans += " "
            i = j
        }
    }
    return ans[0 : len(ans) - 1]
}

func main_LT0151_20211123() {
// func main() {

    fmt.Println("ans:")


}
