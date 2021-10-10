// package sdq
package main

import (
    "fmt"
)

// for(int idx=0; strs.size()>0; prefix+=strs[0][idx], idx++)

// for(int i = 1; i < strsSize ; i++){
//     int j = 0;
//     while(strs[i][j] && t[j] && strs[i][j]==t[j])
//         j++;
//     t[j] = '\0';
// }
// 和每个string 对比， 对比2者的最长。


// sort，对比第一个 和最后一个， 因为这2个是差距最大的。



// Runtime: 4 ms, faster than 18.10% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.7 MB, less than 14.91% of Go online submissions for Longest Common Prefix.
func longestCommonPrefix(strs []string) string {
    ans := ""
    idx := 0
    sz1 := len(strs)
    AAA:
    for true {
        for i := 0; i < sz1; i++ {
            if (len(strs[i]) == idx) {
                break AAA
            }
            if strs[i][idx] != strs[0][idx] {
                break AAA
            }
        }
        ans += string (strs[0][idx])
        idx++
    }
    return ans
}


func main_LT0014_20211006() {
// func main() {

    fmt.Println("ans:")

    // var vs []string = []string{"flower","flow","flight"}
    vs := []string{"dog","racecar","car"}

    fmt.Println(longestCommonPrefix(vs))

}
