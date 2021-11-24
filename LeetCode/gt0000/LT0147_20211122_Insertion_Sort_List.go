// package sdq
package main

import (
    "fmt"
)


// sorted, cur := head, head.Next
// for cur != nil {
//     if sorted.Val <= cur.Val {
//         sorted = sorted.Next         // 已有序。 sorted 后移一位。
//     } else {
//         pre := dummy
//         for pre.Next.Val < cur.Val {
//             pre = pre.Next
//         } 
//         sorted.Next = cur.Next       // cur无序的情况下， 到这行的时候，已经知道 cur的位置了， sorted依然还是 有序list的尾巴，需要保存 下一个结点，让 cur=sorted.Next来使用
//         cur.Next = pre.Next
//         pre.Next = cur
//     }
//     cur = sorted.Next
// }
// 原地操作。 主要是指：下一个需要被比较的结点 就是 已排序的list的 尾巴的  Next。 就是不需要额外的指针来 指明下一个被比较的结点
// sorted指向了 已有序的 list的 尾巴， cur就是 尾巴的后面一个。





// Runtime: 28 ms, faster than 60.42% of Go online submissions for Insertion Sort List.
// Memory Usage: 3.5 MB, less than 27.08% of Go online submissions for Insertion Sort List.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    
    dummy := &ListNode{Val: -10000, Next: nil}
    for head != nil {
        // if t2.Val < head.Val { ; } else {t2 = dummy}
        t2 := dummy
        for t2.Next != nil && t2.Next.Val < head.Val {
            t2 = t2.Next
        }
        headNext := head.Next
        // if t2.Next != nil {
        //     head.Next = t2.Next
        // } else {
        //     head.Next = nil
        // }
        head.Next = t2.Next
        t2.Next = head
        head = headNext
    }
    return dummy.Next
}

func main_LT0147_20211122() {
// func main() {

    fmt.Println("ans:")


}
