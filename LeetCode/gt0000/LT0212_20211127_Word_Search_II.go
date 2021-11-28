// package sdq
package main

import (
    "fmt"
)



// words -> Trie





// Runtime: 1433 ms, faster than 25.07% of Go online submissions for Word Search II.
// Memory Usage: 2.6 MB, less than 76.58% of Go online submissions for Word Search II.
// Runtime: 2308 ms, faster than 6.33% of Go online submissions for Word Search II.
// Memory Usage: 2.6 MB, less than 76.58% of Go online submissions for Word Search II.
func findWords(board [][]byte, words []string) []string {
    map2 := map[byte][]string{}
    for i := range words {
        map2[words[i][0]] = append(map2[words[i][0]], words[i])
    }
    ans := []string{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if _, ok := map2[board[i][j]]; ok {
                for m := 0; m < len(map2[board[i][j]]); m++ {
                    if dfsa212(board, i, j, map2[board[i][j]][m], 0) {
                        ans = append(ans, map2[board[i][j]][m])
                        // map2[board[i][j]] = map2[board[i][j]][0 : m] + map2[board[i][j]][m + 1 : ]
                        map2[board[i][j]] = append(map2[board[i][j]][0 : m], map2[board[i][j]][m + 1 : ]...)
                        m--
                    }
                }
            }
        }
    }
    return ans
}

func dfsa212(board [][]byte, i, j int, word string, idx int) bool {
    if idx == len(word) {
        return true
    }
    if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] != word[idx] {
        return false
    }
    ori := board[i][j]
    board[i][j] = '@'
    if (dfsa212(board, i + 1, j, word, idx + 1) || 
        dfsa212(board, i - 1, j, word, idx + 1) || 
        dfsa212(board, i, j - 1, word, idx + 1) || 
        dfsa212(board, i, j + 1, word, idx + 1)) {
        board[i][j] = ori 
        return true
    }
    board[i][j] = ori 
    return false
}


func main_LT0212_20211127() {
// func main() {

    fmt.Println("ans:")

    arr := [][]byte{{'a','b','c','e'},{'x','x','c','d'},{'x','x','b','a'}}
    arr2 := []string{"abc","abcd"}

    ans := findWords(arr, arr2)

    fmt.Println(ans)

}
