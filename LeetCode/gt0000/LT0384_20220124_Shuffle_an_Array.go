// package sdq
package main

import (
    "fmt"
    "math/rand"
)


// for(int i=0;i<n-1;i++){
//     int index = rand()%(n-i)+i;
//     swap(arr[i],arr[index]);
// }

// for (int i = 0;i < result.size();i++) {
//     int pos = rand()%(result.size()-i);
//     swap(result[i+pos], result[i]);
// }



// Runtime: 47 ms, faster than 78.08% of Go online submissions for Shuffle an Array.
// Memory Usage: 10 MB, less than 26.03% of Go online submissions for Shuffle an Array.
type Solution struct {
    arr []int
    // arr2 []int
}

func Constructor(nums []int) Solution {
    return Solution{ nums }
}

func (this *Solution) Reset() []int {
    return this.arr
}

func (this *Solution) Shuffle() []int {
    ans := make([]int, len(this.arr))
    copy(ans, this.arr)
    for i, _ := range this.arr {
        i2 := rand.Intn(len(this.arr))
        ans[i], ans[i2] = ans[i2], ans[i]
    }
    return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main_LT0384_20220124() {
// func main() {

    fmt.Println("ans:")


}
