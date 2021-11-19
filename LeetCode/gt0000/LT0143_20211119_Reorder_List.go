// package sdq
package main

import (
    "fmt"
    "leetcode/utils"
)




// for second != nil {
//     tmp := second.Next
//     second.Next = rev
//     rev = second
//     second = tmp
// }
// first, second := head, rev
// for second != nil {
//     tmp1, tmp2 := first.Next, second.Next
//     first.Next = second
//     second.Next = tmp1
//     first, second = tmp1, tmp2
// }



// 转成数组。。。


// while(curr!=null){
//     temp = curr.next;
//     curr.next = prev;
//     prev = curr;
//     curr = temp;
// }
// 原地反转List


// 离谱。。。
// Runtime: 12 ms, faster than 39.73% of Go online submissions for Reorder List.
// Memory Usage: 6.6 MB, less than 15.07% of Go online submissions for Reorder List.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 感觉只能暴力啊。或者reverse下，然后 一边取一个。
// 就是 找到中点， 切分， 后半段reverse。 然后 一边取一个。
// 快慢指针 找中点。
func reorderList(head *utils.ListNode) {
    f, s := head, head
    var pres *utils.ListNode = nil
    for f != nil {
        f = f.Next
        if f != nil {
            f = f.Next
        }
        pres = s
        s = s.Next
        if f == s {
            break
        }
    }
    // f = s.Next
    f = s
    // s.Next = nil
    s = head
    if pres != nil {
        pres.Next = nil
    }
    if f == nil {
        return
    }
    _, f = reverse143(f)
    
    // utils.ShowListNode(s)
    // utils.ShowListNode(f)

    t2 := head
    s = s.Next
    isF := true
    for s != nil || f != nil {
        if isF {
            t2.Next = f
            t2 = t2.Next
            f = f.Next
        } else {
            t2.Next = s
            t2 = t2.Next
            s = s.Next
        }
        isF = !isF
    }

    // asd := 1

    // for f != nil {
    //     if asd > 5 {
    //         break
    //     }
    //     asd++
    //     // fmt.Println("z")
    //     fmt.Println(s, s.Next)
    //     fmt.Println(f, f.Next)
    //     t2 := s.Next        // 2
    //     s.Next = f      // 1->4
    //     t2.Next = f.Next        // 2->3
    //     f.Next = t2     // 4->2
    //     s = t2
    //     f = f.Next.Next
    //     fmt.Println(s, s.Next)
    //     fmt.Println(f, f.Next)
    // }
    // if f != nil {
    //     s.Next = f
    // }
}

func reverse143(node *utils.ListNode) (*utils.ListNode, *utils.ListNode) {        // prev, head
    if node.Next == nil {
        return node, node
    }
    prev, head := reverse143(node.Next)
    prev.Next = node
    node.Next = nil
    return node, head
}

// func main_LT0143_20211119() {
func main() {

    fmt.Println("ans:")

    arr := []int{1}

    head := utils.ConvertToListNode(arr)
    utils.ShowListNode(head)

    reorderList(head)

    utils.ShowListNode(head)
}
