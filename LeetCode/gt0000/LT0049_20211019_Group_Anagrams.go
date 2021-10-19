// package sdq
package main

import (
    "fmt"
    "sort"
)


// hashMap := make(map[[26]int][]string)
// map 作为key。


// type sortRunes []rune
// func (s sortRunes) Less(i, j int) bool {
//     return s[i] < s[j]
// }
// func (s sortRunes) Swap(i, j int) {
//     s[i], s[j] = s[j], s[i]
// }
// func (s sortRunes) Len() int {
//     return len(s)
// }
// func sortString(str string) string {
//     runes := []rune(str)
//     sort.Sort(sortRunes(runes))
//     return string(runes)
// }





// Runtime: 24 ms, faster than 87.53% of Go online submissions for Group Anagrams.
// Memory Usage: 8.7 MB, less than 48.49% of Go online submissions for Group Anagrams.
func groupAnagrams(strs []string) [][]string {
    ans := [][]string{}
    map2 := map[string][]string{}

    for _, s := range strs {
        barr := []byte(s)
        arr := make([]int, len(s))
        for i, v := range barr {
            arr[i] = int(v)
        }
        sort.Ints(arr)
        for i, v := range arr {
            barr[i] = byte(v)
        }
        ks := string(barr)
        map2[ks] = append(map2[ks], s)
    }

    for _, arr := range map2 {
        ans = append(ans, arr)
    }

    return ans
}


func main_LT0049_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := []byte{'1','3','2'}
    // arr2 := []int(arr)           // no
    // sort.Sort(arr)               // no

    // arr := []string{"eat","tea","tan","ate","nat","bat"}
    // arr := []string{"a"}
    arr := []string{""}

    ans := groupAnagrams(arr)

    for _, ar := range ans {
        for _, e := range ar {
            fmt.Print(e, ", ")
        }
        fmt.Println()
    }

}
