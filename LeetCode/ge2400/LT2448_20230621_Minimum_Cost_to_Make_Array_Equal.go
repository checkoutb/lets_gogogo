// package sdq
package main

import (
    "fmt"
    "sort"
)




// Runtime57 ms
// Beats
// 77.78%
// Memory11.2 MB
// Beats
// 22.22%

type Pair struct {
    fst int64
    snd int64
}

func minCost(nums []int, cost []int) int64 {
    vp := make([]Pair, len(nums));
    for i, v := range nums {
        vp[i] = Pair{int64(v), int64(cost[i])};
    }
    // sort.Sort(vp);
    sort.Slice(vp, func(i, j int) bool { return vp[i].fst < vp[j].fst; });

    // fmt.Println(vp);

    var ans int64 = 0;         // min
    var t2 int64 = vp[0].fst;     // mid point
    var t3 int64 = 0;             // temp ans
    var cstl int64 = 0;
    var cstr int64 = 0;

    abs := func(num int64) int64 {
        if (num < 0) {
            num = -num;
        }
        return num;
    }

    for _, p := range vp {
        t3 += abs(p.fst - t2) * p.snd;
        cstr += p.snd;
    }
    ans = t3;
    
    for i := 1; i < len(vp); i++ {
        cstr -= vp[i - 1].snd;
        cstl += vp[i - 1].snd;
        t3 -= cstr * (vp[i].fst - vp[i - 1].fst);
        t3 += cstl * (vp[i].fst - vp[i - 1].fst);

        // fmt.Println(i, t3, cstl, cstr);

        if t3 > ans {
            break;
        }
        ans = t3;
    }

    return ans;
}

// // parabola, min
// func minCost_1(nums []int, cost []int) int64 {
//     vp := make([]Pair, len(nums));
//     for i, v := range nums {
//         vp[i] = Pair{v, cost[i]};
//     }
//     sort.Sort(vp);

//     var st int = 0;
//     var en int = len(nums) - 1;

//     fun1 := func(idx int) int64{
//         var ans int64 = 0;
//         var t2 int64 = vp[idx].fst;
//         for i, p := range vp {
//             if i < idx {
//                 ans += (t2 - p.fst) * p.snd;
//             } else {
//                 ans += (p.fst - t2) * p.snd;
//             }
//         }
//         return ans;
//     }

//     while (st <= en)
//     {
//         md := (st + en) / 2;
//         md2 := md + 1;
//     }
//     return -1;
// }



// func main_LT2448_20230621() {
func main() {

    vi := []int{1,3,5,2};
    v2 := []int{2,3,1,14};

    fmt.Println("ans:", minCost(vi, v2));
}
