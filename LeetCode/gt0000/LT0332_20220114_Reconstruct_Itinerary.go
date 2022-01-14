// package sdq
package main

import (
    "fmt"
    "sort"
)




// 为什么 dfs的时候 不需要 size 呢。

// public void dfs(String departure) {
//     PriorityQueue<String> arrivals = flights.get(departure);
//     while (arrivals != null && !arrivals.isEmpty())
//         dfs(arrivals.poll());
//     path.addFirst(departure);
// }

// map<string, multiset<string>> targets;
// vector<string> route;
// void visit(string airport) {
//     while (targets[airport].size()) {
//         string next = *targets[airport].begin();
//         targets[airport].erase(targets[airport].begin());
//         visit(next);
//     }
//     route.push_back(airport);
// }

// 。因为 如果 dfs 里 size 不符合的话， 说明 还有 ticket，那么 while 不会停止。 还会继续。 
// 如果 符合，那么 最终 所有的 ticket 都被 删除了。 while 就没有了。
// 而且 正确的 时候 会删完 所有，这样 route 就是 正确的。
// 但是还是有点。。


// Runtime: 8 ms, faster than 100.00% of Go online submissions for Reconstruct Itinerary.
// Memory Usage: 6.8 MB, less than 35.07% of Go online submissions for Reconstruct Itinerary.
func findItinerary(tickets [][]string) []string {
    mapout := map[string][]string{}
    sz1 := len(tickets)
    for _, ar := range tickets {
        t1, t2 := ar[0], ar[1]
        mapout[t1] = append(mapout[t1], t2)
    }
    for k, _ := range mapout {
        sort.Strings(mapout[k])
    }
    st := "JFK"
    arr := dfs332a(mapout, sz1, st)
    arr = append(arr, st)
    for i, j := 0, len(arr) - 1; i < j; i++ {
        arr[i], arr[j] = arr[j], arr[i]
        j--
    }
    return arr
}

func dfs332a(mapout map[string][]string, sz int, prev string) []string {
    ans := []string{}
    if len(mapout[prev]) == 0 {
        // if sz == 1 {
        //     return []string{prev}
        // } else {
            return ans
        // }
    }
    ar := mapout[prev]
    // fmt.Println(prev, ar)
    // fmt.Println(mapout, prev)
    for idx, nxt := range ar {
        // mapout[prev] = append(ar[0 : idx], ar[idx + 1 : ]...)
        // fmt.Println(idx, nxt, mapout[prev], ar)
        ar2 := make([]string, len(ar))
        copy(ar2, ar)
        mapout[prev] = append(ar2[0 : idx], ar2[idx + 1 : ]...)
        // fmt.Println(ar2, ar, idx, mapout[prev])         // ar2 会被修改，但是 长度不变。 为什么会被修改呢 。 最下面。 是因为 这里的 ar 只是指针， 指向了 同一个底层数组。

        arr := dfs332a(mapout, sz - 1, nxt)
        if len(arr) + 1 == sz {
            ans = append(arr, nxt)
            break
        }
    }
    mapout[prev] = ar
    return ans
}

// error。。。{{"JFK","KUL"},{"JFK","NRT"},{"NRT","JFK"}}
// 票用完，并且 字典顺序 最小。
// 出度 入度 相等， 不可能是 出发 或 结束
// 出度 > 入度， 就是 开始， 应该只有 一个 结点吧。
// 出度 < 入度， 就是 结束， 也应该只有 一个吧？
// You may assume all tickets form at least one valid itinerary.    说明可能有 多个？ 那就是 可能有多个 开始啊。
// If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.
// 不，感觉还是 只有一个，并且 票要全部用完。 上面只是 如果 有 多个双向，选 最小的。

// Your input
// [["JFK","ATL"],["ATL","JFK"]]
// Output
// ["ATL","JFK","ATL"]
// Expected
// ["JFK","ATL","JFK"]
// ++ . who departs from "JFK" ... 我以为是例子。。

func findItinerary2(tickets [][]string) []string {
    mapout := map[string][]string{}      // key -> value
    // mapin := map[string][]string{}       // key <- value
    // sort.
    mapincnt := map[string]int{}
    // mapcnt := map[string]int{}
    
    for _, ar := range tickets {
        t1, t2 := ar[0], ar[1]
        mapout[t1] = append(mapout[t1], t2)
        mapincnt[t2]++
    }
    // st := ""
    // en := ""
    // mn := "ZZZ"
    st := "JFK"
    en := "ZZZ"
    for k, _ := range mapout {
        sz1 := len(mapout[k])
        sort.Strings(mapout[k])
        // if sz1 > mapincnt[k] {
        //     st = k
        // }
        if sz1 < mapincnt[k] {
            en = k
        }
        // if k < mn {
        //     mn = k
        // }
    }
    if en == "ZZZ" {
        en = st
    }
    // fmt.Println(mn)
    // fmt.Println(mapout)
    // if st == "" {     // ... circle...
    //     st, en = mn, mn
    // }

    // st, en := "JFK", 
    ans := []string{}
    for st != "" {
        // fmt.Println(st)
        ans = append(ans, st)
        if len(mapout[st]) > 0 {
            t2 := mapout[st][0]
            if t2 == en && len(mapout[st]) > 1 {
                t2 = mapout[st][1]
                mapout[st] = append(mapout[st][0 : 1], mapout[st][2 : ]...)
            } else {
                mapout[st] = mapout[st][1 : ]
            }
            st = t2
        } else {
            // fmt.Println(st)
            st = ""
        }
    }

    return ans
}

func main_LT0332_20220114() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}
    // arr := [][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}
    // arr := [][]string{{"JFK","ATL"},{"ATL","JFK"}}
    arr := [][]string{{"JFK","KUL"},{"JFK","NRT"},{"NRT","JFK"}}        // ... 感觉得 dfs 硬算了。。。

    ans := findItinerary(arr)

    fmt.Println(ans)

    fmt.Println("-----------------")

    ar := []int{1,2,3,4,5}
    ar2 := ar[0: 2]
    fmt.Println(ar, ar2)
    ar3 := ar[3 : 5]
    fmt.Println(ar, ar2, ar3)

    ar4 := append(ar2, ar3...)          // append 导致 ar3 被修改了。 似乎是把 数据 提前了。
    fmt.Println(ar2, ar3, ar4)
    fmt.Println(ar)     // 切片是指针，所以 ar 也被改了。

    // 应该是 因为 指针的问题。

    // ar4 和 ar2 的开始 是同一个位置， 所以 上面 跳过了一个元素，就 导致 ar3 整体前移 一个。
    // 所以 ar 系列 (除了ar2) 全改了。

}
