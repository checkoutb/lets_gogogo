// package sdq
package main

import (
    "fmt"
)


// res.resize(1<<n);
// auto it = prev.rbegin();
// for (int i = 1<<(n-1); i < 1<<n; i++)
// {
//     res[i] = (1<<(n-1)) | *it;
//     it++;
// }


// for (int i = 0; i < (1 << n); i++) {
//     ret.add(i ^ (i >> 1));
// }




// Runtime: 8 ms, faster than 95.92% of Go online submissions for Gray Code.
// Memory Usage: 6.8 MB, less than 57.14% of Go online submissions for Gray Code.
// 0 1
// 00 01 11 10
// 000 001 011 010 110 111 101 100         100 101 111 110
// 复制, 前面加0 ,  前面加1且反转
// 1 << x | num     + 反序       不原先的0   0<<x 还是0 , 所以  只是复制. 并且反序复制.
func grayCode(n int) []int {
    ans := []int{0, 1}
    for n > 1 {
        n--
        t2 := 1 << n
        for i := len(ans) - 1; i >= 0; i-- {
            ans = append(ans, t2 | ans[i])              // 有点废 切片啊..
        }
    }
    return ans
}


func main_LT0089_20211027() {
// func main() {

    fmt.Println("ans:")


}
