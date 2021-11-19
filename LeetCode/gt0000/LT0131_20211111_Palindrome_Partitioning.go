// package sdq
package main

import (
    "fmt"
)



// dp[start][end] = true;


// def partition(self, s):
//     return [[s[:i]] + rest
//             for i in xrange(1, len(s)+1)
//             if s[:i] == s[i-1::-1]
//             for rest in self.partition(s[i:])] or [[]]

// def partition(s)
//   s == '' ? [[]] : s.size.times.flat_map { |i| s[0..i] != s[0..i].reverse ? [] :
//     partition(s[i+1..-1]).map { |rest| [s[0..i]] + rest }
//   }
// end


// Runtime: 296 ms, faster than 71.55% of Go online submissions for Palindrome Partitioning.
// Memory Usage: 26 MB, less than 31.90% of Go online submissions for Palindrome Partitioning.
func partition(s string) [][]string {
    ans := [][]string{}
    dfsa131(s, 0, []string{}, &ans)
    return ans
}

func dfsa131(s string, idx int, arr []string, ans *[][]string) {
    if idx >= len(s) {
        arr2 := make([]string, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    for i := idx + 1; i <= len(s); i++ {
        s2 := s[idx : i]            // .
        if isPalind(s2) {
            arr = append(arr, s2)
            dfsa131(s, i, arr, ans)
            arr = arr[0 : len(arr) - 1]
        }
    }
}

func isPalind(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}

func main_LT0131_20211111() {
// func main() {

    fmt.Println("ans:")

    s := "aab"

    ans := partition(s)

    fmt.Println(ans)

}
