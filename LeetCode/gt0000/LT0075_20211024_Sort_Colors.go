// package sdq
package main

import (
    "fmt"
)



// nums[i] = 2
// if n <= 1 {
//     nums[one] = 1
//     one++
// }
// if n == 0 {
//     nums[zero] = 0
//     zero++
// }

// while white <= blue:
// if nums[white] == 0:
//     nums[red], nums[white] = nums[white], nums[red]
//     white += 1
//     red += 1
// elif nums[white] == 1:
//     white += 1
// else:
//     nums[white], nums[blue] = nums[blue], nums[white]
//     blue -= 1

// int second=n-1, zero=0;
// for (int i=0; i<=second; i++) {
//     while (A[i]==2 && i<second) swap(A[i], A[second--]);
//     while (A[i]==0 && i>zero) swap(A[i], A[zero++]);
// }


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Sort Colors.
func sortColors(nums []int)  {
    sz1 := len(nums)
    l, m, r := 0, 0, sz1 - 1
    for i := 0; i <= r; i++ {
        if nums[i] == 0 {
            nums[l] = 0
            l++
        } else if (nums[i] == 2) {
            nums[i], nums[r] = nums[r], nums[i]
            r--
            i--
        }
    }
    for m = l; m <= r; m++ {
        nums[m] = 1
    }

    // // for l < r {
    // for l <= m && m < r {
    //     if nums[l] == 0 {
    //         l++
    //         m++
    //     } else if nums[l] == 1 {
    //         nums[l], nums[m] = nums[m], nums[l]
    //         m++
    //     } else {
    //         nums[l], nums[r] = nums[r], nums[l]
    //         r--
    //     }
    //     fmt.Println(nums)
    // }
}

func main_LT0075_20211024() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{2,0,2,1,1,0}
    // arr := []int{2,0,1}
    // arr := []int{0}
    // arr := []int{1}
    // arr := []int{1,1,1}
    // arr := []int{0,0,0}
    // arr := []int{2,2,2}
    arr := []int{1,0,2}

    sortColors(arr)

    fmt.Println(arr)

}
