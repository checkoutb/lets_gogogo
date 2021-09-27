
package main

import (
    "fmt"
)

// D D

// if(land[i][j]==1 && (j==0 || land[i][j-1]==0) && (i==0 || land[i-1][j]==0)){
// don't need fill          ==1 && isLeftTop corner

// if (land[i][j]) {
//     int x = i, y = j;
//     for (x = i; x < m && land[x][j]; ++x)
//         for (y = j; y < n && land[x][y]; ++y)
//             land[x][y] = 0;
//     res.push_back({i, j, x - 1, y - 1});
// }


// if (land[i][j] == 1) {
//     int x = i, y = j;
//     while (y < n && land[i][y])
//         ++y;                
//     for (; x < m && land[x][j]; ++x)
//         land[x][j] = y - j + 1;
//     res.push_back({i, j, x - 1, y - 1});
// }
// if (land[i][j] > 1)
//     j += land[i][j] - 1;
// 把长度保存在[x][j]，后续碰到的时候， [i][j] > 1  直接 j跳过长度个下标。


// var ret [][]int

// arr := make([]int, 4)
// arr[0], arr[1], arr[2], arr[3] = r, c, nr - 1, nc - 1
// ret = append(ret, arr)

// func findFarmland(land [][]int) [][]int {
//     R, C := len(land), len(land[0])
//     isTopLeft := func(r int, c int) bool {
//         return land[r][c] == 1 && (
//             r == 0 && c == 0 || r == 0 && land[r][c-1] == 0 ||
//             c == 0 && land[r-1][c] == 0 || 
//             r > 0 && c > 0 && land[r][c-1] == 0 && land[r-1][c] == 0)
//     }
//     var ret [][]int
//     for r := 0; r < R; r++ {
//         for c := 0; c < C; c++ {
//             if !isTopLeft(r, c) {
//                 continue
//             }

// res = append(res, []int{i,j,bottom[0],bottom[1]})




// Groups of farmland are rectangular in shape.
func findFarmland(land [][]int) [][]int {
    ans := make([][]int, 0, 10)
    sz1 := len(land)
    sz2 := len(land[0])
    for i, arr := range(land) {
        for j, val := range(arr) {
            if val == 1 {
                i2, j2 := 0, 0
                for i2 = i; i2 < sz1 && land[i2][j] == 1; i2++ {
                    
                }
                i2--
                for j2 = j; j2 < sz2 && land[i][j2] == 1; j2++ {}
                j2--
                ans = append(ans, []int{i, j, i2, j2})
                i3 := i
                for i3 <= i2 {
                    j3 := j
                    for j3 <= j2 {
                        land[i3][j3] = 2
                        j3++
                    }
                    i3++
                }
                // for i <= i2 {           // 会改变 下一次i的值吗？
                //     for j <= j2 {
                //         land[i][j] = 2
                //         j++         // ...zhege kendingde ...
                //     }
                // }
            }
        }
    }
    return ans
}



func main_LT1992_20210927() {
// func main() {

    // arr := [][]int{{1,0,0},{0,1,1},{0,1,1}}
    // arr := [][]int{{1,1},{1,1}}
    arr := [][]int{{1}}

    fmt.Println("ans:")

    var ans [][]int = findFarmland(arr)
    fmt.Println(ans)

}
