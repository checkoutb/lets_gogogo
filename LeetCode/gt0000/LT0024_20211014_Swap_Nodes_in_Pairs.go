// package sdq
package main

import (
    "fmt"
)



// if head == nil || head.Next == nil {
//     return head
// }
// newHead := head.Next
// head.Next = swapPairs(head.Next.Next)
// newHead.Next = head
// return newHead





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
// Memory Usage: 2.2 MB, less than 39.37% of Go online submissions for Swap Nodes in Pairs.
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head};
    np := dummy
    for np != nil && np.Next != nil && np.Next.Next != nil {
        t2 := np.Next.Next
        np.Next.Next = t2.Next
        t2.Next = np.Next
        np.Next = t2
        np = np.Next.Next
    }
    return (*dummy).Next
}

func main_LT0024_20211014() {
// func main() {

    fmt.Println("ans:")


}
