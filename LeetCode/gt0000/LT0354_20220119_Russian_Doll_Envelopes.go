// package sdq
package main

import (
    "fmt"
    "sort"
)



    // struct comp {
    //     bool operator()(vector<int>v1,vector<int>v2)
    //     {
    //         return (v1[0]<v2[0] || (v1[0]==v2[0] && v1[1]>v2[1]));
    //     }
    // };
    // int lengthOfLIS(vector<int>& nums) {
    //     vector<int> sub;
    //     for (int x : nums) {
    //         if (sub.empty() || sub[sub.size() - 1] < x) {
    //             sub.push_back(x);
    //         } else {
    //             auto it = lower_bound(sub.begin(), sub.end(), x); // Find the index of the smallest number >= x
    //             *it = x; // Replace that number with x
    //         }
    //     }
    //     return sub.size();
    // }
    // int maxEnvelopes(vector<vector<int>>& es) {
    //     int n=es.size();
    //     sort(es.begin(),es.end(),comp());
    //     vector<int>dp(n,1);
    //     int mx=1;
    //     vector<int>ls;
    //     for(int i=0;i<n;i++)
    //         ls.push_back(es[i][1]);
       
    //     return lengthOfLIS(ls);
    // }
// 。。 按 [0] 从小到达，然后找 LIS。 最长增长序列(?)
// 而且[0] 相等的时候 是 [1]降序， 所以 能LIS。 这个能确保 相同的[0]，根据LIS 的时候，不会 有 套娃。



// sort.Slice(e, func(i, j int) bool {
//     return e[i][0] < e[j][0] || (e[i][0] == e[j][0] && e[i][1] > e[j][1])
// })
// var lis []int
// for i, v := range e {
//     loc := sort.Search(len(lis), func(a int) bool {
//         return e[lis[a]][1] >= v[1]
//     })
//     if loc >= len(lis) {
//         lis = append(lis, i)
//     } else {
//         lis[loc] = i
//     }
// }
// return len(lis)


// for(int i = 1; i < size; ++i)
// {
//     maxrolls[i] = 1;
//     for(int j = i-1; j >= 0; --j)
//         if(envelopes[i].first>envelopes[j].first && envelopes[i].second>envelopes[j].second)
//             maxrolls[i] = max(maxrolls[i], maxrolls[j]+1);
//     maxroll = max(maxroll, maxrolls[i]);
// }
// sort后，从后往前，直接dp。


// Runtime: 285 ms, faster than 20.69% of Go online submissions for Russian Doll Envelopes.
// Memory Usage: 6.2 MB, less than 89.66% of Go online submissions for Russian Doll Envelopes.
// 记录套层。缩减耗时。
// 记录下 外层，如果其他的 大于这个外层，就没有必要计算。 但是。  而且 有dp了。 应该无所谓的。
// sort
func maxEnvelopes(envelopes [][]int) int {
    sz1 := len(envelopes)
    depths := make([]int, sz1)
    sort.Slice(envelopes, func(i, j int) bool {
        return envelopes[i][0] < envelopes[j][0]
    })

    // fmt.Println(envelopes)

    ans := 0
    for i := 0; i < sz1; i++ {
        t2 := dfsa354(envelopes, depths, i)
        if t2 > ans {
            ans = t2
        }
    }
    return ans
}

func dfsa354(ep [][]int, depths []int, idx int) int {
    if depths[idx] > 0 {
        return depths[idx]
    }
    ans := 1
    a0, a1 := ep[idx][0], ep[idx][1]
    for i := idx + 1; i < len(ep); i++ {
        if ep[i][0] > a0 && ep[i][1] > a1 {
            t2 := dfsa354(ep, depths, i)
            if t2 + 1 > ans {
                ans = t2 + 1
            }
        }
    }

    depths[idx] = ans
    return ans
}

func main_LT0354_20220119() {
// func main() {

    fmt.Println("ans:")


}
