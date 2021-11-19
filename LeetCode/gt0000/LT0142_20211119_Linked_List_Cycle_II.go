// package sdq
package main

import (
    "fmt"
)



// while fast and fast.next:
//     slow = slow.next
//     fast = fast.next.next
//     if slow == fast:
//         break
// else:
//     return None
// .py




// Runtime: 0 ms, faster than 100.00% of Go online submissions for Linked List Cycle II.
// Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Linked List Cycle II.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 记得是，快慢指针相遇后， 快/慢 指针从头开始， 另一个不变。 相遇的时候就是了。  还是说 都是 慢速？ 应该是 都慢速
func detectCycle(head *ListNode) *ListNode {
    f, s := head, head
    for f != nil && f.Next != nil {
        f = f.Next.Next
        s = s.Next
        if f == s {
            break
        }
    }
    if f == nil || f.Next == nil {
        return nil
    }
    f = head
    for f != s {
        f = f.Next
        s = s.Next
    }
    return f
}


func main_LT0142_20211119() {
// func main() {

    fmt.Println("ans:")


}
