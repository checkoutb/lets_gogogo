// package sdq
package main

import (
    "fmt"
)


// ... KMP ... 



// The concatenated w1+w2 is palindrom if and only if:
//     len(w1)==len(w2) and the reversed w1 equals w2
//     len(w1)>len(w2) and the w1 must start with the reversed w2 and the remaining part of w1 should be a palindrome itself
//     len(w1)<len(w2) and w2 must end with the reversed w1 and the remaining part of w2 should be a palindrome itself
// 比硬算 跳过了 min(len1, len2)。 即 只需要计算 abs(len1 - len2)。

// if wrd[:-j] in rmap and wrd[-j:]==rev[:j]:
// go 能切片对比吗？


// Starting with the string "bot". We start checking all prefixes. 
// If "", "b", "bo", "bot" are themselves palindromes. 
// The empty string and "b" are palindromes. 
// We work with the corresponding suffixes ("bot", "ot") and check to see 
// if their reverses ("tob", "to") are present in our initial word list. 
// 对每个word，看每个 前缀 是否是一个 回文， 如果是，那么看 剩下的反转 后 是否 存在，存在 就能组成 回文。
// 消耗 在于 需要计算 每个前缀 是否是 回文。 空 和 一个char 必然是，



// 。。。anyway, accepted....
// Runtime: 1979 ms, faster than 6.06% of Go online submissions for Palindrome Pairs.
// Memory Usage: 131.5 MB, less than 24.24% of Go online submissions for Palindrome Pairs.
type MyTrie struct {
    Idx int
    Nxt map[byte]*MyTrie
}

// 看关联了Trie。 HashTable。
// 难道 反转 然后 前缀？ 取几位前缀呢。长度小于前缀呢。  ""
// Trie，是不是 先 反序 全部放到 Trie，然后 正序 搜索Trie，如果 有 相同前缀，就 走到底 找到 是哪条，然后 判断下
// 但是 Trie。。
func palindromePairs(words []string) [][]int {
    root := &MyTrie{-1, map[byte]*MyTrie{}}
    for i, w := range words {
        insertToTrie336(w, len(w) - 1, root, i)
    }
    ans := [][]int{}
    for idx, w := range words {
        arr := findTrie336(w, 0, root)
        // fmt.Println(idx, w, arr)
        for _, j := range arr {
            if j == idx { continue }
            if len(w) >= len(words[j]) {
                if isPalindrome336(w, len(words[j]), len(w) - 1) {
                    ans = append(ans, []int{idx, j})
                }
            } else {
                // if isPalindrome336(words[j], len(w), len(words[j]) - 1) {            // 反序的。。所以是 前面
                if isPalindrome336(words[j], 0, len(words[j]) - len(w) - 1) {
                    ans = append(ans, []int{idx, j})
                }
            }
        }
    }
    return ans
}

func isPalindrome336(w string, st, en int) bool {
    for st < en {
        if w[st] != w[en] {
            return false
        }
        st++
        en--
    }
    return true
}

// ssllllll  llss
// llllllss  ssll   碰到的都要记录。。
func findTrie336(word string, idx int, prev *MyTrie) []int {
    if prev == nil {
        return []int{}
    }
    // pIdx := prev.Idx
    ans := []int{}
    if prev.Idx != -1 {
        ans = append(ans, prev.Idx)
    }
    if idx >= len(word) {
        for _, v := range prev.Nxt {
            ans = append(ans, findTrie336(word, idx, v)...)
        }
    } else {
        ch := word[idx]
        if _, ok := prev.Nxt[ch]; ok {
            ans = append(ans, findTrie336(word, idx + 1, prev.Nxt[ch])...)
        }
    }
    return ans
}

// idx : len(word) - 1
func insertToTrie336(word string, idx int, prev *MyTrie, wordIdx int) {
    // if idx == len(word) {
    if idx == -1 {
        prev.Idx = wordIdx
        return
    }
    ch := word[idx]
    if _, ok := prev.Nxt[ch]; !ok {
        node := &MyTrie{-1, map[byte]*MyTrie{}}
        prev.Nxt[ch] = node
    }
    insertToTrie336(word, idx - 1, prev.Nxt[ch], wordIdx)
}


func main_LT0336_20220117() {
// func main() {

    fmt.Println("ans:")

    // arr := []string{"abcd","dcba","lls","s","sssll"}
    // arr := []string{"bat","tab","cat"}
    arr := []string{"a", ""}

    ans := palindromePairs(arr)

    fmt.Println(ans)

}
