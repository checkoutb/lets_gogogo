// package sdq
package main

import (
    "fmt"
)


// int l = 0, r = m * n - 1;
// while (l != r){
//     int mid = (l + r - 1) >> 1;
//     if (matrix[mid / m][mid % m] < target)
//         l = mid + 1;
//     else 
//         r = mid;
// }
// return matrix[r / m][r % m] == target;
// 视为一维。


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search a 2D Matrix.
// Memory Usage: 2.7 MB, less than 48.14% of Go online submissions for Search a 2D Matrix.
func searchMatrix(matrix [][]int, target int) bool {
    sz1, sz2 := len(matrix), len(matrix[0])
    l, r := 0, sz1 - 1
    for l <= r {
        mid := l + (r - l) / 2
        // fmt.Println(l, r, mid)
        // t1, t2, t3 := matrix[l][sz2 - 1], matrix[mid][sz2 - 1], matrix[r][sz2 - 1]
        t2 := matrix[mid][sz2 - 1]
        if t2 == target {
            return true
        } else if t2 > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    // fmt.Println(l)
    if l >= sz1 {
        return false
    }
    row := l
    l, r = 0, sz2 - 1
    for l <= r {
        mid := (l + r) / 2
        t2 := matrix[row][mid]
        if t2 == target {
            return true
        } else if t2 > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false
}

func main_LT0074_20211024() {
// func main() {

    fmt.Println("ans:")

    arr := [][]int{{1,3,5,7},{11,13,15,18},{21,25,31,32}}
    // t := 3
    // t := 12
    // t := 18
    // t := 21
    // t := 22
    // t := 111
    t := -123

    // arr := [][]int{{1}}
    // t := 2

    ans := searchMatrix(arr, t)

    fmt.Println(ans)
}
