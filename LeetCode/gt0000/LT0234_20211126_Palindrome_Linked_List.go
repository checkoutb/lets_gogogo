// package sdq
package main

import (
    "fmt"
)

// class Solution {
//     public:
//         ListNode* temp;
//         bool isPalindrome(ListNode* head) {
//             temp = head;
//             return check(head);
//         }
//
//         bool check(ListNode* p) {
//             if (NULL == p) return true;
//             bool isPal = check(p->next) & (temp->val == p->val);
//             temp = temp->next;
//             return isPal;
//         }
//     };
// temp保存的是头，  dfs到最后一个， 最后一个 和 头对比。
// temp 保存 第二个， dfs出栈一次，是倒数第二个，  倒数第二个 和 第二个对比。
// temp 保存第三个， dfs出栈，是倒数第三个。



// def isPalindrome(self, head):
//     rev = None
//     slow = fast = head
//     while fast and fast.next:
//         fast = fast.next.next
//         rev, rev.next, slow = slow, rev, slow.next
//     if fast:
//         slow = slow.next
//     while rev and rev.val == slow.val:
//         slow = slow.next
//         rev = rev.next
//     return not rev



// for fast != nil && fast.Next != nil {
//     fast = fast.Next.Next
//     slow = slow.Next
// }
// var prev *ListNode
// for slow != nil {
//     tmp := slow.Next
//     slow.Next = prev
//     prev = slow
//     slow = tmp
// }
// left, right := head, prev
// 反转后半部分。



// .. 这个时间是真的不稳定啊。。
// Runtime: 144 ms, faster than 95.60% of Go online submissions for Palindrome Linked List.
// Memory Usage: 11.2 MB, less than 43.52% of Go online submissions for Palindrome Linked List.
// Runtime: 268 ms, faster than 16.04% of Go online submissions for Palindrome Linked List.
// Memory Usage: 8.9 MB, less than 97.58% of Go online submissions for Palindrome Linked List.
// Runtime: 140 ms, faster than 99.12% of Go online submissions for Palindrome Linked List.
// Memory Usage: 10 MB, less than 80.66% of Go online submissions for Palindrome Linked List.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 应该是 快慢指针， 并且 慢的那个指针一路反转， 找到中点后， 开始对比。  而且计数下，知道是长度是 奇数还是偶数，也不太需要，看 fast.Next==nil 就可以判断了
// .. 代码顺序 来决定 中间的结点是否需要。
func isPalindrome(head *ListNode) bool {
    f, s := head, head
    var hd *ListNode = nil
    for f != nil {
        f = f.Next
        if f == nil {          // 奇数长度
            s = s.Next
            break
        }
        f = f.Next
        t2 := s.Next
        s.Next = hd
        hd = s
        s = t2
    }
    for hd != nil {
        if hd.Val != s.Val {
            return false
        }
        hd, s = hd.Next, s.Next
    }
    return true
}


func main_LT0234_20211126() {
// func main() {

    fmt.Println("ans:")


}
