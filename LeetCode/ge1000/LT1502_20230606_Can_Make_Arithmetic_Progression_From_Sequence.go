// package sdq
package main

import (
    "fmt"
    "sort"
)



// Runtime5 ms
// Beats
// 25.51%
// Memory2.5 MB
// Beats
// 7.14%
func canMakeArithmeticProgression(arr []int) bool {
    sort.Ints(arr);
    diff := arr[1] - arr[0];
    for i := 2; i < len(arr); i++ {
        if diff != arr[i] - arr[i - 1] {
            return false;
        }
    }
    return true;
}


func main_LT1502_20230606() {
// func main() {

    fmt.Println("ans:")


}
