// package sdq
package main

import (
    "fmt"
    "container/heap"
)



// private Queue<Long> small = new PriorityQueue(),
//                     large = new PriorityQueue();
// public void addNum(int num) {
//     large.add((long) num);
//     small.add(-large.poll());
//     if (large.size() < small.size())
//         large.add(-small.poll());
// }
// public double findMedian() {
//     return large.size() > small.size()
//                     ? large.peek()
//                     : (large.peek() - small.peek()) / 2.0;
// }


// Queue<Integer> q = new PriorityQueue(), z = q, t, Q = new PriorityQueue(Collections.reverseOrder()); 
// public void addNum(int num) {
//     (t=Q).add(num);
//     (Q=q).add((q=t).poll());
// }
// public double findMedian() {
//     return (Q.peek() + z.peek()) / 2.;
// }


// Queue[] q = {new PriorityQueue(), new PriorityQueue(Collections.reverseOrder())};
// int i = 0;
// public void addNum(int num) {
//     q[i].add(num);
//     q[i^=1].add(q[i^1].poll());
// }
// public double findMedian() {
//     return ((int)(q[1].peek()) + (int)(q[i].peek())) / 2.0;
// }




// 太繁琐了。。。吐了。
// Runtime: 364 ms, faster than 66.91% of Go online submissions for Find Median from Data Stream.
// Memory Usage: 22 MB, less than 47.79% of Go online submissions for Find Median from Data Stream.
// max-heap, min-heap  但是 go。。。 有min-heap， 所以 之前 就不需要什么 pri_queue， 直接 heap。。。 堆 就是 最小堆？ 。 不不不。。。
type MedianFinder struct {
    MnHeap1 *IntHeap     // 存放 取负 后的数，  前半部分   这样就是 前半部分的 最大值
    MnHeap2 *IntHeap      // 后半部分的 最小值
}
func Constructor() MedianFinder {
    h1 := &IntHeap{}
    h2 := &IntHeap{}
    heap.Init(h1)
    heap.Init(h2)
    return MedianFinder{ h1, h2 }
}
func (this *MedianFinder) AddNum(num int)  {
    if this.MnHeap1.Len() == 0 {
        heap.Push(this.MnHeap1, -num)
        return
    } else if this.MnHeap2.Len() == 0 {
        t2 := -heap.Pop(this.MnHeap1).(int)
        if t2 > num {
            heap.Push(this.MnHeap1, -num)
            heap.Push(this.MnHeap2, t2)
        } else {
            heap.Push(this.MnHeap1, -t2)
            heap.Push(this.MnHeap2, num)
        }
        return
    }
    hp := *this.MnHeap2
    if hp[0] <= num {             // .
        heap.Push(this.MnHeap2, num)
        if this.MnHeap2.Len() > this.MnHeap1.Len() {
            t2 := heap.Pop(this.MnHeap2).(int)
            heap.Push(this.MnHeap1, -t2)
        }
    } else {
        heap.Push(this.MnHeap1, -num)
        if this.MnHeap2.Len() < this.MnHeap1.Len() - 1 {
            t2 := heap.Pop(this.MnHeap1).(int)
            heap.Push(this.MnHeap2, -t2)
        }
    }
    // fmt.Println("----- ", this.MnHeap1.Len(), this.MnHeap2.Len(), " - ", *this.MnHeap1, *this.MnHeap2)
}
func (this *MedianFinder) FindMedian() float64 {
    if this.MnHeap1.Len() > this.MnHeap2.Len() {
        // ans := (*this.MnHeap1)[this.MnHeap1.Len() - 1]
        ans := (*this.MnHeap1)[0]
        // fmt.Println(". .. ", ans, this.MnHeap1.Len(), this.MnHeap2.Len(), " - ", *this.MnHeap1, *this.MnHeap2)
        return float64(-ans)
    } else {
        ans := (*this.MnHeap2)[0] - (*this.MnHeap1)[0]
        return float64(ans) / 2
    }
}

// copy
type IntHeap []int  // 定义一个类型
func (h IntHeap) Len() int { return len(h) }  // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool {  // 绑定less方法
	return h[i] < h[j]  // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) {  // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Pop() interface{} {  // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]                   // 被这个耍了。。 (this.MnHeap1)[this.MnHeap1.Len() - 1] ....
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {  // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}






/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */


func main_LT0295_20211204() {
// func main() {

    fmt.Println("ans:")

    // hp := heap{}

    // mf := MedianFinder{}         /....... 
    mf := Constructor()
    mf.AddNum(1)
    mf.AddNum(2)
    fmt.Println(mf.FindMedian())
    mf.AddNum(3)
    mf.AddNum(4)
    mf.AddNum(5)
    mf.AddNum(6)
    fmt.Println(mf.FindMedian())

    // t2 := IntHeap{}
    // heap.Init(&t2)
    // heap.Push(&t2, 2)
    // // heap.Push(t2, 5)
    // // heap.Push(t2, 4)
    // fmt.Println(heap.Pop(&t2))
    // // fmt.Println(heap.Pop(t2))
    // // fmt.Println(heap.Pop(t2))

}
