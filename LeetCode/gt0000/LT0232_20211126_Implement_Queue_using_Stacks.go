// package sdq
package main

import (
    "fmt"
)


// discuss
// 用于 读写分离， 多线程 操作一个queue时，需要加锁，这时 读写都 加锁。
// 如果 分开，那么 读锁，不影响 写锁，   不过 读的stack为空的时候， 依然要 对写stack加锁 ，来 复制数据吧。
// 但是感觉 操作一个queue， 也可以 读加一个锁，  写加一个锁 啊。 是的。 juc.LinkedBlockingQueue 里有2个 ReentrantLock： takeLock, putLock.





// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Queue using Stacks.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Implement Queue using Stacks.
// 2个stack 完成一个 queue
// 每次都要 把stack 翻一个个啊。
// O(1)
// 要对比 2个stk的长度 对对对，这样就 差不多是O(1)
type MyQueue struct {
    Stk1 []int
    Stk2 []int
}
func Constructor() MyQueue {
    return MyQueue{ []int{}, []int{} }
}
func (this *MyQueue) Push(x int)  {
    this.Stk1 = append(this.Stk1, x)
}
func (this *MyQueue) Pop() int {
    ans := this.Peek()
    this.Stk2 = this.Stk2[0 : len(this.Stk2) - 1]
    return ans
}
func (this *MyQueue) Peek() int {
    if len(this.Stk2) == 0 {
        this.fill()
    }
    return this.Stk2[len(this.Stk2) - 1]
}
func (this *MyQueue) Empty() bool {
    return len(this.Stk1) == 0 && len(this.Stk2) == 0
}
func (this *MyQueue) fill() {
    for len(this.Stk1) > 0 {
        sz1 := len(this.Stk1)
        this.Stk2 = append(this.Stk2, this.Stk1[sz1 - 1])
        this.Stk1 = this.Stk1[0 : sz1 - 1]
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */


func main_LT0232_20211126() {
// func main() {

    fmt.Println("ans:")


}
