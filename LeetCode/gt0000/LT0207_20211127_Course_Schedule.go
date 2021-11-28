// package sdq
package main

import (
    "fmt"
)


// BFS uses the indegrees of each node. We will first try to find a node with 0 indegree. 
// If we fail to do so, there must be a cycle in the graph and we return false. 
// Otherwise we set its indegree to be -1 to prevent from visiting it again and reduce the indegrees of its neighbors by 1. 
// This process will be repeated for n (number of nodes) times.
// 剔除入度为0的课程，然后把这些课程的后置课程的 入度-1。 再剔除入度0的课程。


// Runtime: 17 ms, faster than 32.19% of Go online submissions for Course Schedule.
// Memory Usage: 6.8 MB, less than 17.69% of Go online submissions for Course Schedule.
// topological
func canFinish(numCourses int, prerequisites [][]int) bool {
    map2 := map[int][]int{}
    for _, v := range prerequisites {
        if _, ok := map2[v[0]]; !ok {
            map2[v[0]] = []int{v[1]}
        } else {
            map2[v[0]] = append(map2[v[0]], v[1])
        }
    }
    vst := map[int]bool{}
    for k, _ := range map2 {
        if !dfsa207(map2, vst, k, map[int]bool{}) {
            return false
        }
    }
    return true
}

// can complete crs
func dfsa207(map2 map[int][]int, vst map[int]bool, crs int, tmp map[int]bool) bool {
    if _, ok := map2[crs]; !ok {
        return true
    }
    if _, ok := vst[crs]; ok {
        return true
    }
    if _, ok := tmp[crs]; ok {
        if tmp[crs] {
            return false
        }
    }
    tmp[crs] = true
    for _, preCrs := range map2[crs] {
        if !dfsa207(map2, vst, preCrs, tmp) {
            tmp[crs] = false
            return false
        }
    }
    tmp[crs] = false
    vst[crs] = true
    return true
}


func main_LT0207_20211127() {
// func main() {

    fmt.Println("ans:")


}
