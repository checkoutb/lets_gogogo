// package sdq
package main

import (
    "fmt"
)





// Runtime: 3 ms, faster than 94.51% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 3.2 MB, less than 38.64% of Go online submissions for Two Sum II - Input Array Is Sorted.
// 应该是 2个指针逼近，但不知道怎么证明。 总觉得会跳过。。
func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1
    for numbers[l] + numbers[r] != target {
        if numbers[l] + numbers[r] > target {
            r--
        } else {
            l++
        }
    }
    return []int{l + 1, r + 1}
}


func main_LT0167_20211124() {
// func main() {

    fmt.Println("ans:")


}
