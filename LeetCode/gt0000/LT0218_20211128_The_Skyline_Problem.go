// package sdq
package main

import (
    "fmt"
    "sort"
)



// https://briangordon.github.io/2014/08/the-skyline-problem.html



// Step 1:
    // Use a multimap to sort all boundary points. 
    // For a start point of an interval, let the height be positive; 
    // otherwise, let the height be negative. Time complexity: O(n log n)
// Step 2:
    // Use a multiset (rather than a heap/priority_queue) to maintain the current set of heights to be considered. 
    // If a new start point is met, insert the height into the set, otherwise, delete the height. 
    // The current max height is the back() element of the multiset. 
    // For each point, the time complexity is O(log n). The overall time complexity is O(n log n).
// Step 3:
    // Delete the points with equal heights. Time: O(n)

// // Step 1:
// multimap<int, int> coords;
// for (const vector<int> & building : buildings) {
// 	coords.emplace(building[0], building[2]);
// 	coords.emplace(building[1], -building[2]);
// }

// // Step 2:
// multiset<int> heights = { 0 };
// map<int, int> corners;
// for (const pair<int, int> & p : coords) {
// 	if (p.second > 0) {
// 		heights.insert(p.second);
// 	}
// 	else {
// 		heights.erase(heights.find(-p.second));
// 	}
// 	int cur_y = *heights.rbegin();
// 	corners[p.first] = cur_y;
// }

// // Step 3:
// function<bool(pair<int, int> l, pair<int, int> r)> eq2nd = [](pair<int, int> l, pair<int, int> r){ return l.second == r.second;  };
// vector<pair<int, int>> results;
// unique_copy(corners.begin(), corners.end(), back_insert_iterator<vector<pair<int, int>>>(results), eq2nd);
// return results;









// 很难，但是又不简单。。。很少碰到这种类型的。 没有思路。
// Runtime: 36 ms, faster than 41.03% of Go online submissions for The Skyline Problem.
// Memory Usage: 8.1 MB, less than 12.82% of Go online submissions for The Skyline Problem.
func getSkyline(buildings [][]int) [][]int {
    map2 := map[int]int{}
    stall := map[int][]int{}
    enall := map[int][]int{}
    arr := []int{}
    for _, v := range buildings {
        st,en,hi := v[0],v[1],v[2]
        stall[st] = append(stall[st], hi)
        enall[en] = append(enall[en], hi)
        arr = append(arr, st)
        arr = append(arr, en)
    }
    sort.Ints(arr)
    ans := [][]int{}
    lst := -1
    mxh := -1
    for i := 0; i < len(arr); i++ {
        t2 := arr[i]
        if t2 == lst {
            continue
        }
        mxh2 := mxh
        if _, ok := stall[t2]; ok {
            for _, v := range stall[t2] {
                map2[v]++
                if v > mxh {
                    mxh = v
                }
            }
        }
        if _, ok := enall[t2]; ok {
            for _, v := range enall[t2] {
                map2[v]--
            }
        }
        if map2[mxh] == 0 {
            mxh = getMaxHigh_Remove0_a218(map2)
        }
        if mxh != mxh2 {
            ans = append(ans, []int{t2, mxh})
        }
        lst = arr[i]
    }
    return ans
}

// func getSkyline(buildings [][]int) [][]int {
//     map2 := map[int]int{}
//     mapst := map[int]int{}
//     mapen := map[int]int{}
//     arr := []int{}              // 怀念pri_que
//     // sz1 := len(buildings)
//     for i, v := range buildings {
//         if mapst[v[0]] < v[2] {
//             mapst[v[0]] = v[2]
//         }
//         if mapen[v[1]] < v[2] {
//             mapst[v[1]] = v[2]
//         }
//         arr = append(arr, v[0])
//         arr = append(arr, v[1])
//     }
//     sort.Ints(arr)
//     sz2 := len(arr)
//     ans := [][]int{}
//     lst := -1
//     mxh := -1
//     for i := 0; i < sz2; i++ {
//         if arr[i] == lst {
//             continue
//         }
//         t2 := arr[i]
//         if _, ok := mapen[t2]; ok {
//             map2[mapen[t2]]--
//         }
//         if _, ok := mapst[t2]; ok {
//             map2[mapst[t2]]++               // 只记录 最高值，有问题的，没有办法清除干净。
//         }

//         if t2 == mxh && map2[t2] == 0 {
//             mxh = getMaxHigh_Remove0_a218(map2)
//         }
//         lst = arr[i]
//     }
// }

func getMaxHigh_Remove0_a218(map2 map[int]int) int {
    mx := 0         // 平地。。
    for k, v := range map2 {
        if v == 0 {
            delete(map2, k)         // ?   可以的，遍历时删除 是安全的。  删不删呢？
        } else {
            if k > mx {
                mx = k
            }
        }
    }
    return mx
}


// // 分段，然后判断每段的 最大高度，合并。
// func getSkyline(buildings [][]int) [][]int {
//     sz1 := len(buildings)
//     mapst := map[int]int{}
//     mapen := map[int]int{}
//     arr := []int{}
//     for _m v := range buildings {
//         if mapst[v[0]] < v[2] {
//             mapst[v[0]] = v[2]
//         }
//         if mapen[v[1]] < v[2] {
//             mapen[v[1]] = v[2]
//         }
//         arr = append(arr, v[0])
//         arr = append(arr, v[1])
//     }
//     ans := [][]int{}
//     sort.Ints(arr)
//     lst := -1
//     lsten := -1
//     lsthi := -1
//     for i := 0; i < len(arr); i++ {
//         if arr[i] == lst {
//             continue
//         }
//         hi := mapst[v[i][0]]
//         if hi > lsthi {

//         } else {

//         }
//         lst = arr[i]
//     }



    // arr := []int{}
    // sz1 := len(buildings)
    // for _, v := range buildings {
    //     arr = append(arr, v[0])
    //     arr = append(arr, v[1])
    // }
// }

// // gg
// // 高度, 结束
// // 后续 且在 right内 第一个高于它的值。 如果right内没有比它高的，那么 它的right 就是 下一个点。
// func getSkyline2(buildings [][]int) [][]int {
//     stk := [][]int{ []int{buildings[0][0], buildings[0][2]} }
//     ans := [][]int{}
//     sz1 := len(buildings)
//     deal := 0
//     // ans = append(ans, []int{buildings[0][0], buildings[0][2]})
//     // for i := 1; i < sz1; i++ {
//     for deal < sz1 {
//         for j := deal + 1; j < sz1; j++ {
//             if buildings[j][0] < buildings[deal][1] {
//                 if buildings[j][2] > buildings[deal][2] {
//                     deal = j
//                     break
//                 }
//             } else {
//                 break
//             }
//         }
//     }
//     return ans
// }

// func main_LT0218_20211128() {
func main() {

    fmt.Println("ans:")

    // arr := [][]int{{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}
    arr := [][]int{{0,2,3},{2,5,3}}

    ans := getSkyline(arr)

    fmt.Println(ans)

}
