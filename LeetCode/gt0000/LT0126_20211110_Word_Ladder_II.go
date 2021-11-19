// package sdq
package main

import (
    "fmt"
    // "sort"
)


// 都很长。
// 很多(基本都是)都使用了 字符替换 来 获得 diff 1
// 很多都是 bfs + dfs， 也有 一个bfs 搞定的。




// sth changed
// Runtime: 4 ms, faster than 96.55% of Go online submissions for Word Ladder II.
// Memory Usage: 3.7 MB, less than 58.62% of Go online submissions for Word Ladder II.

// 字符替换试下
// bu ,好像也不是， 应该是计算  到达这个 string的 最短长度。 dp？ floyd？ 点对间最短路径。 不， 不求全部的。只需要 2个点，类似最小生成树。
// 没有注意到 前一题 有没有这种解法。。
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    ans := [][]string{}
    // vst := map[string]bool{}
    b2 := false
    b3 := false
    for _, s := range wordList {
        if s == beginWord {
            b2 = true
            // break
        }
        if s == endWord {
            b3 = true
        }
    }
    if !b3 {
        return ans
    }
    if !b2 {
        wordList = append(wordList, beginWord)
    }
    sz1 := len(wordList)

    arr := make([]int, sz1)         // 0: not visited,  >0 : min_path
    arr3 := make([][]int, sz1)      // prefixes
    map2 := map[string][]int{}      // diff 1
    map3 := map[string]int{}
    am := make([]map[int]bool, sz1)
    for i, _ := range am {
        am[i] = map[int]bool{}
    }
    // map2 :=

    for i, s := range wordList {
        for j := i + 1; j < sz1; j++ {
            if diff1126(s, wordList[j]) {
                map2[s] = append(map2[s], j)
                map2[wordList[j]] = append(map2[wordList[j]], i)
            }
        }
        map3[s] = i
    }

    que1 := []string{beginWord}
    mnPath := 1
    arr[map3[beginWord]] = mnPath
    got := false
    for len(que1) > 0 {
        que2 := []string{}
        mnPath++

        for _, s := range que1 {
            idx := map3[s]
            for _, i2 := range map2[s] {    // idx of diff1
                // idx2 := map3[wordList[i2]]
                idx2 := i2
                s2 := wordList[i2]
                if arr[idx2] == 0 || arr[idx2] == mnPath {     // not visited  or  same-length
                    que2 = append(que2, s2)
                    arr[idx2] = mnPath
                    if _, ok := am[idx2][idx]; !ok {
                        arr3[idx2] = append(arr3[idx2], idx)
                        am[idx2][idx] = true
                    }

                    if s2 == endWord {
                        got = true
                        // goto AAA;
                    }
                }
            }
        }
        if got {
            goto AAA
        }

        que1 = que2
    }
    AAA:
    fmt.Println(arr3)
    fmt.Println(arr)
    if arr[map3[endWord]] != 0 {
        // for i := 0; i < len(arr3); i++ {
        //     arr3[i] = sort.
        // }
        dfsa126(&ans, arr, arr3, map3[endWord], []string{}, wordList)
    }

    return ans
}

func dfsa126(ans *[][]string, arr []int, arr3 [][]int, idx int, tmp []string, words []string) {
    if arr[idx] == 1 {  // beginWord
        tmp2 := make([]string, len(tmp))
        copy(tmp2, tmp)
        tmp2 = append(tmp2, words[idx])
        l, r := 0, len(tmp2) - 1
        for l < r {
            tmp2[l], tmp2[r] = tmp2[r], tmp2[l]
            l++
            r--
        }
        *ans = append(*ans, tmp2)
        return
    }
    tmp = append(tmp, words[idx])
    for _, i := range arr3[idx] {     // for each prefix
        dfsa126(ans, arr, arr3, i, tmp, words)
    }
    tmp = tmp[0 : len(tmp) - 1]
}

func diff1126(s1, s2 string) bool {
    cnt := 0
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            cnt++
            if cnt > 1 {
                return false
            }
        }
    }
    return cnt == 1
}

func main_LT0126_20211110() {
// func main() {

    fmt.Println("ans:")

    // st, en := "hit", "cog"
    // arr := []string{"hot","dot","dog","lot","log","cqg"}

    st, en := "red", "tax"
    arr := []string{"ted","tex","red","tax","tad","den","rex","pee"}        // prefix 重复了。 1-3-0-4  1-2-0-4   4就会有2个0的前缀。
    

    ans := findLadders(st, en, arr)

    fmt.Println(ans)

}
