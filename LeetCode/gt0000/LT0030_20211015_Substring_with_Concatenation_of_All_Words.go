// package sdq
package main

import (
    "fmt"
)



// n := len(words[0]) * len(words)
// .....  words of the same length .....................
// 



// for i := 0; i < len(words[0]); i++ {
//     pos := i
//     for pos+n <= len(s) {
//         exist := map[string]int{}
//         times := len(words)
//         for ; times > 0; times-- {
    // str := s[pos+len(words[0])*(times-1) : pos+len(words[0])*times]
    // exist[str]++
    // if exist[str] > wordsMap[str] {
    //     break
    // }
    // if times == 0 {
    //     ans = append(ans, pos)
    //     pos += len(words[0])
    // } else {
    //     pos += len(words[0]) * times
    // }
// word长度相同，所以以 [0, word.size) 做一个划分。   0 的时候 会 把所有  sz * n  为开始的 substr 全部计算的，所以 不需要 sz
// 如果 substr 多于 word， 就跳出。  times是--。





// Runtime: 128 ms, faster than 50.00% of Go online submissions for Substring with Concatenation of All Words.
// Memory Usage: 10.7 MB, less than 5.66% of Go online submissions for Substring with Concatenation of All Words.
// ["aaa","aaaaa"]
// trie?  ["abc","abc","abc"]
func findSubstring(s string, words []string) []int {
    ans := []int{}
    dummy := generateNode(0, false)
    map2 := map[string]int{}
    // map5 := map[string]int{}            // just a converter    words duplicate
    //                                 // .. 不如直接记录到 Node里。。。。。。
    lensum := 0
    for _, w := range words {
        map2[w]++
        lensum += len(w)
        // map5[w] = idx
        generateTrie(dummy, w, 0)
    }

    map3 := map[int][]int{}     // start : end[]
    map11 := map[int][]*Node{}      // start : Node[]
    map12 := map[int][]int{}            // end : start[]

    for i := 0; i < len(s); i++ {
        arr2 := &[]int{}
        arr3 := &[]*Node{}
        findByTrie(dummy, s, i, arr2, arr3)
        map3[i] = *arr2         // s[i : arr2[x]] is a word: arr3[x].word
        map11[i] = *arr3
        for _, v := range *arr2 {
            map12[v] = append(map12[v], i)
        }
    }

    // 那么就是求连续的数字串。至少得 全部word相加的长度。 。 可以加速的， 只要 subString[x, x+sum(len(word)))]  不是一个完整的划分 就肯定不行。
    // 是完整的话费，就看里面是否 全部出现一次 。。。。 bitmap,,dan...
    // 不过 我这里 把长度记录在 头上了。。。 还得再转一次。。。。。
    // 记录的是 结束的下标，不是 长度。。
    // 而且没有 word的 信息。。

    // map4 := map[int][]int{}

    // for _, v := range map3 {
    //     for t6 := range v {
    //         // map4[k + t6]

    //     }
    // }

    // // 滚动，而不是重新计算。  可能性太多了。。。。
    // lst := 0
    // for i := 0; i + lensum <= len(s); i++ {
    //     _, ok := map3[i]
    //     if ok {
    //         for lst < i + lensum - 1 {

    //         }
    //     }

    //     // _, ok := map3[i]
    //     // if ok {
    //     //     _, ok2 := map12[i + lensum - 1]
    //     //     if ok2 {

    //     //     }
    //     // }

    // }

    // fmt.Println(", ", lensum)
    // fmt.Println(map3)
    // fmt.Println(dummy)
    // fmt.Println(map11)
    // fmt.Println("=========")
    // showNode(dummy)

    for i := 0; i + lensum <= len(s); i++ {
        _, ok := map3[i]
        if ok {
            _, ok2 := map12[i + lensum - 1]
            if (ok2) {
                if dfsa6(s, i, map3, map11, map2, i + lensum - 1) {
                    ans = append(ans, i)
                }
            }
        }
    }

    return ans
}

// 全都需要预先加载。 不能直接计算。。
// map3 := map[int][]int{}     // start : end[]
// map11 := map[int][]*Node{}      // start : Node[]
// en]
func dfsa6(s string, idx int, map3 map[int][]int, map11 map[int][]*Node, map2 map[string]int, en int) bool {
    if idx > en {
        if idx == en + 1 {
            return true
        }
        return false
    }
    _, ok := map3[idx]
    if ok {
        for i, en2 := range map3[idx] {
            s2 := map11[idx][i].word
            if map2[s2] > 0 {
                map2[s2]--
                if dfsa6(s, en2 + 1, map3, map11, map2, en) {
                    map2[s2]++
                    return true;
                }
                map2[s2]++
            }
        }
    }

    return false
}

type Node struct {
    ch byte
    isEnd bool
    childs map[byte]*Node
    word string
}

// (parent(previous-char)'s Node, word, word's index)
func generateTrie(node *Node, word string, idx int) {
    var t2 byte = (word[idx] - 'a')
    if (idx == len(word) - 1) {
        _, ok := node.childs[t2]
        if ok {
            node.childs[t2].isEnd = true
        } else {
            node.childs[t2] = generateNode(t2, true)
        }
        node.childs[t2].word = word
        return;
    }
    _, ok := node.childs[t2]
    if !ok {
        node.childs[t2] = generateNode(t2, false)
    }
    t3 := node.childs[t2]
    generateTrie(t3, word, idx + 1)
}

func generateNode(val byte, isEnd bool) *Node {
    return &Node{val, isEnd, make(map[byte]*Node), ""}
}

// parent's node
func findByTrie(node *Node, s string, idx int, ans *[]int, arr3 *[]*Node) {
    if (idx == len(s) || node == nil) {
        return
    }
    t2 := s[idx] - 'a'          // ... - 'a'
    t3 := node.childs[t2]
    
    // fmt.Println("       ", idx)
    // fmt.Println(node.childs)
    // fmt.Println(t2)

    if t3 == nil {
        return
    }
    if t3.isEnd {
        *ans = append(*ans, idx)
        *arr3 = append(*arr3, t3)
    }
    findByTrie(t3, s, idx + 1, ans, arr3)
}

// func dfsa5(s string, idx int, )


func main_LT0030_20211015() {
// func main() {

    fmt.Println("ans:")

    // fmt.Println("asd"[1] - 'a')

    // s := "barfoothefoobarman"
    // ws := []string{"foo","bar"}

    // s := "wordgoodgoodgoodbestword"
    // ws := []string{"word","good","best","word"}

    s := "barfoofoobarthefoobarman"
    ws := []string{"bar","foo","the"}

    ans := findSubstring(s, ws)

    fmt.Println("sz: ", len(ans))
    for _, val := range ans {
        fmt.Println(val)
    }
}

func showNode(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.ch, ", ", node.isEnd, ", ", node.word)
    for _, n := range node.childs {
        showNode(n)
    }
}
