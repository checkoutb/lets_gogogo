// package sdq
package main

import (
    "fmt"
)









// t := &ListNode{
//     Next: head,
// }
// c := t
// for c.Next != nil && c.Next.Next != nil {
//     if c.Next.Val == c.Next.Next.Val {
//         v := c.Next.Val
//         for c.Next != nil && c.Next.Val == v {
//             c.Next = c.Next.Next
//         }
//     } else {
//         c = c.Next
//     }
// }
// return t.Next



// ListNode dummy = new ListNode(0);
// dummy.next = head;
// ListNode prev = dummy;
// while(head != null){
//     if(head.next != null && head.val == head.next.val){
//         while(head.next != null && head.val == head.next.val){
//             head = head.next;
//         }
//         prev.next = head.next;               // 这里并没有移动 prev.
//     } else {
//         prev = prev.next;
//     }
//     head = head.next;
// }
// return dummy.next;



// while(cur!=null){
//     while(cur.next!=null&&cur.val==cur.next.val){
//         cur=cur.next;
//     }
//     if(pre.next==cur){
//         pre=pre.next;
//     }
//     else{
//         pre.next=cur.next;
//     }
//     cur=cur.next;
// }


// Runtime: 4 ms, faster than 80.60% of Go online submissions for Remove Duplicates from Sorted List II.
// Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List II.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummy := ListNode{-1, nil}
    tail := head
    tail2 := &dummy
    for tail != nil {
        t2 := tail.Val
        dup := false
        for tail.Next != nil && tail.Next.Val == t2 {
            dup = true
            tail = tail.Next
        }
        if !dup {
            tail2.Next = tail
            tail2 = tail
        }
        tail = tail.Next
    }
    tail2.Next = nil

    return dummy.Next




    // while(curr != null && curr.next != null){
    //     if(curr.val == curr.next.val){
    //         curr.next = curr.next.next;          // 如果和 后面相同, 则直接指向 后后面.
    //     }
    //     else curr = curr.next;
    // }

    // Runtime: 4 ms, faster than 82.28% of Go online submissions for Remove Duplicates from Sorted List.
    // Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
    // ... LT0083...
    // ... distinct, 不是 remove all duplicate
    // if head == nil || head.Next == nil {
    //     return head
    // }
    // dummy := ListNode{-1, head}
    // tail1, tail2 := head.Next, head             // original, new
    // for tail1 != nil {
    //     if tail1.Val != tail2.Val {
    //         tail2.Next = tail1
    //         tail2 = tail1
    //     }
    //     tail1 = tail1.Next
    // }
    // tail2.Next = nil
    // return dummy.Next
}

func main_LT0082_20211026() {
// func main() {

    fmt.Println("ans:")


}
