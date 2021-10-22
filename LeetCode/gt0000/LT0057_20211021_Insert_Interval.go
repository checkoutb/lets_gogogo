// package sdq
package main

import (
    "fmt"
)




// left := make([][]int, 0)
// right := make([][]int, 0)
// start, end := newInterval[0], newInterval[1]
// for i := range intervals{
//     if intervals[i][1] < start{
//         left = append(left, intervals[i])
//     } else if end < intervals[i][0]{
//         right = append(right, intervals[i])
//     } else {
//         start = min(start, intervals[i][0])
//         end = max(end, intervals[i][1])
//     }
// }
// return append(append(left, []int{start, end}), right...)
// 。。。。。。。。。。。。。。。。。。。。。。。



// 
// s, e = newInterval.start, newInterval.end
// left = [i for i in intervals if i.end < s]
// right = [i for i in intervals if i.start > e]
// if left + right != intervals:
//     s = min(s, intervals[len(left)].start)
//     e = max(e, intervals[~len(right)].end)
// return left + [Interval(s, e)] + right


// s, e = newInterval.start, newInterval.end
// parts = merge, left, right = [], [], []
// for i in intervals:
//     parts[(i.end < s) - (i.start > e)].append(i)
// if merge:
//     s = min(s, merge[0].start)
//     e = max(e, merge[-1].end)
// return left + [Interval(s, e)] + right


// s, e = newInterval.start, newInterval.end
// left, right = [], []
// for i in intervals:
//     if i.end < s:
//         left += i,
//     elif i.start > e:
//         right += i,
//     else:
//         s = min(s, i.start)
//         e = max(e, i.end)
// return left + [Interval(s, e)] + right




// 第四次18ms， 第5次 8ms 。。 是不是有缓存之类的。。 反正就是不稳定啊。
// 最下面和 上面2次的 区别是  最下面有 fmt.Println(st, ", ", en)
// Runtime: 9 ms, faster than 47.19% of Go online submissions for Insert Interval.
// Memory Usage: 4.8 MB, less than 61.56% of Go online submissions for Insert Interval.
// Runtime: 8 ms, faster than 87.81% of Go online submissions for Insert Interval.
// Memory Usage: 4.6 MB, less than 99.69% of Go online submissions for Insert Interval.
// Runtime: 23 ms, faster than 9.06% of Go online submissions for Insert Interval.
// Memory Usage: 4.8 MB, less than 61.56% of Go online submissions for Insert Interval.
// [0:xx] (+ newInterval) + [xx:]
func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }
    if newInterval[0] >= intervals[len(intervals) - 1][0] {     // kaishi  >= kaishi
        if newInterval[0] > intervals[len(intervals) - 1][1] {  // new st > end
            intervals = append(intervals, newInterval)
        } else {
            if newInterval[1] > intervals[len(intervals) - 1][1] {
                intervals[len(intervals) - 1][1] = newInterval[1]
            }
        }
        return intervals
    }
    if newInterval[1] < intervals[0][0] {
        return append([][]int{newInterval}, intervals...)
    }
    st, en := -1, -1
    for i := 0; i < len(intervals); i++ {
        if intervals[i][1] >= newInterval[0] {
            if st == -1 {
                st = i
            }
            if intervals[i][0] > newInterval[1] {
                break
            }
            en = i
        }
    }
    // fmt.Println(st, ", ", en)
    if en < st {
        t3 := make([][]int, len(intervals) - st)
        copy(t3, intervals[st : len(intervals)])
        intervals = append(append(intervals[0 : st], newInterval), t3...)           // 先计算所有 slice，然后append，还是一边slice一边append？ 前者吧。 houzhe...
    } else {
        t2en := newInterval[1]
        if intervals[en][1] > t2en {            // [] faster then [][], but...
            t2en = intervals[en][1]
        }
        t2st := intervals[st][0]
        if newInterval[0] < t2st {
            t2st = newInterval[0]
        }
        intervals = append(intervals[0 : st], intervals[en : len(intervals)]...)
        intervals[st][0], intervals[st][1] = t2st, t2en
    }
    return intervals
}


func main_LT0057_20211021() {
// func main() {

    // t2 := []int{1,2,3,5,6}
    // fmt.Println(t2[3 : len(t2)])
    // t2 = append(append(t2[0 : 3], 4), t2[3 - 1 : len(t2)]...)           // ...append 和 [] 混合进行的。。 而且 append是 映射到 底层的数组的， 所以 后面的 t2[] 就取的不对的。

    // t3 := t2[3 : len(t2)]
    // fmt.Println(t3)
    // t2 = append(append(t2[0 : 3], 4), t3...)
    // fmt.Println(t3)         // 2次 t3 的值 不同。 。 是映射到 同一个数组，  天然 指针。
    // fmt.Println(t2)

    // t4 := make([]int, len(t2) - 3)
    // copy(t4, t2[3 : len(t2)])
    // fmt.Println(t4)
    // t2 = append(append(t2[0 : 3], 4), t4...)
    // fmt.Println(t4)         // ok
    // fmt.Println(t2)

    fmt.Println("ans:")

    // arr1 := [][]int{{1,3},{6,9}}
    // arr2 := []int{2,5}

    // arr1 := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
    // arr2 := []int{4,8}

    // arr1 := [][]int{}
    // arr2 := []int{5,7}

    // arr1 := [][]int{{1,5}}
    // // arr2 := []int{2,3}
    // arr2 := []int{2,7}

    // arr1 := [][]int{{5,7},{8,9}}
    // arr2 := []int{1,2}

    // arr1 := [][]int{{3,5},{7,8}}
    // arr2 := []int{6,10}
    // // arr2 := []int{6,7}

    arr1 := [][]int{{1,2},{11,22}}
    arr2 := []int{5,5}

    ans := insert(arr1, arr2)

    fmt.Println(ans)

}
