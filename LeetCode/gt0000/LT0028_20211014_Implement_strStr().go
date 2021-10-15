// package sdq
package main

import (
    "fmt"
)



// return strings.Index(haystack, needle)



// func getNext(next []int, s string) {
// 	j := -1 // j表示 最长相等前后缀长度
// 	next[0] = j
// 	for i := 1; i < len(s); i++ {
// 		for j >= 0 && s[i] != s[j+1] {
// 			j = next[j] // 回退前一位
// 		}
// 		if s[i] == s[j+1] {
// 			j++
// 		}
// 		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
// 	}
// }

// private void computeKMP(String s, int[] kmp){
//     int prefixEnd = -1;
//     int suffixEnd = 0;
//     kmp[0] = -1;
//     //while loop updates kmp[suffixEnd + 1]
//     while (suffixEnd < s.length() - 1) {
//         if (prefixEnd == -1 || s.charAt(prefixEnd) == s.charAt(suffixEnd)) {
//             kmp[suffixEnd +1] = prefixEnd + 1;
//             prefixEnd ++;
//             suffixEnd ++;
//         }
//         else {
//             prefixEnd = kmp[prefixEnd];   
//         }
//     }
// }


// Runtime: 200 ms, faster than 40.75% of Go online submissions for Implement strStr().
// Memory Usage: 2.6 MB, less than 51.65% of Go online submissions for Implement strStr().
func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    // .... ganjue kmp yao lai le, dan wo buhui
    mxi := len(haystack) - len(needle) + 1
    AAA:
    for i := 0; i < mxi; i++ {
        // b2 := true
        for j := 0; j < len(needle); j++ {
            if needle[j] != haystack[i + j] {
                continue AAA;
            }
        }
        return i
    }
    return -1
}


func main_LT0028_20211014() {
// func main() {

    fmt.Println("ans:")


}
