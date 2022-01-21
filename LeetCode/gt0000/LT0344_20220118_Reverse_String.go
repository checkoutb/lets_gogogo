// package sdq
package main

import (
    "fmt"
)


// for x:= 0; x< l/2; x++{
//     s[l-1-x], s[x] = s[x], s[l-1-x]
// } 



// Runtime: 65 ms, faster than 11.87% of Go online submissions for Reverse String.
// Memory Usage: 6.6 MB, less than 88.22% of Go online submissions for Reverse String.
func reverseString(s []byte)  {
    st, en := 0, len(s) - 1
    for st < en {
        s[st], s[en] = s[en], s[st]
        st, en = st + 1, en - 1
    }
}

func main_LT0344_20220118() {
// func main() {

    fmt.Println("ans:")


}
