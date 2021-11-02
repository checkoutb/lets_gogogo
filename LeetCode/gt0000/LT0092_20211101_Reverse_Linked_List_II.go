// package sdq
package main

import (
    "fmt"
    // "leetcode/utils"         // ok的.
)




// ListNode **a = &head, **b;
// for (;m--; n--)
//     a = &(*(b=a))->next;
// for (;n--; swap(*b, *a))
//     swap(*b, (*a)->next);
// return head;



// 乱

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.1 MB, less than 45.71% of Go online submissions for Reverse Linked List II.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == right {
        return head
    }
    dummy := ListNode{-1, head}
    p2 := &dummy                // left's prev
    for i := 0; i < left - 1; i++ {
        p2 = p2.Next
    }
    p3 := p2            // right
    pre := p3
    p3 = p3.Next
    // fmt.Println(p2.Val, ", ", p3.Val)
    for i := left; i < right; i++ {
        t2 := p3.Next
        p3.Next = pre
        pre = p3
        p3 = t2
    }
    // fmt.Println(p3.Val, ", ", p2.Next.Val, ", ", p3.Next, ", ", (pre == &dummy))
    // t2 := p3.Next
    p2.Next.Next = p3.Next
    if pre != &dummy {
        p3.Next = pre
    } else {
        p3.Next = nil
    }
    p2.Next = p3

    return dummy.Next
}



// type ListNode struct {
//     Val int
//     Next *ListNode
// }

func main_LT0092_20211101() {
// func main() {

    fmt.Println("ans:")

    // list := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
    // ans := reverseBetween(&list, 2, 4)

    // list := ListNode{11, nil}
    // ans := reverseBetween(&list, 1, 1)

    list := ListNode{3, &ListNode{5, nil}}
    ans := reverseBetween(&list, 1, 1)

    // utils.ShowListNode(ans)              // ... 类型不一样,而且 好像没有办法 直接使用啊.. 类型必须带 包名..但是带了包名 就不是同一个类了..

    for ans != nil {
        fmt.Print(ans.Val, ", ")
        ans = ans.Next
    }

}
