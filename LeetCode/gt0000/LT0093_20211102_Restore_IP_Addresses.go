// package sdq
package main

import (
    "fmt"
    "strings"
)



// for(int i = 1; i<4 && i<len-2; i++){
//     for(int j = i+1; j<i+4 && j<len-1; j++){
//         for(int k = j+1; k<j+4 && k<len; k++){
//             String s1 = s.substring(0,i), s2 = s.substring(i,j), s3 = s.substring(j,k), s4 = s.substring(k,len);
//             if(isValid(s1) && isValid(s2) && isValid(s3) && isValid(s4)){
//                 res.add(s1+"."+s2+"."+s3+"."+s4);
//             }
//         }
//     }
// }



// Runtime: 24 ms, faster than 7.89% of Go online submissions for Restore IP Addresses.
// Memory Usage: 6.3 MB, less than 5.26% of Go online submissions for Restore IP Addresses.
func restoreIpAddresses(s string) []string {
    ans := []string{}
    dfsa93(&s, 0, []string{}, &ans)
    return ans
}

func dfsa93(s *string, idx int, arr []string, ans *[]string) {
    if idx == len(*s) {
        if len(arr) == 4 {
            t2 := strings.Join(arr, ".")
            *ans = append(*ans, t2)
        }
        return
    }
    if len(arr) == 4 {
        return
    }
    t2 := 0
    arr = append(arr, "")
    for idx < len(*s) {
        t2 *= 10
        t2 += int((*s)[idx] - '0')

        idx++
        if t2 <= 255 {
            arr[len(arr) - 1] = arr[len(arr) - 1] + (*s)[idx - 1 : idx]
            dfsa93(s, idx, arr, ans)
        }
        if t2 == 0 {
            break
        }
    }
    arr = arr[0 : len(arr) - 1]
}

func main_LT0093_20211102() {
// func main() {

    fmt.Println("ans:")

    // s := "25525511135"
    // s := "1111"
    // s := "0000"
    // s := "010010"
    s := "101023"

    ans := restoreIpAddresses(s)

    fmt.Println(ans)

}
