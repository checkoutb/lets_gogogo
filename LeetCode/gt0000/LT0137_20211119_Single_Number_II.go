// package sdq
package main

import (
    "fmt"
)




// a, b := 0, 0
// for _, n := range nums {
//     a, b = ^b & (a ^ n), n & ^b & a | ^n & b & ^a
// }

// sort

// map


// for( int i = 0 ; i < 32 ; ++i){          // 32
//     int count = 0;
//     for(auto ele: nums){
//         if(ele & shift)
//             count+=1;
//     }
//     if( count % 3 != 0){
//         ans += shift ;
//     }
//     shift *= 2 ;
// }


// int ones = 0, twos = 0;
// for(int i = 0; i < A.length; i++){
//     ones = (ones ^ A[i]) & ~twos;
//     twos = (twos ^ A[i]) & ~ones;
// }
// return ones;




// Runtime: 4 ms, faster than 95.88% of Go online submissions for Single Number II.
// Memory Usage: 3.6 MB, less than 49.48% of Go online submissions for Single Number II.

// 线性  固定空间
// 线性， hash越等于线性。
// 固定空间， 至少 1/3
// 想起来了， bit。
// bit[x].count % 3 ..... !!!
// 还记得 某道题目的discuss 里 有 通解 。 重复n次，m个只出现一次。。。。我记得空间是和 m有关的。 
func singleNumber(nums []int) int {
    arr := [32]int{}
    for _, v := range nums {
        for j := 0; j < 32; j++ {
            if v & (1 << j) != 0 {
                arr[j]++
            }
        }
    }
    var ans int32 = 0
    fmt.Println(arr)
    for i, v := range arr {
        if v % 3 == 1 {
            ans |= (1 << i)
        }
    }
    return (int) (ans)
}


func main_LT0137_20211119() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{2,2,3,2}
    arr := []int{-2,-2,1,1,4,1,4,4,-4,-2}

    ans := singleNumber(arr)

    fmt.Println(ans)

}
