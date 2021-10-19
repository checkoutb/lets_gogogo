// package sdq
package main

import (
    "fmt"
)


// for(int i = order;i<num.size();i++) // iterative component
// {
//     if(num[i]>target) return;
//     if(i&&num[i]==num[i-1]&&i>order) continue; // check duplicate combination
//     local.push_back(num[i]),
//     findCombination(res,i+1,target-num[i],local,num); // recursive componenet
//     local.pop_back();
// }
// 这种好像是， 第一次的时候 会 dfs， 后续相同值 不会 dfs， 但是 第一次的dfs 里面可以继续获得相同值。
// 这样，保证， 第一次的dfs 必然只有 1个 值，  dfs的内部 是2个 相同值， dfs再dfs的内部 是3个相同值。






// Runtime: 8 ms, faster than 28.11% of Go online submissions for Combination Sum II.
// Memory Usage: 5.1 MB, less than 19.46% of Go online submissions for Combination Sum II.
func combinationSum2(candidates []int, target int) [][]int {
    ans := &[][]int{}
    
    map2 := map[int]int{}
    for _, v := range candidates {
        map2[v]++
    }
    cs := []int{}
    for k, _ := range map2 {
        cs = append(cs, k)
    }
    dfsa40a(cs, 0, target, ans, []int{}, map2)
    return *ans
}

func dfsa40a(cs []int, idx int, need int, ans *[][]int, arr []int, map2 map[int]int) {
    if need == 0 {
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    if idx >= len(cs) || need < 0 {             // need < 0
        return
    }
    dfsa40a(cs, idx + 1, need, ans, arr, map2);
    sz1 := len(arr);
    for i := 0; i < map2[cs[idx]]; i++ {
        need -= cs[idx]
        arr = append(arr, cs[idx])
        dfsa40a(cs, idx + 1, need, ans, arr, map2)
    }

    arr = arr[0 : sz1]
}


// 没有去重
func combinationSum2a(candidates []int, target int) [][]int {
    ans := &[][]int{}
    dfsa40(candidates, 0, target, ans, []int{})
    return *ans
}

func dfsa40(cs []int, idx int, need int, ans *[][]int, arr []int) {
    if need == 0 {
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    if idx >= len(cs) {
        return
    }
    dfsa40(cs, idx + 1, need, ans, arr)     // not choose
    arr = append(arr, cs[idx])
    dfsa40(cs, idx + 1, need - cs[idx], ans, arr)   // choose
    arr = arr[0 : len(arr) - 1]
}

func main_LT0040_20211018() {
// func main() {

    fmt.Println("ans:")


}
