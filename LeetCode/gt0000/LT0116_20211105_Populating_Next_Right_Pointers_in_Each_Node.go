// package sdq
package main

import (
    "fmt"
    "container/list"
)





// root.left.next = root.right;
// if( root.next != null ){                 // ... 上一层的时候设置了next，所以 这里能知道，然后。。
//     root.right.next = root.next.left;
// }
// connect(root.left);
// connect(root.right);



// Runtime: 8 ms, faster than 32.66% of Go online submissions for Populating Next Right Pointers in Each Node.
// Memory Usage: 6.8 MB, less than 22.07% of Go online submissions for Populating Next Right Pointers in Each Node.
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// bfs
func connect(root *Node) *Node {
    if root == nil {
        return root
    }
	que := list.New()
    que.PushBack(root)
    for que.Len() > 0 {
        // sz1 := que.Len() - 1
        // prev := que.Remove(que.Front())
        // if prev.Left != nil {
        //     que.PushBack(prev.Left)
        // }
        // if prev.Right != nil {
        //     que.PushBack(prev.Left)
        // }
        sz1 := que.Len()
        var prev *Node = nil
        for sz1 > 0 {
            sz1--
            node := que.Remove(que.Front()).(*Node)
            if prev != nil {
                prev.Next = node
            }
            if node.Left != nil {
                que.PushBack(node.Left)
            }
            if node.Right != nil {
                que.PushBack(node.Right)
            }
            prev = node
        }
    }
    return root
}


func main_LT0116_20211105() {
// func main() {

    fmt.Println("ans:")


}
