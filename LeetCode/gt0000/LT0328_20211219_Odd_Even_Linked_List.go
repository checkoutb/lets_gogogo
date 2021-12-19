// package sdq
package main

import (
    "fmt"
    "leetcode/utils"
)

// while (even != null && even.next != null) {
//     odd.next = odd.next.next; 
//     even.next = even.next.next; 
//     odd = odd.next;
//     even = even.next;
// }



// Runtime: 5 ms, faster than 7.40% of Go online submissions for Odd Even Linked List.
// Memory Usage: 3.4 MB, less than 14.04% of Go online submissions for Odd Even Linked List.
func oddEvenList(head *utils.ListNode) *utils.ListNode {
    odd := &utils.ListNode{-1, nil}
    even := &utils.ListNode{-1, nil}
    ans := odd
    even2 := even
    for head != nil {
        odd.Next = head
        head = head.Next
        odd = odd.Next
        if head != nil {
            even.Next = head
            even = even.Next
            head = head.Next
        }
    }
    even.Next = nil         // 。
    odd.Next = even2.Next
    return ans.Next
}


func main_LT0328_20211219() {
// func main() {

    fmt.Println("ans:")

    head := utils.ConvertToListNode([]int{1,2,3,4,5})

    ans := oddEvenList(head)

    utils.ShowListNode(ans)

}
