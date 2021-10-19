// package sdq
package main

import (
    "fmt"
)



// for i := idx; i < len(candidates); i++ {
//     x := candidates[i]
//     combinationSumR(candidates, target-x, i, append(path, x), result)
// }
// dfs，每一步 减去 一个。 并且 后续的 idx 就是 >= i的


// def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
// dp = [[] for _ in range(target+1)]
// for c in candidates:                                  # O(N): for each candidate
//     for i in range(c, target+1):                      # O(M): for each possible value <= target
//         if i == c: dp[i].append([c])
//         for comb in dp[i-c]: dp[i].append(comb + [c]) # O(M) worst: for each combination
// return dp[-1]
// 竟然可以不回车。。py3的。 有点复杂。不看。


// Runtime: 12 ms, faster than 25.97% of Go online submissions for Combination Sum.
// Memory Usage: 6 MB, less than 33.53% of Go online submissions for Combination Sum.
func combinationSum(candidates []int, target int) [][]int {
    ans := &[][]int{}
    dfsa39(candidates, target, ans, 0, []int{})
    return *ans
}

func dfsa39(cs []int, need int, ans *[][]int, idx int, arr []int) {
    if need == 0 {
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)            // == ?  same object?  不能直接 append(arr) ,, 需要 copy一个。
        return
    }
    if idx >= len(cs) || need < 0 {
        return
    }
    sz1 := len(arr)
    dfsa39(cs, need, ans, idx + 1, arr)
    for need > 0 {
        need -= cs[idx]
        arr = append(arr, cs[idx])
        dfsa39(cs, need, ans, idx + 1, arr)
    }
    arr = arr[0 : sz1]
}

func main_LT0039_20211018() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{2,3,6,7}
    // tar := 7

    // arr := []int{2,3,5}
    // tar := 8

    // arr := []int{}
    // tar := 1

    arr := []int{2,7,6,3,5,1}
    tar := 9
        

    ans := combinationSum(arr, tar)

    fmt.Println(len(ans))
    for _, ar := range ans {
        for _, val := range ar {
            fmt.Print(val, ", ")
        }
        fmt.Println()
    }

}
