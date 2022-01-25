// package sdq
package main

import (
    "fmt"
)


// char res = 0;
// for(char c : s) res -= c;
// for(char c : t) res += c;
// return res;


// char ans=0;
// for(char ch:s)
//     ans^=ch;
// for(char ch:t)
//     ans^=ch;
// return ans;   



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find the Difference.
// Memory Usage: 2.5 MB, less than 20.88% of Go online submissions for Find the Difference.
func findTheDifference(s string, t string) byte {
    arr := make([]int, 123)
    for i := 0; i < len(s); i++ {
        arr[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        arr[t[i]]--
        if arr[t[i]] < 0 {
            return t[i]
        }
    }
    return s[0]
}


func main_LT0389_20220124() {
// func main() {

    fmt.Println("ans:")


}
