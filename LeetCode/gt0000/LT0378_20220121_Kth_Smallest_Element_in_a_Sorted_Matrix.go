// package sdq
package main

import (
    "fmt"
)



// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/discuss/85170/O(n)-from-paper.-Yes-O(rows).


// while (left < right) {     
//     int mid = left + (right - left) / 2;
//     int cnt = search_less_equal(matrix, mid);
//     if (cnt < k) left = mid + 1;
//     else right = mid;
// }
//
// while (i >= 0 && j < n) {
//     if (matrix[i][j] <= target) {
//         res += i + 1;
//         ++j;
//     } else {
//         --i;
//     }
// }


// func kthSmallest(matrix [][]int, k int) int {
// 	n := len(matrix)
// 	l, r := matrix[0][0], matrix[n-1][n-1]+1 // [l,r)
// 	for l < r {
// 		m := l + (r-l)>>1
// 		count := 0 // how many <= m
// 		j := n - 1
// 		for i := range matrix {
// 			for ; j >= 0 && matrix[i][j] > m; j-- {
// 			}
// 			count += j + 1 // [i][0] to [i][j] <= m
// 			if j < n-1 {
// 				j++
// 			}
// 		}
// 		if k <= count {
// 			r = m
// 		} else {
// 			l = m + 1
// 		}
// 	}
// 	return l
// }


// Runtime: 37 ms, faster than 63.70% of Go online submissions for Kth Smallest Element in a Sorted Matrix.
// Memory Usage: 7.3 MB, less than 52.59% of Go online submissions for Kth Smallest Element in a Sorted Matrix.
// bs
func kthSmallest(matrix [][]int, k int) int {
    sz1, sz2 := len(matrix), len(matrix[0])
    l, r := matrix[0][0], matrix[sz1 - 1][sz2 - 1]
    ans := 0
    for l <= r {
        mid := (l + r) / 2
        cnt := 0
        mxv := matrix[0][0]
        mxcnt := 1
        for i := 0; i < sz1; i++ {
            st, en := 0, sz2 - 1
            for st <= en {
                m := (st + en) / 2
                if matrix[i][m] == mid {
                    st, en = m + 1, m
                } else if matrix[i][m] > mid {
                    en = m - 1
                } else if matrix[i][m] < mid {
                    st = m + 1
                }
            }
            for st < sz2 && matrix[i][st] == mid {
                st++
            }
            if st > 0 {
                if matrix[i][st - 1] > mxv {
                    mxv = matrix[i][st - 1]
                    mxcnt = 0
                }
                for x := st; x > 0 && matrix[i][x - 1] == mxv; x-- {
                    mxcnt++
                }
            }
            // 6,6,6,6  for, not if.
            // if st < sz2 && matrix[i][st] == mid {
            //     st++
            // }
            // fmt.Print(st, ", ")
            // if st > 0 {
            //     if matrix[i][st - 1] > mxv {
            //         mxv = matrix[i][st - 1]
            //         mxcnt = 1
            //     } else if matrix[i][st - 1] == mxv {
            //         mxcnt++
            //     }
            // }
            cnt += st
        }
        // fmt.Println()
        // if cnt < k && cnt + mxcnt >= k {
            // fmt.Println(cnt, mxcnt, k, mxv)
        if cnt >= k && cnt - mxcnt < k {
            // fmt.Println(mxv)
            return mxv
        } else if cnt > k {
            ans = r
            r = mid - 1
        } else {
            ans = l
            l = mid + 1
        }
        // fmt.Println(l, r, mid, cnt, mxcnt)
    }
    // fmt.Println(" --- ")
    return ans
}


func main_LT0378_20220121() {
// func main() {

    fmt.Println("ans:")

    arr := [][]int{{1,5,9},{10,11,13},{12,13,15}}
    k := 8

    // arr := [][]int{{-5}}
    // k := 1

    // arr := [][]int{{1,2},{3,3}}
    // k := 4

    // [[1,2],[1,3]]
    // 1
    // arr := [][]int{{1,2},{1,3}}
    // k := 1


    // 18
    // arr := [][]int{{1,1,3,8,13},
    //                 {4,4,4,8,18},
    //                 {9,14,18,19,20},
    //                 {14,19,23,25,25},
    //                 {18,21,26,28,29}}
    // k := 13

    // 25
    // arr := [][]int{
    //             {2,6,6,7,10,14,18},
    //             {6,11,14,14,15,20,23},
    //             {11,11,17,21,25,30,30},
    //             {11,12,20,25,25,35,37},
    //             {16,16,20,29,34,35,39},
    //             {16,18,22,33,37,37,40},
    //             {17,23,26,36,38,41,42}}
    // k := 32


    // 6
    // arr := [][]int{{2,3,6,6,10},{5,9,12,17,19},{10,14,17,20,20},{15,17,20,24,24},{20,20,25,26,29}}
    // k := 4
    

    fmt.Println(kthSmallest(arr, k))

}
