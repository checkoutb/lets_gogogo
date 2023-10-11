// package sdq
package main

import (
    "fmt"
    "math"
)

// D D

//   func gcd(a int, b int) int {
//     if a % b == 0 {
//       return b
//     } else {
//       return gcd(b, a % b)
//     }
//   }
  
//   func gcd2(a, b int) int {
//       for b != 0 {
//           a, b = b, a%b  
//       }
//       return a
//   }


// 。。 所以 a > b 是必须的吗？


// Runtime8 ms
// Beats
// 92.9%
// Memory6.8 MB
// Beats
// 7.44%
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// gcd
func insertGreatestCommonDivisors(head *ListNode) *ListNode {

    var pt *ListNode = head;

    for pt.Next != nil {
                            // pt.Val is int, Max need float64... cannot auto cast...
        // var t2 = &ListNode{gcda1(int(math.Max(pt.Val, pt.Next.Val)), int(math.Min(pt.Val, pt.Next.Val))), pt.Next};
        var t2 = &ListNode{gcda1(pt.Val, pt.Next.Val), pt.Next};
        pt.Next = t2;
        pt = pt.Next.Next;
    }
    return head;
}

func gcda1(a, b int) int {
    if a < b {
        a, b = b, a
    }
    if b == 0 {
        return a;
    } else {
        return gcda1(b, a % b);
    }
}

func main_LT2807_20231010() {
// func main() {

    fmt.Println("ans:")


}
