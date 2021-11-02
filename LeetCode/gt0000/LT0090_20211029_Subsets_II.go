// package sdq
package main

import (
    "fmt"
    "sort"
)


// func backtrack(nums []int,pos int,list []int,res *[][]int){
//     ans:=make([]int,len(list))
//     copy(ans,list)
//     *res=append(*res,ans)
//     for i:=pos;i<len(nums);i++{
//         if i!=pos&&nums[i]==nums[i-1]{
//             continue
//         }
//         list=append(list,nums[i])
//         backtrack(nums,i+1,list,res)
//         list=list[:len(list)-1]
//     }
// }


// if _, ok := visited[str]; !ok {


// set<vector<int>> set;            // ...... go.map.key.type == Slice ?     no, invalid map key



// Runtime: 4 ms, faster than 25.50% of Go online submissions for Subsets II.
// Memory Usage: 2.6 MB, less than 68.46% of Go online submissions for Subsets II.
// 做一个 arr.sort.join(',') 的 key.    这种是 一层 一个元素
// 或者  一层 一个 数字的所有可能.
func subsetsWithDup(nums []int) [][]int {
    ans := [][]int{}
    sort.Ints(nums)
    dfsa90(nums, 0, &ans, []int{})
    return ans
}

func dfsa90(nums []int, idx int, ans *[][]int, arr []int) {
    fmt.Println(idx)
    if idx == len(nums) {
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    mxi := idx
    for (mxi + 1) < len(nums) && nums[mxi + 1] == nums[mxi] {               // mxi < len || (mxi > 0 && [mxi] == [mxi - 1])
        mxi++
    }
    mxi++
    // sz1 := len(arr)
    for t2 := nums[idx]; idx <= mxi; idx++ {
        dfsa90(nums, mxi, ans, arr)
        arr = append(arr, t2)
    }
    // arr = arr[0 : sz1]
}


func main_LT0090_20211029() {
// func main() {

    // map2 := map[[]int]int{}

    fmt.Println("ans:")

    // arr := []int{1,2,2}
    arr := []int{0}

    ans := subsetsWithDup(arr)

    fmt.Println(ans)

}
