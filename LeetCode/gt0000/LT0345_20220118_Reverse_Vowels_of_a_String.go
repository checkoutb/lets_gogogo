// package sdq
package main

import (
    "fmt"
    "reflect"
)


// return re.sub('(?i)[aeiou]', lambda m, v=re.findall('(?i)[aeiou]', s): v.pop(), s)

// def reverseVowels(self, s):
//     vowels = re.findall('(?i)[aeiou]', s)
//     return re.sub('(?i)[aeiou]', lambda m: vowels.pop(), s)


// Runtime: 12 ms, faster than 26.17% of Go online submissions for Reverse Vowels of a String.
// Memory Usage: 4.2 MB, less than 71.03% of Go online submissions for Reverse Vowels of a String.
// 'a', 'e', 'i', 'o', and 'u',
func reverseVowels(s string) string {
    st, en := 0, len(s) - 1
    sb := []byte(s)
    for st < en {
        for st < en && !isVowel345(sb[st]) {
            st++
        }
        for st < en && !isVowel345(sb[en]) {
            en--
        }
        if st < en {
            sb[st], sb[en] = sb[en], sb[st]         // string is immutable...
            st++
            en--
        }
    }
    return string(sb)
}

func isVowel345(ch byte) bool {
    vs := "aeiouAEIOU"
    for i, _ := range vs {
        if ch == vs[i] {
            return true
        }
    }
    return false
}

func main_LT0345_20220118() {
// func main() {

    fmt.Println("ans:")

    // s := "hello"
    s := "leetcode"

    ans := reverseVowels(s)

    fmt.Println(ans)
    fmt.Println(reflect.TypeOf(ans))
    fmt.Println(reflect.TypeOf(ans[0]))
    fmt.Println(reflect.TypeOf("a"))
    fmt.Println(reflect.TypeOf('a'))        // int32

    for _, v := range ans {
        fmt.Println(reflect.TypeOf(v), v)       // int32
    }

}
