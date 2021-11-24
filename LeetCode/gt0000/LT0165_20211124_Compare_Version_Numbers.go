// package sdq
package main

import (
    "fmt"
    "strings"
    "strconv"
)


// vs1 := strings.Split(version1, ".")



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Compare Version Numbers.
// Memory Usage: 2.2 MB, less than 7.14% of Go online submissions for Compare Version Numbers.
func compareVersion(version1 string, version2 string) int {
    fun1 := func (c rune) bool {
        return c == '.'
    }
    arr1 := strings.FieldsFunc(version1, fun1)
    arr2 := strings.FieldsFunc(version2, fun1)

    // fmt.Println(arr1)
    // fmt.Println(arr2)

    ar1 := str2inta165(arr1)
    ar2 := str2inta165(arr2)

    // fmt.Println(ar1)
    // fmt.Println(ar2)

    for i := 0; i < len(ar1) && i < len(ar2); i++ {
        if ar1[i] > ar2[i] {
            return 1
        } else if ar1[i] < ar2[i] {
            return -1
        }
    }
    if len(ar1) > len(ar2) {
        return 1
    } else if len(ar1) < len(ar2) {
        return -1
    } else {
        return 0
    }
}

func str2inta165(arr []string) []int {
    sz1 := len(arr)
    ans := make([]int, sz1)
    for i, v := range arr {
        ans[i], _ = strconv.Atoi(v)
    }
    for i := sz1 - 1; i >= 0; i-- {
        if ans[i] != 0 {
            ans = ans[0 : i + 1]
            break
        }
    }
    return ans
}

// func main_LT0165_20211124() {
func main() {

    fmt.Println("ans:")

    s1, s2 := "1.01", "1.0000"

    fmt.Println(compareVersion(s1, s2))

}
