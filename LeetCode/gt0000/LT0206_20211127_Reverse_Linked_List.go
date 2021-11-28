// package sdq
package main

import (
    "fmt"
)


// dfs


// Runtime: 5 ms, faster than 15.69% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.8 MB, less than 36.66% of Go online submissions for Reverse Linked List.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 
func reverseList(head *ListNode) *ListNode {
    t2 := head
    var pre *ListNode = nil
    for t2 != nil {
        t2.Next, pre, t2 = pre, t2, t2.Next
    }
    return pre

//    dummy := &ListNode{ -1, head }
//    t2 := dummy
//    for t2.Next != nil {

//    }
//    return dummy.Next
}


func main_LT0206_20211127() {
// func main() {

    fmt.Println("ans:")


}
