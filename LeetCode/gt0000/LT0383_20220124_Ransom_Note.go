// package sdq
package main

import (
    "fmt"
)


// std::multiset<char> mag(magazine.begin(), magazine.end());
// std::multiset<char> ransom(ransomNote.begin(), ransomNote.end());
// return std::includes(mag.begin(), mag.end(), ransom.begin(), ransom.end());


// Runtime: 4 ms, faster than 90.79% of Go online submissions for Ransom Note.
// Memory Usage: 4.2 MB, less than 39.47% of Go online submissions for Ransom Note.
func canConstruct(ransomNote string, magazine string) bool {
    arr := make([]int, 123)
    for _, v := range magazine {
        arr[v]++
    }
    for _, v := range ransomNote {
        arr[v]--
        if arr[v] < 0 {
            return false
        }
    }
    return true
}

func main_LT0383_20220124() {
// func main() {

    fmt.Println("ans:")


}
