// package sdq
package main

import (
    "fmt"
    // "container/heap"
    "math"
)





// private Stack<int[]> stack = new Stack<>();
// public MinStack() { }
// public void push(int x) {
//     if (stack.isEmpty()) {
//         stack.push(new int[]{x, x});
//         return;
//     }
//     int currentMin = stack.peek()[1];
//     stack.push(new int[]{x, Math.min(x, currentMin)});
// }
// .. stack 的顺序， min只会缩小，不会增大。。。


// func (this *MinStack) Push(val int)  {
//     this.s = append(this.s, val)
//     if len(this.m) == 0 || val <= this.m[len(this.m)-1] {
//         this.m = append(this.m, val)
//     } else {
//         this.m = append(this.m,  this.m[len(this.m)-1])
//     }
// }
// 2个数组代替 stack<pair<int,int>>


// int min = Integer.MAX_VALUE;
// Stack<Integer> stack = new Stack<Integer>();
// public void push(int x) {
//     // only push the old minimum value when the current 
//     // minimum value changes after pushing the new value x
//     if(x <= min){          
//         stack.push(min);
//         min=x;
//     }
//     stack.push(x);
// }
// public void pop() {
//     // if pop operation could result in the changing of the current minimum value, 
//     // pop twice and change the current minimum value to the last minimum value.
//     if(stack.pop() == min) min=stack.pop();
// }
// 当 min 发生变化时， 把原min push进去， 
// 3,2,1,1,1,1
// 3,3,2,2,1,1,1,1  . 不，是 <===== min的时候 就 push。
// 3,2,2,2,1,1,1,1,1,1,1。。



// Runtime: 16 ms, faster than 97.56% of Go online submissions for Min Stack.
// Memory Usage: 8.6 MB, less than 46.58% of Go online submissions for Min Stack.
type MinStack struct {
    Map2 map[int]int
    Arr []int
    Mn int
}
func Constructor() MinStack {
    return MinStack{map[int]int{}, []int{}, math.MinInt32}
}
func (this *MinStack) Push(val int)  {
    this.Arr = append(this.Arr, val)
    this.Map2[val]++
    if val < this.Mn {
        this.Mn = val
    }
}
func (this *MinStack) Pop()  {
    this.Map2[this.Arr[len(this.Arr) - 1]]--
    this.Arr = this.Arr[0 : len(this.Arr) - 1]
}
func (this *MinStack) Top() int {
    return this.Arr[len(this.Arr) - 1]
}
func (this *MinStack) GetMin() int {
    if this.Map2[this.Mn] != 0 {
        return this.Mn
    } else {
        this.Mn = math.MaxInt32
        for k, v := range this.Map2 {
            if v != 0 && k < this.Mn {
                this.Mn = k
            }
        }
    }
    return this.Mn
}




// // An IntHeap is a min-heap of ints.
// type IntHeap []int
// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *IntHeap) Push(x interface{}) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }
// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
// // // This example inserts several ints into an IntHeap, checks the minimum,
// // // and removes them in order of priority.
// // func main() {
// // 	h := &IntHeap{2, 1, 5}
// // 	heap.Init(h)
// // 	heap.Push(h, 3)
// // 	fmt.Printf("minimum: %d\n", (*h)[0])
// // 	for h.Len() > 0 {
// // 		fmt.Printf("%d ", heap.Pop(h))
// // 	}
// // }

// /**
//  * Your MinStack object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Push(val);
//  * obj.Pop();
//  * param_3 := obj.Top();
//  * param_4 := obj.GetMin();
//  */
// // stack + ordered_map/priority_queue/min-heap
// // Node {val, valid, stack.Next, num.Prev, num.Next}
// type MinStack struct {
//     Mnh *IntHeap
//     Stk []int
// }
// func Constructor() MinStack {
//     mnh := &IntHeap{}
//     stk := []int{}
//     return MinStack{mnh, stk}
// }
// func (this *MinStack) Push(val int)  {
//     this.Stk = append(this.Stk, val)
//     heap.Push(this.Mnh, val)
// }
// func (this *MinStack) Pop()  {
//     heap.Remove(this.Mnh, this.Stk[len(this.Stk) - 1])          // 这个是下标。。。不是值。。
//     this.Stk = this.Stk[0 : len(this.Stk) - 1]
// }
// func (this *MinStack) Top() int {
//     return this.Stk[len(this.Stk) - 1]
// }
// func (this *MinStack) GetMin() int {
//     return (*(this.Mnh))[0]
//     // t2 := heap.Pop(this.Mnh)
//     // fmt.Println(t2, "... ")
//     // heap.Push(this.Mnh, t2)
//     // return t2
// }



func main_LT0155_20211123() {
// func main() {

    fmt.Println("ans:")

    obj := Constructor()
    obj.Push(1)
    obj.Push(2)
    fmt.Println(obj.Top())
    fmt.Println(obj.GetMin())
    obj.Pop()
    fmt.Println(obj.GetMin())

}
