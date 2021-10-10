// package sdq
package main

import (
    "fmt"
)

// if(value(s[i])>= value(s[i+1])){
//     ans = ans + value(s[i]);
// }
// else
// {
//     ans = ans - value(s[i]);
// }
// value 就是 char - int 的映射。

// for i in range(0, len(s) - 1):
// if roman[s[i]] < roman[s[i+1]]:
//     z -= roman[s[i]]
// else:
//     z += roman[s[i]]
// return z + roman[s[-1]]

// unordered_map<char, int> T = { { 'I' , 1 },
// { 'V' , 5 },
// { 'X' , 10 },
// { 'L' , 50 },
// { 'C' , 100 },
// { 'D' , 500 },
// { 'M' , 1000 } };


// numMap:=map[string]int{
//     "I": 1,
//     "V": 5,
//     "X": 10,
//     "L": 50,
//     "C": 100,
//     "D": 500,
//     "M": 1000,
// }

// table := map[rune]int{
//     'I': 1,
//     'V': 5,
//     'X': 10,
//     'L': 50,
//     'C': 100,
//     'D': 500,
//     'M': 1000,
// }

// runes := []rune(s)
// for i :=0 ;i < len(runes)-1; i++ {
//     val, _ := vals[runes[i]]
//     val2, _ := vals[runes[i+1]]

// Runtime: 4 ms, faster than 91.29% of Go online submissions for Roman to Integer.
// Memory Usage: 3.1 MB, less than 46.57% of Go online submissions for Roman to Integer.
func romanToInt(s string) int {
    arr1 := []int{1,5,10,50,100,500,1000}       // 0-6
    arr2 := []string{"I", "V", "X", "L", "C", "D", "M"}
    ans := 0
    sz1 := len(s)
    idx := 6
    for i := 0; i < sz1; i++ {
        for arr2[idx][0] != s[i] {
            idx--
        }
        if idx != 6 && i + 1 < sz1 && s[i + 1] == arr2[idx + 1][0] {
            ans += (arr1[idx + 1] - arr1[idx])
            i++
        } else if idx < 5 && i + 1 < sz1 && s[i + 1] == arr2[idx + 2][0] {
            ans += (arr1[idx + 2] - arr1[idx])
            i++
        } else {
            // fmt.Println(idx)
            ans += arr1[idx]
        }
        // fmt.Println(ans)
    }
    return ans
}


func main_LT0013_20211006() {
// func main() {

    fmt.Println("ans:")

    // s := "III"
    // s := "IV"
    // s := "IX"
    // s := "LVIII"
    s := "MCMXCIV"

    fmt.Println(romanToInt(s))

}
