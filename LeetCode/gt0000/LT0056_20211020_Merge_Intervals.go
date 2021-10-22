// package sdq
package main

import (
    "fmt"
    "sort"
)



// sort.Slice(intervals, func (i, j int) bool { return intervals[i][0] < intervals[j][0] })
// res = append(res, intervals[0])
// for i, j := 1, 0; i < len(intervals); i++ {
//     if intervals[i][0] <= res[j][1] {
//         if intervals[i][1] > res[j][1] {
//             res[j][1] = intervals[i][1]
//         }
//     } else {
//         res = append(res, intervals[i])
//         j++
//     }
// }


// for i := 1; i < len(intervals); i++ {
//     lastEnd = res[len(res)-1][1]
//     if intervals[i][0] <= lastEnd {
//         res[len(res)-1][1] = max(intervals[i][1], lastEnd)
//         continue
//     }
//     res = append(res, intervals[i])
// }


// for (int i = 1; i < ins.size(); i++) {
//     if (res.back().end < ins[i].start) res.push_back(ins[i]);
//     else
//         res.back().end = max(res.back().end, ins[i].end);
// }


// Runtime: 8 ms, faster than 92.53% of Go online submissions for Merge Intervals.
// Memory Usage: 4.9 MB, less than 31.92% of Go online submissions for Merge Intervals.
// 是否有序？
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }
        return intervals[i][1] > intervals[j][1]
    })

    // fmt.Println(intervals)

    ans := [][]int{}
    st, en := intervals[0][0], intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= en {
            if intervals[i][1] > en {
                en = intervals[i][1]
            }
        } else {
            ans = append(ans, []int{st, en})
            st, en = intervals[i][0], intervals[i][1]
        }
    }
    ans = append(ans, []int{st, en})
    return ans
}


func main_LT0056_20211020() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]int{{1,3},{2,6},{8,10},{15,18}}
    arr := [][]int{{1,4},{4,5}}

    ans := merge(arr)

    fmt.Println(ans)

}
