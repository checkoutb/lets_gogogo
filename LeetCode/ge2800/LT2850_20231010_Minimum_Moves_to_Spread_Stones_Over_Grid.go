// package sdq
package main

import (
    "fmt"
    // "math"
)

// D D

// func dfs(cell int, moves int, grid [][]int, res *int) {
//     if moves >= *res {
//         return
//     }
//     // base case
//     if cell == 9 {
//         // fmt.Println(moves)
//         *res = min(*res, moves)
//         return
//     }



// 这个谁提交的？？？ 直接拔高了一个层次。。
/*
func minimumMoves(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	src := m * n   // 超级源点
	dst := src + 1 // 超级汇点
	type edge struct{ to, rid, cap, cost int }
	g := make([][]edge, m*n+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], edge{to, len(g[to]), cap, cost})
		g[to] = append(g[to], edge{from, len(g[from]) - 1, 0, -cost})
	}
	for x, row := range grid {
		for y, v := range row {
			if v > 1 {
				addEdge(src, x*n+y, v-1, 0)
				for i, r := range grid {
					for j, w := range r {
						if w == 0 {
							addEdge(x*n+y, i*n+j, 1, abs(x-i)+abs(y-j))
						}
					}
				}
			} else if v == 0 {
				addEdge(x*n+y, dst, 1, 0)
			}
		}
	}

	// 下面是最小费用最大流模板
	const inf int = 1e9
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = 1e9
		}
		dist[src] = 0
		inQ := make([]bool, len(g))
		inQ[src] = true
		q := []int{src}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[dst] < inf
	}
	ek := func() (maxFlow, minCost int) {
		for spfa() {
			// 沿 st-end 的最短路尽量增广
			minF := inf
			for v := dst; v != src; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := dst; v != src; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dist[dst] * minF
		}
		return
	}
	_, cost := ek()
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
*/



// class Solution {
//     public:
//         int minimumMoves(vector<vector<int>>& grid) {
//             // Base Case
//             int t = 0;
//             for (int i = 0; i < 3; ++i)
//                 for (int j = 0; j < 3; ++j)
//                     if (grid[i][j] == 0)
//                         t++;
//             if (t == 0)
//                 return 0;
            
//             int ans = INT_MAX;
//             for (int i = 0; i < 3; ++i)
//             {
//                 for (int j = 0; j < 3; ++j)
//                 {
//                     if (grid[i][j] == 0)
//                     {
//                         for (int ni = 0; ni < 3; ++ni)
//                         {
//                             for (int nj = 0; nj < 3; ++nj)
//                             {
//                                 int d = abs(ni - i) + abs(nj - j);
//                                 if (grid[ni][nj] > 1)
//                                 {
//                                     grid[ni][nj]--;
//                                     grid[i][j]++;
//                                     ans = min(ans, d + minimumMoves(grid));
//                                     grid[ni][nj]++;
//                                     grid[i][j]--;
//                                 }
//                             }
//                         }
//                     }
//                 }
//             }
//             return ans;
//         }
//     };
    




// Runtime15 ms
// Beats
// 54.43%
// Memory2.5 MB
// Beats
// 48.10%

// 9!
func minimumMoves(grid [][]int) int {
    emp := []Pair{}
    dup := []Pair{}

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[i][j] == 0 {
                emp = append(emp, Pair{i, j})
            } else if grid[i][j] > 1 {
                for k := 1; k < grid[i][j]; k++ {
                    dup = append(dup, Pair{i, j});
                }
            }
        }
    }

    ans := dfsa1(emp, dup, 0, 0)
    return ans
}

func dfsa1(emp, dup []Pair, idx, stp int) int {
    if idx == len(emp) {
        return stp
    }
    ans := 123123123
    t2 := 0
    for i := 0; i < len(dup); i++ {
        if dup[i].fst != -1 {
            t2 = dup[i].fst
            dup[i].fst = -1
            ans = mn(ans, dfsa1(emp, dup, idx + 1, stp + absa1(emp[idx].fst - t2) + absa1(emp[idx].snd - dup[i].snd)));
            dup[i].fst = t2
        }
    }
    return ans
}

func mn(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

type Pair struct {
    fst int
    snd int
}



// g...

// 3 2 0
// 0 1 0
// 0 3 0

// O(9 * 9)
// min distance to empty cell
func minimumMoves__(grid [][]int) int {
    var ans = 0
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[i][j] == 0 {
                ni, nj := 100, 100
                for a := 0; a < 3; a++ {
                    for b := 0; b < 3; b++ {
                        if grid[a][b] > 1 && (absa1(ni - i) + absa1(nj - j)) > (absa1(a - i) + absa1(b - j)) {
                            ni, nj = a, b
                        }
                    }
                }
                ans += absa1(ni - i) + absa1(nj - j)
                grid[ni][nj]--
                grid[i][j] = 1
            }
        }
    }
    return ans
}

func absa1(a int) int {
    if a < 0 {
        a = -a
    }
    return a
}


// func main_LT2850_20231010() {
func main() {

    vv := [][]int{{3,2,0},{0,1,0},{0,3,0}}
    
    fmt.Println("ans:", minimumMoves(vv))


}
