// package sdq
package main

import (
	"fmt"
	"math"
)

// D D

// So cost[i] to finish time[i] + 1 walls,
// So cost[i] to finish time[i] + 1 walls,
// So cost[i] to finish time[i] + 1 walls,
// So cost[i] to finish time[i] + 1 walls,
// So cost[i] to finish time[i] + 1 walls,
// So cost[i] to finish time[i] + 1 walls,

// int n = cost.size();
// vector<int> dp(n + 1, 1e9);
// dp[0] = 0;
// for (int i = 0; i < n; ++i)
// 	for (int j = n; j > 0; --j)
// 		dp[j] = min(dp[j], dp[max(j - time[i] - 1, 0)] + cost[i]);
// return dp[n];

//
// 不知道 vscode的哪个配置，ctrl+s 就删除空行。 已经不勾了啊。还是这样。
// 已经好点了， 之前是 删除空行 + 删除空格，比如 a[i + j] 会变成 a[i+j] 。
// 现在是 删除空行。
// 应该也是 vscode go 的插件的问题， 我打开 txt， md， 就没有 删除空行的行为。
//

// Runtime23 ms
// Beats
// 77.78%
// Memory6.9 MB
// Beats
// 66.67%

// Paid painters will be used for a maximum of N/2 units of time.
// There is no need to use paid painter for a time greater than this.
// 我还是无法理解这个hint，在我看来： cost [1,333,333,333]  time[2,333,333,333]
// time[0] >= N/2 , 但是 只让paid printer 画[0], free 画[1,2,3] 肯定不行啊。
// 这个hint 应该是错的。

// **==So cost[i] to finish time[i] + 1 walls,     time[i] 是 free的， 1 是paid

func paintWalls(cost []int, time []int) int {
    sz1 := len(cost)
	vi := make([]int, sz1 + 1)			// [count paid+free wall] minCost
	for i := range vi {
		vi[i] = math.MaxInt
	}
	vi[0] = 0

	fmn := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i, tm := range time {
		cst := cost[i]

		for j := sz1; j >= 0; j-- {
			if vi[j] == math.MaxInt {
				continue
			}
			var t2 int = j + tm + 1;
			if t2 > sz1 {
				t2 = sz1
			}
			vi[t2] = fmn(vi[t2], vi[j] + cst)
		}
	}
	return vi[sz1]
}

// tle 2550 / 2557
// [idx][paid painter paint count] = cost
// [time][count] = min cost
// [cnt] = max(time)
// [cnt][time] = min(cost)
func paintWalls__tle__(cost []int, time []int) int {

	sz1 := len(cost)
	vmap := make([]map[int]int, sz1)		// [count of paid wall] map [time] minCost
	for i := range vmap {
		vmap[i] = make(map[int]int)
	}

	fmn := func(a, b int) int {
		if a == 0 {
			return b
		} else if b == 0 {
			return a
		} else if a > b {
			return b
		} else {
			return a
		}
	}
	ans := math.MaxInt
	for i := range cost {
		for j := i - 1; j >= 0; j-- {
			for k, v := range vmap[j] {
				if k + time[i] >= sz1 - j - 2 {
					ans = fmn(ans, v + cost[i])
				} else {
					vmap[j + 1][k + time[i]] = fmn(vmap[j + 1][k + time[i]], v + cost[i])
				}
			}
		}
		if time[i] >= sz1 - 1 {
			ans = fmn(ans, cost[i])
		} else {
			vmap[0][time[i]] = fmn(vmap[0][time[i]], cost[i])
		}
	}

	fmt.Println(vmap)

	return ans

	// sz1 := len(cost)
	// vvi := make([][]int, sz1)
	// for i := range vvi {
	// 	vvi[i] = make([]int, sz1)
	// 	for j := range vvi[i] {
	// 		vvi[i][j] = math.MaxInt
	// 	}
	// }

	// for i := range cost {

	// }

	// var sz1 int = len(cost)
	// var vi []int = make([]int, sz1 / 2 + 2);
	// for i := range vi {
	// 	vi[i] = math.MaxInt
	// }

	// fmx := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }

	// for i := range cost {
	// 	ct := cost[i]
	// 	tm := time[i]

	// 	for j := 1; j < len(vi) - 1; j++ {
	// 		if vi[j] == math.MaxInt {
	// 			break
	// 		}
	// 		vi[j + 1] = fmx(vi[j + 1], vi[j] + time)
	// 	}
	// }

	// var vvi [][]int = make([][]int, sz1)
	// for i := range vvi {
	// 	vvi[i] = make([]int, sz1)
	// 	for j := range vvi[i] {
	// 		vvi[i][j] = math.MaxInt
	// 	}
	// }

	// fmn := func(a, b int) int {
	// 	if a > b {
	// 		return b
	// 	}
	// 	return a
	// }

	// for i := range vvi {		// cost[i]
	// 	for j := range vvi[i] {		// count

	// 	}
	// }

}






// D

// int n = cost.size();
// vector<int> dp(n + 1, 1e9);
// dp[0] = 0;
// for (int i = 0; i < n; ++i)
//     for (int j = n; j > 0; --j)
//         dp[j] = min(dp[j], dp[max(j - time[i] - 1, 0)] + cost[i]);
// return dp[n];

// 2 3 4 2

// hint: ...
// discuss: 01 beibao.
// [time][wall count]
//
// func paintWalls(cost []int, time []int) int {

// }

// 1 : cost / time
// 2 : 0 / 1
// <= 500
// sum(time) >= 1*(sz1 - count(time))
// min sum(cost)
// dp[500][500] ?
// time[x] <= 500
// min cost that paid painter spend sz1/2 days
// not sz1/2,  it is changable...
func paintWalls_____err__(cost []int, time []int) int {
	sz1 := len(cost)
	fmna1 := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	var vi = make([]int, sz1)
	for i := range vi {
		vi[i] = math.MaxInt
	}
	// ans := math.MaxInt
	for i, tm := range time {
		// for j := 1; j <= sz1/2; j++ {
		for j := sz1 - 2; j > 0; j-- {
			if vi[j] != math.MaxInt {
				// if j+tm > sz1/2 {
				// 	ans = fmna1(ans, j+tm)
				if j+tm >= sz1 {
					vi[sz1-1] = fmna1(vi[sz1-1], vi[j]+cost[i])
				} else {
					vi[j+tm] = fmna1(vi[j+tm], vi[j]+cost[i])
				}
			}
		}
		vi[tm] = fmna1(cost[i], vi[tm])
	}
	return vi[sz1-1]
	// if sz1%2 == 0 {
	// 	ans = fmna1(ans, vi[sz1/2])
	// }
	// return ans
}

// func main_LT2742_20231014() {
func main() {

	// vi := []int{1,2,3,2}
	// v2 := []int{1,2,3,2}

	vi, v2 := []int{2,3,4,2}, []int{1,1,1,1}

	fmt.Println("ans:", paintWalls(vi, v2))

}
