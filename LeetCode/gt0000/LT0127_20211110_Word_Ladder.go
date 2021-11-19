// package sdq
package main

import (
    "fmt"
)


// 对自己的每个char进行替换。


// wSet := make(map[string]bool)
// delete(wSet, beginWord)
// bSet := map[string]bool{beginWord:true}


// Runtime: 212 ms, faster than 48.03% of Go online submissions for Word Ladder.
// Memory Usage: 7.2 MB, less than 63.82% of Go online submissions for Word Ladder.
// map. bfs 。只会硬算
func ladderLength(beginWord string, endWord string, wordList []string) int {
    map2 := map[string][]string{}
    sz1 := len(wordList)
    b2 := false
    for i, s := range wordList {
        for j := i + 1; j < sz1; j++ {
            if diffOne127(s, wordList[j]) {
                map2[s] = append(map2[s], wordList[j])
                map2[wordList[j]] = append(map2[wordList[j]], s)
            }
        }
        if !b2 && s == endWord {
            b2 = true
        }
    }
    if !b2 {
        return 0
    }
    // beMap := map[string]int{}
    que1 := []string{}
    enMap := map[string]int{}
    gotMap := map[string]int{}      // visited
    for _, s := range wordList {
        if diffOne127(beginWord, s) {
            // beMap[s] = 1
            que1 = append(que1, s)
            gotMap[s] = 1
            if s == endWord {
                return 2
            }
        }
        if diffOne127(endWord, s) {
            enMap[s] = 2
        }
    }
    ans := 2
    for len(que1) > 0 {
        que2 := []string{}
        // fmt.Println(que1)
        for _, s := range que1 {
            if _, ok := map2[s]; ok {
                for _, s2 := range map2[s] {
                    if s2 == endWord {
                        return ans + 1
                    }
                    if _, ok2 := gotMap[s2]; !ok2 {
                        // if _, ok3 := enMap[s2]; ok3 {
                        //     return ans + 2
                        // }
                        que2 = append(que2, s2)
                        gotMap[s2] = 1
                    }
                }
            }
        }
        ans++
        // fmt.Println("    ", que2)
        que1 = que2
    }

    return 0
}

func diffOne127(s1, s2 string) bool {
    cnt := 0
    for i, ch := range s1 {
        if byte(ch) != s2[i] {          // ch is rune
            cnt++
            if cnt > 2 {
                break
            }
        }
    }
    return cnt == 1
}

func main_LT0127_20211110() {
// func main() {

    fmt.Println("ans:")

    st, en := "hit", "cog"
    arr := []string{"hot","dot","dog","lot","log","cog"}

    // st, en := "a", "c"
    // arr := []string{"a","b","c"}

    // st, en := "hit", "lot"
    // arr := []string{"hot","dot","dog","lot","log","cog"}
    

    ans := ladderLength(st, en, arr)

    fmt.Println(ans)

}
