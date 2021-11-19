// package sdq
package main

import (
    "fmt"
)



//   // loop 1. copy all the nodes
//   RandomListNode node = head;
//   while (node != null) {
//     map.put(node, new RandomListNode(node.label));
//     node = node.next;
//   }
//   // loop 2. assign next and random pointers
//   node = head;
//   while (node != null) {
//     map.get(node).next = map.get(node.next);
//     map.get(node).random = map.get(node.random);
//     node = node.next;
//   }


//   // First round: make copy of each node,
//   // and link them together side-by-side in a single list.
//   while (iter != null) {
//     next = iter.next;
//     RandomListNode copy = new RandomListNode(iter.label);
//     iter.next = copy;                // ...
//     copy.next = next;                // ... 直接扩写。把 新的结点放在老的结点后面
//     iter = next;
//   }
//   // Second round: assign random pointers for the copy nodes.
//   iter = head;
//   while (iter != null) {
//     if (iter.random != null) {
//       iter.next.random = iter.random.next;       // iter是原来结点，iter.next是新结点，所以这里就是 新结点.random = 老结点.ramdon.新结点
//     }
//     iter = iter.next.next;
//   }
//   // Third round: restore the original list, and extract the copy list.
//   iter = head;
//   RandomListNode pseudoHead = new RandomListNode(0);
//   RandomListNode copy, copyIter = pseudoHead;
//   while (iter != null) {
//     next = iter.next.next;
//     // extract the copy
//     copy = iter.next;
//     copyIter.next = copy;
//     copyIter = copy;
//     // restore the original list
//     iter.next = next;
//     iter = next;
//   }
// 



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Copy List with Random Pointer.
// Memory Usage: 3.6 MB, less than 79.46% of Go online submissions for Copy List with Random Pointer.
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// 先按Next来一遍，并且map下，然后再处理 Random
func copyRandomList(head *Node) *Node {
    map2 := map[*Node]*Node{}         // old node -> new node
    dummy := Node{-1, nil, nil}
    t2 := head
    t3 := &dummy
    for t2 != nil {
        n2 := &Node{t2.Val, nil, nil}
        t3.Next = n2
        map2[t2] = n2
        t2 = t2.Next
        t3 = t3.Next
    }
    t2, t3 = head, dummy.Next
    for t2 != nil {
        if t2.Random != nil {
            t3.Random = map2[t2.Random]
        }
        t2, t3 = t2.Next, t3.Next
    }
    return dummy.Next
}


func main_LT0138_20211119() {
// func main() {

    fmt.Println("ans:")


}
