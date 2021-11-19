// package sdq
package main

import (
    "fmt"
    "sort"
)





// def wordBreak(self, s, words):
//     ok = [True]
//     for i in range(1, len(s)+1):
//         ok += any(ok[j] and s[j:i] in words for j in range(i)),
//     return ok[-1]



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Break.
// Memory Usage: 2.2 MB, less than 57.36% of Go online submissions for Word Break.
// dp, 不标准的dp。。。 memo。。好吧，也是标准dp。   top-down == dfs+memo 。  bottom-up == [][]int or []int
func wordBreak(s string, wordDict []string) bool {
    sz1 := len(s)
    // arr := make([][]int, sz1)
    // for i := 0; i < sz1; i++ {
    //     arr[i] = make([]int, sz1)
    // }
    map2, map3 := map[string]bool{}, map[int]bool{}
    for _, v := range wordDict {
        map2[v] = true
        map3[len(v)] = true
    }
    arrLen := []int{}
    for k := range map3 {
        arrLen = append(arrLen, k)
    }

    vst := make([]bool, sz1 + 1)
    vst[0] = true
    for i := 0; i < sz1; i++ {
        if vst[i] {
            for _, v := range arrLen {
                if i + v > sz1 {
                    continue
                }
                if vst[i + v] {
                    continue
                }
                t2 := s[i : i + v]
                if _, ok := map2[t2]; ok {
                    if i + v == sz1 {
                        return true
                    }
                    vst[i + v] = true
                }
            }
        }
    }
    return false
}

// tle * 2
// trie
func wordBreak2(s string, wordDict []string) bool {
    map2 := make(map[string]bool)
    map3 := make(map[int]bool)
    for _, v := range wordDict {
        map2[v] = true
        map3[len(v)] = true
    }
    arr := []int{}
    for k := range map3 {           // keys
        arr = append(arr, k)
    }
    sort.Ints(arr)
    return dfsa139(s, 0, map2, arr)
}

func dfsa139(s string, idx int, map2 map[string]bool, arr []int) bool {
    if idx == len(s) {
        return true
    }
    // for i := idx + 1; i <= len(s); i++ {
    for i := len(arr) - 1; i >= 0; i-- {
        en := idx + arr[i]
        if en > len(s) {
            continue
        }
        t2 := s[idx : en]
        if _, ok := map2[t2]; ok {
            if dfsa139(s, en, map2, arr) {
                return true
            }
        }
    }
    return false
}


func main_LT0139_20211119() {
// func main() {

    fmt.Println("ans:")


}
