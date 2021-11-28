// package sdq
package main

import (
    "fmt"
)


// trie
// . 就26个字母走一边







// Runtime: 93 ms, faster than 42.98% of Go online submissions for Design Add and Search Words Data Structure.
// Memory Usage: 12.6 MB, less than 93.86% of Go online submissions for Design Add and Search Words Data Structure.
type WordDictionary struct {
    Map2 map[int][]string
}
func Constructor() WordDictionary {
    return WordDictionary{ Map2: map[int][]string{} }    
}
func (this *WordDictionary) AddWord(word string)  {
    this.Map2[len(word)] = append(this.Map2[len(word)], word)
}
func (this *WordDictionary) Search(word string) bool {
    for _, v := range this.Map2[len(word)] {
        for i, v2 := range word {
            if v2 != '.' {
                if v2 != rune(v[i]) {
                    goto AAA
                }
            }
        }
        return true
        AAA:
        continue
    }
    return false
}


func main_LT0211_20211127() {
// func main() {

    fmt.Println("ans:")

    // var arr []int
    // fmt.Println(arr)

}
