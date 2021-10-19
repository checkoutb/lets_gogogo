// package sdq
package main

import (
    "fmt"
)




// const char* star=NULL;
// const char* ss=s;
// while (*s){
//     //advancing both pointers when (both characters match) or ('?' found in pattern)
//     //note that *p will not advance beyond its length 
//     if ((*p=='?')||(*p==*s)){s++;p++;continue;} 

//     // * found in pattern, track index of *, only advancing pattern pointer 
//     if (*p=='*'){star=p++; ss=s;continue;} 

//     //current characters didn't match, last pattern pointer was *, current pattern pointer is not *
//     //only advancing pattern pointer
//     if (star){ p = star+1; s=++ss;continue;} 

//    //current pattern pointer is not star, last patter pointer was not *
//    //characters do not match
//     return false;
// }
// //check for remaining characters in pattern
// while (*p=='*'){p++;}
// return !*p;  



// 或者 从后往前， 计算 后续各个char的个数。 然后开始从头开始对比。
// 不行，这个 结尾没有办法。


// tle.
// 感觉只会dfs， 估计tle的。
func isMatch(s string, p string) bool {
    for i := 1; i < len(p); i++ {
        if p[i] == '*' && p[i] == p[i - 1] {
            p = p[0 : i] + p[i + 1 : len(p)]
            i--
        }
    }
    // fmt.Println(p)
    ans := dfsa44(s, p, 0, 0)
    return ans
}

func dfsa44(s string, p string, si int, pi int) bool {
    // fmt.Println(si, ", ", pi)
    if si == len(s) || pi == len(p) {
        return si == len(s) && (pi == len(p) || (pi == len(p) - 1 && p[pi] == '*'))
    }
    t2 := s[si]
    t3 := p[pi]
    if (t3 == '*') {
        if dfsa44(s, p, si + 1, pi) || dfsa44(s, p, si, pi + 1) || dfsa44(s, p, si + 1, pi + 1) {
            return true
        }
        // if dfsa44(s, p, si, pi + 1) {
        //     return true
        // }
        // if dfsa44(s, p, si + 1, pi + 1) {
        //     return true
        // }
        
    // } else if (t3 == '?') {

    } else {
        if t2 != t3 && t3 != '?' {
            return false
        }
        return dfsa44(s, p, si + 1, pi + 1)
    }
    return false
}


func main_LT0044_20211019() {
// func main() {

    fmt.Println("ans:")

    // s, p := "aa", "a"
    // s, p := "aa", "****"
    // s, p := "cb", "?a"
    // s, p := "abcde", "*a*e"
    // s, p := "acdcb", "a*c?b"
    s, p := "", "*****"

    ans := isMatch(s, p)

    fmt.Println(ans)

}
