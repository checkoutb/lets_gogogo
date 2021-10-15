// package sdq
package main

import (
    "fmt"
    // "./utils/utils"      // 真的不知道怎么设计的这种东西。
)


// ListNode** t1 = &head, *t2 = head;
// c的，不需要dummy。


// def index(node):
// if not node:
//     return 0
// i = index(node.next) + 1
// if i > n:
//     node.next.val = node.val
// return i
// index(head)
// return head.next
// 。。。dfs到尾巴，然后返回1, 每次+1, 所以 从尾巴到头，是1,2,3,4.....
// 这里是 next.VAL = VAL.... 等于就是 移除 倒数第n个，然后把 前面的 向后移动一位。
// 这样 head.next 就是 原先的head的值。

// 这样的话，dfs，然后 i==n 删除一个节点。 也可以。 不过需要 dummy， 或者 ListNode**。  后者需要 手动delete。


// def remove(head):
// if not head:
//     return 0, head
// i, head.next = remove(head.next)
// return i+1, (head, head.next)[i+1 == n]
// return remove(head)[1]
// dfs， 从尾巴向头 进行编号，并且返回自己这个Node， 如果 i+1==n 那么就 返回 head.next。 等于跳过了 head 这个结点。
// 重新创建了 Node的 关系。并且如果是 第n个， 就 用 Next 代替自己， 使得 前驱 直接 指向 自己的后继， 这样就跳过了一个节点。


// fast = slow = head
// for _ in range(n):
//     fast = fast.next
// if not fast:
//     return head.next
// while fast.next:
//     fast = fast.next
//     slow = slow.next
// slow.next = slow.next.next
// return head
// fast为空，说明删除 第一个，就直接返回了 head.next




// 删链表头。。。
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.4 MB, less than 43.56% of Go online submissions for Remove Nth Node From End of List.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    p1, p2 := head, head
    dummy := &ListNode{-1, head}
    p2 = dummy
    for n > 0 {             // no   n-- > 0
        p1 = (*p1).Next
        n--
    }
    for p1 != nil {
        p1 = p1.Next
        p2 = (*p2).Next
    }
    if p2.Next == nil {         // not need.
        p2.Next = nil
    } else {
        p2.Next = p2.Next.Next
    }
    return dummy.Next
}


func main_LT0019_20211013() {
// func main() {

    fmt.Println("ans:")


}
