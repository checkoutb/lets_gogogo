// package sdq
package main

import (
    "fmt"
)




// Runtime0 ms
// Beats
// 100%
// Memory3.6 MB
// Beats
// 34.21%
// x/y = x2/y2  x y2 = y x2
func checkStraightLine(coordinates [][]int) bool {
    x := coordinates[0][0] - coordinates[1][0];
    y := coordinates[0][1] - coordinates[1][1];
    // for i in range(2 to len(coordinates)) {
    for i := 2; i < len(coordinates); i++ {
        x2, y2 := coordinates[i][0] - coordinates[1][0], coordinates[i][1] - coordinates[1][1];
        if x * y2 != x2 * y {
            return false;
        }
    }
    return true;
}


// func main_LT1232_20230605() {
func main() {

    fmt.Println("ans:")

    vvi := [][]int{{1,2},{2,3},{3,4},{4,5}};        // [3][2]int != [][]int

    fmt.Println(checkStraightLine(vvi));
}
