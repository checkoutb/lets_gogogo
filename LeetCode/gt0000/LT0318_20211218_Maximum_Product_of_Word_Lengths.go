// package sdq
package main

import (
    "fmt"
)


// if n, ok := m[bitmasks[i]]; ok {
//     m[bitmasks[i]] = Max(n, len(word))
// } else {
//     m[bitmasks[i]] = len(word)
// }


// bitset<32>wordBitset;
// for(auto ch : word) {
//     wordBitset.set(ch - 'a');
// }
// if((bitsetList[i] & bitsetList[j]).none()){



// Runtime: 30 ms, faster than 21.43% of Go online submissions for Maximum Product of Word Lengths.
// Memory Usage: 7.1 MB, less than 10.71% of Go online submissions for Maximum Product of Word Lengths.
// 316 的 bitmask...
func maxProduct(words []string) int {
    sz1 := len(words)
    mask := make([]int, sz1)
    // arr := make([]int, sz1)          // len(xx) is O(1), i remember
    ans := 0
    for i := 0; i < sz1; i++ {
        t2 := 0
        for j := 0; j < len(words[i]); j++ {
            t2 |= (1 << int(words[i][j] - 'a'))
        }
        mask[i] = t2
        for j := 0; j < i; j++ {
            if mask[j] & t2 == 0 && len(words[j]) * len(words[i]) > ans {
                ans = len(words[j]) * len(words[i])
            }
        }
    }
    return ans
    // for i := 0; i < sz1; i++ {
    // }
}

func main_LT0318_20211218() {
// func main() {

    fmt.Println("ans:")


}
