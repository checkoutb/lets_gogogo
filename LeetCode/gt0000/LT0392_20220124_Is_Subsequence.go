// package sdq
package main

import (
    "fmt"
)


// def isSubsequence(self, s, t):
//     t = iter(t)
//     return all(c in t for c in s)


// while (*t)
//     s += *s == *t++;
// return !*s;


// int sLen = s.length(), sIdx = 0, tLen = t.length();
// for (int i=0; i<tLen && sIdx<sLen; i++) 
//     if (t[i]==s[sIdx]) sIdx++;
// return sIdx==sLen;


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Is Subsequence.
// Memory Usage: 2 MB, less than 56.83% of Go online submissions for Is Subsequence.
func isSubsequence(s string, t string) bool {
    si, ti, sz1, sz2 := 0, 0, len(s), len(t)
    for ; si < sz1; si++ {
        for ti < sz2 && s[si] != t[ti] {
            ti++
        }
        if ti < sz2 {
            ti++
        } else {
            return false
        }
    }
    return true
}


func main_LT0392_20220124() {
// func main() {

    fmt.Println("ans:")


}
