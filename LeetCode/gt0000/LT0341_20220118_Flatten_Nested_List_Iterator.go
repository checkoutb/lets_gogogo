// package sdq
package main

import (
    "fmt"
)














// 并不需要 下标，获得 list后，全部入队/stack就可以了。

// Runtime: 7 ms, faster than 40.35% of Go online submissions for Flatten Nested List Iterator.
// Memory Usage: 5.8 MB, less than 40.35% of Go online submissions for Flatten Nested List Iterator.
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
// 构造的时候 直接遍历完。
// 维护一个 list。 就是 dfs 的 栈。.....数组，需要 int 来维护下标，每层都需要。。.....gg
// .给的是数组。
type NestedIterator struct {
    arr []int
    idx int
}

func dfsa341(ni *NestedInteger, arr *[]int) {
    if ni.IsInteger() {
        *arr = append(*arr, ni.GetInteger())
        return
    }
    for _, v := range ni.GetList() {
        dfsa341(v, arr)
    }
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    arr := []int{}
    for _, v := range nestedList {
        dfsa341(v, &arr)
    }
    return &NestedIterator{arr, 0}
}

func (this *NestedIterator) Next() int {
    this.idx++
    return this.arr[this.idx - 1]
}

func (this *NestedIterator) HasNext() bool {
    return this.idx < len(this.arr)
}

func main_LT0341_20220118() {
// func main() {

    fmt.Println("ans:")


}
