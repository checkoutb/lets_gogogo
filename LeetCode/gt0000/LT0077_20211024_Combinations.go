// package sdq
package main

import (
    "fmt"
)


// public List<List<Integer>> combine(int n, int k) {
//     List<List<Integer>> result = new ArrayList<List<Integer>>();
//     if (k > n || k < 0) {
//         return result;
//     }
//     if (k == 0) {
//         result.add(new ArrayList<Integer>());
//         return result;
//     }
//     result = combine(n - 1, k - 1);
//     for (List<Integer> list : result) {
//         list.add(n);
//     }
//     result.addAll(combine(n - 1, k));
//     return result;
// }


// vector<vector<int>> result;
// int i = 0;
// vector<int> p(k, 0);
// while (i >= 0) {
//     p[i]++;
//     if (p[i] > n) --i;
//     else if (i == k - 1) result.push_back(p);
//     else {
//         ++i;
//         p[i] = p[i - 1];
//     }
// }
// return result;
// 。。 好像是遍历全部。



// 之前Java的代码。。。是 从1 遍历到 1<<n， 看1的位数是不是k，是k就转化为数组。。

// Runtime: 26 ms, faster than 22.43% of Go online submissions for Combinations.
// Memory Usage: 8.4 MB, less than 40.90% of Go online submissions for Combinations.
func combine(n int, k int) [][]int {
    ans := [][]int{}
    vst := make([]bool, n + 1)
    for i, _ := range vst {
        vst[i] = false
    }
    dfsa77([]int{}, vst, &ans, k, 1)
    return ans
}

func dfsa77(arr []int, vst []bool, ans *[][]int, k int, idx int) () {
    if len(arr) == k {
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    // mxi := len(vst) - k + idx           // .
    mxi := len(vst)
    for i := idx; i < mxi; i++ {
        if !vst[i] {
            arr = append(arr, i)
            vst[i] = true
            dfsa77(arr, vst, ans, k, i + 1)
            vst[i] = false
            // *arr = *arr[0 : len(*arr) - 1]           // cannot slice arr (type *[]int)
            arr = arr[0 : len(arr) - 1]
        }
    }
}


func main_LT0077_20211024() {
// func main() {

    fmt.Println("ans:")

    // n, k := 4, 2
    n, k := 1, 1


    ans := combine(n, k)

    fmt.Println(ans)

}
