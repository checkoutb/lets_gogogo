// package sdq
package main

import (
    "fmt"
)



// while(fast.next != null && fast.next.next != null)

// for slow!=fast {
//     if fast == nil || fast.Next == nil {
//         return false
//     }


// try:
//     slow = head
//     fast = head.next
//     while slow is not fast:
//         slow = slow.next
//         fast = fast.next.next
//     return True
// except:
//     return False


// while head:
//     if sys.getrefcount(head) > 4:
//         return True
//     head = head.next
// https://docs.python.org/3/library/sys.html#sys.getrefcount




// Runtime: 8 ms, faster than 86.55% of Go online submissions for Linked List Cycle.
// Memory Usage: 4.4 MB, less than 100.00% of Go online submissions for Linked List Cycle.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    f, s := head.Next, head
    for f != nil {
        if f == s {
            return true
        }
        s = s.Next
        f = f.Next
        if f != nil {
            f = f.Next
        }
    }
    return false
}

func main_LT0141_20211119() {
// func main() {

    fmt.Println("ans:")


}
