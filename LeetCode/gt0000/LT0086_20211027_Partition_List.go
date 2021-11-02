// package sdq
package main

import (
    "fmt"
)



// chabuduo


// Runtime: 0 ms, faster than 100.00% of Go online submissions for Partition List.
// Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Partition List.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    dummy := ListNode{-1, nil}       // < x
    dy2 := ListNode{-1, nil}         // >= x
    dy1tail := &dummy
    dy2t := &dy2
    for head != nil {
        t2 := head.Next
        if head.Val < x {
            dy1tail.Next = head
            dy1tail = dy1tail.Next
        } else {
            dy2t.Next = head
            dy2t = dy2t.Next
        }
        head = t2
    }
    // if dy1tail == nil {
    //     dummy.Next = dy2.Next
    // } else {
        dy1tail.Next = dy2.Next
    // }
    dy2t.Next = nil
    return dummy.Next
}

func main_LT0086_20211027() {
// func main() {

    fmt.Println("ans:")


}
