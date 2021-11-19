// package sdq
package main

import (
    "fmt"
)



// while (!stack.isEmpty()) {
//     Node top = stack.pop();
//     for (Node neighbor : top.neighbors) {
//         if (!visited.containsKey(neighbor)) {
//             stack.push(neighbor);
//             visited.put(neighbor, new Node(neighbor.val));
//         }
//         visited.get(top).neighbors.add(visited.get(neighbor));
//     }
// }


// func helper(node *Node, visited map[*Node]*Node) *Node{
//     if visited, ok := visited[node]; ok {
//         return visited
//     }
//     root := &Node{
//         Val: node.Val,
//         Neighbors: make([]*Node, 0),
//     }




// Runtime: 0 ms, faster than 100.00% of Go online submissions for Clone Graph.
// Memory Usage: 2.9 MB, less than 35.05% of Go online submissions for Clone Graph.
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }
    map2 := map[int]*Node{}
    dfsa133(node, map2);
    return map2[node.Val]
}

func dfsa133(node *Node, map2 map[int]*Node) *Node {
    if _, ok := map2[node.Val]; ok {
        return map2[node.Val]
    }
    t2 := &Node{node.Val, nil}
    map2[node.Val] = t2
    if node.Neighbors != nil {
        t2.Neighbors = []*Node{}
        for _, val := range node.Neighbors {
            t2.Neighbors = append(t2.Neighbors, dfsa133(val, map2))
        }
    }
    return t2
}

func main_LT0133_20211118() {
// func main() {

    fmt.Println("ans:")


}
