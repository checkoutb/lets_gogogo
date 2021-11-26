// package sdq
package main

import (
    "fmt"
)


// func deleteNode(node *ListNode) {
//     node.Val = node.Next.Val
//     node.Next = node.Next.Next
// }
// 。。。。。。。。。。。。。。。。。。。。。。。。。


// func deleteNode(node *ListNode) {
//     *node = *node.Next
// }
// 。。。。。。。。。。。暴击了。。。


// Runtime: 5 ms, faster than 17.35% of Go online submissions for Delete Node in a Linked List.
// Memory Usage: 3.2 MB, less than 9.71% of Go online submissions for Delete Node in a Linked List.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// ？？？？？ 直接给需要删除的node， 不给head。。怎么找到 删除node 的 prev ？  还是单链表啊
// 。。。 移动数字，而不是 删除结点。。 可以。。 脑筋急转弯。。  确实是删除一个结点
func deleteNode(node *ListNode) {
    for node.Next.Next != nil {     // 不是 tail 。
        node.Val = node.Next.Val
        node = node.Next
    }
    node.Val = node.Next.Val
    node.Next = nil
}

func main_LT0237_20211126() {
// func main() {

    fmt.Println("ans:")


}
