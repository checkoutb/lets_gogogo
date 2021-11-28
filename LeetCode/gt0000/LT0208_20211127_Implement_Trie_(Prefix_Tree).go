// package sdq
package main

import (
    "fmt"
)


// word[i]            byte
// _ ch := range word      rune





// Runtime: 109 ms, faster than 11.32% of Go online submissions for Implement Trie (Prefix Tree).
// Memory Usage: 19.4 MB, less than 10.00% of Go online submissions for Implement Trie (Prefix Tree).
type Trie struct {
    Val rune
    End bool
    Next [26]*Trie
}
func Constructor() Trie {
    return Trie{ '@', false, [26]*Trie{} }
}
func (this *Trie) Insert(word string)  {
    t2 := this
    for _, ch := range word {
        if t2.Next[ch - 'a'] == nil {
            t2.Next[ch - 'a'] = &Trie{ ch, false, [26]*Trie{} }
        }
        t2 = t2.Next[ch - 'a']
    }
    t2.End = true
}
func (this *Trie) Search(word string) bool {
    t2 := this
    for _, ch := range word {
        if t2.Next[ch - 'a'] == nil {
            return false
        }
        t2 = t2.Next[ch - 'a']
    }
    return t2.End
}
func (this *Trie) StartsWith(prefix string) bool {
    t2 := this
    for _, ch := range prefix {
        if t2.Next[ch - 'a'] == nil {
            return false
        }
        t2 = t2.Next[ch - 'a']
    }
    return true
}



func main_LT0208_20211127() {
// func main() {

    fmt.Println("ans:")


}
