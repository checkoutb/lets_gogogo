// package sdq
package main

import (
    "fmt"
    "reflect"
)


// 切掉第一行，然后 逆时针转90度， 继续切掉第一行， 逆时针90度。。。

// return matrix and [*matrix.pop(0)] + self.spiralOrder([*zip(*matrix)][::-1])

// matrix[0] ? matrix.shift + spiral_order(matrix.transpose.reverse) : []



// for len(ret) < n*m {
//     for i := left; i <= right && len(ret) < n*m; i++ {
//       ret = append(ret, matrix[up][i])
//     }
//     for i := up+1; i <= down-1 && len(ret) < n*m; i++ {
//       ret = append(ret, matrix[i][right])
//     }
//     for i := right; i >= left && len(ret) < n*m; i-- {
//       ret = append(ret, matrix[down][i])
//     }
//     for i := down-1; i >= up+1 && len(ret) < n*m; i-- {
//       ret = append(ret, matrix[i][left])
//     }
//     up++
//     down--
//     left++
//     right--
//   }


// vector<vector<int> > dirs{{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
// vector<int> res;
// int nr = matrix.size();     if (nr == 0) return res;
// int nc = matrix[0].size();  if (nc == 0) return res;
// vector<int> nSteps{nc, nr-1};                            // 记录了 这一次可以走 多少步。
// int iDir = 0;   // index of direction.
// int ir = 0, ic = -1;    // initial position
// while (nSteps[iDir%2]) {
//     for (int i = 0; i < nSteps[iDir%2]; ++i) {
//         ir += dirs[iDir][0]; ic += dirs[iDir][1];
//         res.push_back(matrix[ir][ic]);
//     }
//     nSteps[iDir%2]--;
//     iDir = (iDir + 1) % 4;
// }




// Runtime: 2 ms, faster than 31.47% of Go online submissions for Spiral Matrix.
// Memory Usage: 2.1 MB, less than 13.49% of Go online submissions for Spiral Matrix.
func spiralOrder(matrix [][]int) []int {
    used := 10000
    sz1, sz2 := len(matrix), len(matrix[0])
    ans := make([]int, sz1 * sz2)
    arr := []func(int, int)(int, int, int){
        func(a int, b int)(int, int, int){
            a1, b1, c1 := a, b + 1, 0
            if b1 == sz2 || matrix[a1][b1] == used {
                a1, b1, c1 = a + 1, b, 1
            }
            return a1, b1, c1
        },func(a int, b int)(int, int, int){
            a1, b1, c1 := a + 1, b, 1
            if a1 == sz1 || matrix[a1][b1] == used {
                a1, b1, c1 = a, b - 1, 2
            }
            return a1, b1, c1
        },func(a int, b int)(int, int, int){
            a1, b1, c1 := a, b - 1, 2
            if b1 < 0 || matrix[a1][b1] == used {
                a1, b1, c1 = a - 1, b, 3
            }
            return a1, b1, c1
        },func(a int, b int)(int, int, int){
            a1, b1, c1 := a - 1, b, 3
            if a1 < 0 || matrix[a1][b1] == used {
                a1, b1, c1 = a, b + 1, 0
            }
            return a1, b1, c1
        }}
    
    a, b, c := 0, 0, 0
    mxsz := sz1 * sz2
    for i := 0; i < mxsz; i++ {
        ans[i] = matrix[a][b]
        matrix[a][b] = used
        a, b, c = arr[c](a, b)
    }
    return ans
}


func main_LT0054_20211020() {
// func main() {
    var fff func(int) = func(a int) {fmt.Println("asdasd ", a)}
    fff(5)
    fmt.Println(reflect.TypeOf(fff).String())

    fmt.Println("ans:")
    
    // arr := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
    arr := [][]int{{1,2,3},{4,5,6},{7,8,9}}

    ans := spiralOrder(arr)

    fmt.Println(ans)


}
