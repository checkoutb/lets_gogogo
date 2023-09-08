// package sdq
package main

import (
    "fmt"
)

// D D

// m[curHead] = &Node{Val: curHead.Val}

// 		p.Next = &Node{p.Val, p.Next, nil}

// 		prevCopy, prevCopy.Next = p.Next, p.Next

// p.Next = &Node{
//     Val:    p.Val,
//     Next:   p.Next,
//     Random: nil,
// }




// Runtime0 ms
// Beats
// 100%
// Memory3.6 MB
// Beats
// 28.38%
type Node struct {
    Val int
    Next *Node
    Random *Node
}
func copyRandomList(head *Node) *Node {
    if head == nil {
        return head;
    }
    // var ans *Node = new Node{head.Val, head.Next};
    var ans *Node = new(Node);
    (*ans).Val = head.Val;
    // ans.Next 

    map2 := map[*Node]*Node{};
    var ptr *Node = head;
    var ptr2 *Node = ans;
    map2[head] = ans;
    // map2[ptr] = ptr2;
    for ptr != nil {
        map2[ptr] = ptr2;
        ptr2.Val = ptr.Val;
        ptr = ptr.Next;
        // ptr2.Next = new Node{ptr.Val};
        if ptr != nil {
            ptr2.Next = new(Node);
            ptr2 = ptr2.Next;
        }
        // ptr2.Val = 
    }

    ptr = head;
    ptr2 = ans;
    for ptr != nil {
        if ptr.Random != nil {
            ptr2.Random = map2[ptr.Random];
        }
        ptr = ptr.Next
        ptr2 = ptr2.Next
    }
    return ans;
}


func main_LT0138_20230905() {
// func main() {

    fmt.Println("ans:")


}
