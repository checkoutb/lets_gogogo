// package sdq
package main

import (
    "fmt"
)



// ["aaa","aaaaa"]
// trie?  ["abc","abc","abc"]
func findSubstring(s string, words []string) []int {
    ans := []int{}
    dummy := generateNode(0, false)
    map2 := map[string]int{}
    // map5 := map[string]int{}            // just a converter    words duplicate
    //                                 // .. 不如直接记录到 Node里。。。。。。
    lensum := 0
    for idx, w := range words {
        map2[w]++
        lensum += len(w)
        // map5[w] = idx
        generateTrie(dummy, w, 0)
    }

    map3 := map[int][]int{}
    map11 := map[int][]*Node{}

    for i := 0; i < len(s); i++ {
        arr2 := &[]int{}
        arr3 := &[]*Node{}
        findByTrie(dummy, s, i, arr2, arr3)
        map3[i] = *arr2
        map11[i] = *arr3
    }

    // 那么就是求连续的数字串。至少得 全部word相加的长度。 。 可以加速的， 只要 subString[x, x+sum(len(word)))]  不是一个完整的划分 就肯定不行。
    // 是完整的话费，就看里面是否 全部出现一次 。。。。 bitmap,,dan...
    // 不过 我这里 把长度记录在 头上了。。。 还得再转一次。。。。。
    // 记录的是 结束的下标，不是 长度。。
    // 而且没有 word的 信息。。

    map4 := map[int][]int{}

    for _, v := range map3 {
        for t6 := range v {
            // map4[k + t6]

        }
    }




    

    return ans
}

// 全都需要预先加载。 不能直接计算。。
func dfsa6() {

}

type Node struct {
    ch byte
    isEnd bool
    childs map[byte]*Node
    words []string
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
        node.childs[t2].words = append(node.childs[t2].words, word)
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
    return &Node{val, isEnd, make(map[byte]*Node), []string{}}
}

// parent's node
func findByTrie(node *Node, s string, idx int, ans *[]int, arr3 *[]*Node) {
    if (idx == len(s) || node == nil) {
        return
    }
    t2 := s[idx]
    t3 := node.childs[t2]
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


// func main_LT0030_20211015() {
func main() {

    fmt.Println("ans:")

    fmt.Println("asd"[1] - 'a')

}
