// package sdq
package main

import (
    "fmt"
)


// if (head == null) return null;
// head.next = removeElements(head.next, val);
// return head.val == val ? head.next : head;



// Runtime: 11 ms, faster than 12.81% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 5 MB, less than 20.08% of Go online submissions for Remove Linked List Elements.
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{-1, head}
    t2 := dummy
    for t2.Next != nil {
        if t2.Next.Val == val {
            t2.Next = t2.Next.Next
        } else {
            t2 = t2.Next
        }
    }
    return dummy.Next
}

func main_LT0203_20211123() {
// func main() {

    fmt.Println("ans:")


}
