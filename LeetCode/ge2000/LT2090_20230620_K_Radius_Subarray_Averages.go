// package sdq
package main

import (
    "fmt"
)


// Runtime220 ms
// Beats
// 98.73%
// Memory9.2 MB
// Beats
// 78.48%
func getAverages(nums []int, k int) []int {
    ans := make([]int, len(nums));
    for i := range ans {
        ans[i] = -1;
    }
    st := 0;
    sum2 := 0;
    kk := 2 * k + 1;
    for i, v := range nums {
        sum2 += v;
        if i - st + 1 == kk {
            ans[i - k] = sum2 / kk;
            sum2 -= nums[st];
            st++;
        }
    }
    return ans;
}

func main_LT2090_20230620() {
// func main() {

    arr := []int{7,4,3,9,1,8,5,2,6};
    k := 3;

    fmt.Println("ans:", getAverages(arr, k));


}
