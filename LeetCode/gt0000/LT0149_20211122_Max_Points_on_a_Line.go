// package sdq
package main

import (
    "fmt"
    "math"
)


// for (int i = 0; i < n - 1; i++) {
//     for (int j = i + 1; j < n; j++) {
//         int[] point_a = points[i];
//         int[] point_b = points[j];
//         int abDeltaX = point_b[0] - point_a[0];
//         int abDeltaY = point_b[1] - point_a[1];
//         int count = 2;
//         for (int k = j + 1; k < n; k++) {
//             int[] point_c = points[k];
//             int acDeltaX = point_c[0] - point_a[0];
//             int acDeltaY = point_c[1] - point_a[1];
//             if (abDeltaY * acDeltaX == abDeltaX * acDeltaY) {
//                 count += 1;
//             }
//         }
//         max = Math.max(max, count);
//     }
// }

// (y3-y1)/(x3-x1)=(y2-y1)/(x2-x1)=a   除法 改成 乘法



// 精度真的。。。 估计是 两边都除以 gcd。
// Runtime: 4 ms, faster than 84.06% of Go online submissions for Max Points on a Line.
// Memory Usage: 7.1 MB, less than 11.59% of Go online submissions for Max Points on a Line.
// y = a*x + b
// a = (y2-y1)/(x2-x1)      // x2==x1
// b = y1 - a*x1
func maxPoints(points [][]int) int {
    map2 := map[float64]map[float64]int{}
    var x1,x2,y1,y2,t1,t2 float64
    for i, v := range points {
        // fmt.Println(i, ", ", v)
        x1, y1 = float64(v[0]), float64(v[1])
        for j := i + 1; j < len(points); j++ {
            x2, y2 = float64(points[j][0]), float64(points[j][1])
            if x1 == x2 {
                t1, t2 = float64(2100000000), x1
                // map2[t1][t2]++
            } else {
                t1 = (y2 - y1) / (x2 - x1)
                t2 = y1 - t1 * x1
                // map2[t1][t2]++
            }
            // fmt.Println(t1, ", ", t2, ", ", i, ", ", j)
            t1, t2 = math.Trunc(t1 * 1e11 + 0.5) * 1e-11, math.Trunc(t2 * 1e11 + 0.5) * 1e-11
            if _, ok := map2[t1]; !ok {
                map2[t1] = map[float64]int{}
            }
            map2[t1][t2]++          // ..4point, 3+2+1 == 6   x^2 + x - 12 = 0   (x-3)*(x+4)   x=3   3+1, sqrt
        }
    }

    ans := 0
    for _, v := range map2 {
        for _, v2 := range v {
            if v2 > ans {
                ans = v2
            }
        }
    }
    // sqr := math.Sqrt(float64(ans))
    ans = int(math.Floor(math.Sqrt(float64(ans * 2)))) + 1
    return ans
}


func main_LT0149_20211122() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}}
    // arr := [][]int{{0,1},{0,0},{0,4},{0,-2},{0,-1},{0,3},{0,-4}}

    arr := [][]int{{54,153},{1,3},{0,-72},{-3,3},{12,-22},{-60,-322},{0,-5},{-5,1},{5,5},{36,78},
    {3,-4},{5,0},{0,4},{-30,-197},{-5,0},{60,178},{0,0},{30,53},{24,28},{4,5},{2,-2},{-18,-147},{-24,-172},
    {-36,-222},{-42,-247},{2,3},{-12,-122},{-54,-297},{6,-47},{-5,3},{-48,-272},{-4,-2},{3,-2},{0,2},{48,128},{4,3},{2,4}}

    // arr := [][]int{{5151,5150},{0,0},{5152,5151}}

    ans := maxPoints(arr)

    fmt.Println(ans)

}
