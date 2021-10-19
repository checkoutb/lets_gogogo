// package sdq
package main

import (
    "fmt"
)


// LinkedList<List<Integer>> result = new LinkedList<List<Integer>>();
// result.add(new ArrayList<Integer>());
// for (int n: nums) {
//     int size = result.size();
//     while(size > 0) {
//         List<Integer> current = result.pollFirst();
//         for (int i = 0; i <= current.size(); i++) {
//             List<Integer> temp = new ArrayList<Integer>(current);
//             temp.add(i, n);
//             result.add(temp);
//         }
//         size--;
//     }
// }
// queue， 第一次，加 第一个元素，  第二次加 第二个元素， 第三次 家第三个元素。  加的时候 是在 所有地方加。


// ans := [][]int{}
// for i := range nums {
//     q := make([]int, 0, len(nums) - 1)
//     q = append(q, nums[:i]...)
//     q = append(q, nums[i+1:]...)
//     for _, p := range permute(q) {               // dfs
//         ap := make([]int, len(p), len(p)+1)
//         copy(ap, p)
//         ap = append(ap, nums[i])
//         ans = append(ans, ap)
//     }
// }


// for i in xrange(len(nums)):
//     self.dfs(nums[:i]+nums[i+1:], path+[nums[i]], res)





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
// Memory Usage: 2.7 MB, less than 69.28% of Go online submissions for Permutations.
// num -10 - 10 ... + 11   1-21  int bit
func permute(nums []int) [][]int {
    ans := [][]int{}
    dfsa46(nums, 0, []int{}, &ans)
    return ans
}

func dfsa46(nums []int, vst int, arr []int, ans *[][]int) {
    if len(arr) == len(nums) {              // or vst == 0,  initialize vst with nums
        arr2 := make([]int, len(arr))
        copy(arr2, arr)
        *ans = append(*ans, arr2)
        return
    }
    for _, val := range nums {
        t2 := val + 11
        if vst & (1 << t2) == 0 {
            arr = append(arr, val)
            dfsa46(nums, vst | (1 << t2), arr , ans)
            arr = arr[0 : len(arr) - 1]
        }
    }
}


func main_LT0046_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{1,2,3}
    // arr := []int{-1,-2}
    arr := []int{-1}

    ans := permute(arr)

    for _, ar := range ans {
        for _, e := range ar {
            fmt.Print(e, ", ")
        }
        fmt.Println()
    }

}
