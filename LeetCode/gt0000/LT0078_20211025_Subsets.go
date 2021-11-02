// package sdq
package main

import (
    "fmt"
)


// int n = 1 << nums.size();
// for (int i = 0; i < n; ++i) {
//     vector<int> sub;
//     for (int j = 0; j < nums.size(); ++j) {
//         if ((1 << j) & i) sub.push_back(nums[j]);
//     }
//     subs.push_back(move(sub));
// }


// for (int i = 0; i < nums.size(); i++)
// {
//     vector<vector<int>> temp=res;
//     for (int j =0; j< temp.size(); j++)
//     {
//         temp[j].push_back(nums[i]);
//         res.push_back(temp[j]);
//     }
// }
// 对现有的所有 增加一个 nums[i], 然后再放到 res中, 等于就是 res中 元素个数*2, 新增的元素 都带有 nums[i]


// res = append(res, []int{})
// if nums == nil {
//     return res
// }
// for k, v := range nums {
//     nr := subsets(nums[k+1:])
//     for _, s := range nr {
//         ns := []int{v}
//         ns = append(ns, s...)
//         res = append(res, ns)
//     }  
// }




// Runtime: 2 ms, faster than 28.31% of Go online submissions for Subsets.
// Memory Usage: 2.6 MB, less than 21.16% of Go online submissions for Subsets.
func subsets(nums []int) [][]int {
    ans := [][]int{}
    dfsa78(nums, 0, &ans, []int{})
    return ans
}

func dfsa78(nums []int, idx int, ans *[][]int, arr []int) {
    if idx == len(nums) {
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    dfsa78(nums, idx + 1, ans, arr)
    arr = append(arr, nums[idx])
    dfsa78(nums, idx + 1, ans, arr)
    arr = arr[ : len(arr) - 1]
}

func main_LT0078_20211025() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,2,3}

    ans := subsets(arr)

    fmt.Println(ans)

}
