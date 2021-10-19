// package sdq
package main

import (
    "fmt"
)



// for(int i = 0; i < nums.length; i++)
// {
//     if(visited[i] || i > 0 && nums[i] == nums[i - 1] && !visited[i - 1]) continue;
//     per.add(nums[i]);
//     visited[i] = true;
//     dfs(nums, visited, uniquePermutations, per);
//     per.remove(per.size() - 1);
//     visited[i] = false;
// }   
// dfs, 先 sort nums
// 访问过， 或 和前面的相等且前面的没有访问过(且本元素没有访问过)  就 跳过。
// 就是 相同值，需要 一个个 取。  这样的话， 本次for循环 就不会 取到 重复的值。


// Runtime: 16 ms, faster than 17.36% of Go online submissions for Permutations II.
// Memory Usage: 4.1 MB, less than 54.96% of Go online submissions for Permutations II.
// 1,1,2
// 总不能把 array 转为 string 做key吧。。
// 或者 dfs  只能插入到 前面不能自己的地方， 不，这个是 bfs
// dfs的话， 要一次性把 所有 这个值 生成，并且 还需要 判断不能等于之前的 最后一个。
// 直接用 下标作为 位移 就可以了。 。。 还是 map吧。
func permuteUnique(nums []int) [][]int {
    map2 := map[int]int{}
    for _, v := range nums {
        map2[v]++
    }
    ans := [][]int{}
    dfsa47(len(nums), []int{}, &ans, map2)
    return ans
}

func dfsa47(sz1 int, arr []int, ans *[][]int, map2 map[int]int) {
    if len(arr) == sz1 {
        arr2 := make([]int, sz1)
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    lst := -1000
    if len(arr) > 0 {
        lst = arr[len(arr) - 1]
    }
    for k, v := range map2 {
        if k == lst {
            continue
        }
        v2 := v
        sz2 := len(arr)
        for v > 0 {     // map2[k] > 0
            v--
            arr = append(arr, k)
            map2[k]--
            dfsa47(sz1, arr, ans, map2)
        }
        map2[k] = v2
        arr = arr[0 : sz2]
    }
}


func main_LT0047_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{1,1,2}
    arr := []int{1,2,3}

    ans := permuteUnique(arr)

    for _, ar := range ans {
        for _, e := range ar {
            fmt.Print(e, ", ")
        }
        fmt.Println()
    }

}
