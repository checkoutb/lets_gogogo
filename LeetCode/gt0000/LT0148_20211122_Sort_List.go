// package sdq
package main

import (
    "fmt"
)

// ....  merge sort 快。。。  快慢指针 。。。 多遍历一次， 所以应该是 2nlogn 等于 nlogn


// sort， merge  2部分。

// 步长增长。 和shell排序的 缩小增量  反过来。
// 步长×2 以后 可以merge。


// Runtime: 670 ms, faster than 17.88% of Go online submissions for Sort List.
// Memory Usage: 7.6 MB, less than 50.84% of Go online submissions for Sort List.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// merge sort ? insert的话，用 上一题discuss里的 也差不多是nlogn，不，那个方差有点大，最好情况O(n)，最差O(n^2)，还是O(n^2)
func sortList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    dummy := &ListNode{-11111, head}
    end, endNext := head, head.Next
    for endNext != nil {
        if end.Val < endNext.Val {
            end = endNext
        } else {
            pre := dummy
            for pre.Next.Val < endNext.Val {
                pre = pre.Next
            }
            end.Next = endNext.Next
            endNext.Next = pre.Next
            pre.Next = endNext
        }
        endNext = end.Next
    }
    return dummy.Next
}

// func mergeSorta147(node *ListNode) *ListNode {
//     if node.Next == nil {
//         return node
//     }
//     mergeSorta147()         // ... divide 长度不知道。。而且divide的时候需要遍历，无法O(1)完成 切分。
// }

func main_LT0148_20211122() {
// func main() {

    fmt.Println("ans:")


}
