// package sdq
package main

import (
    "container/heap"
    "fmt"
)



// class Solution {
//     public:
//         using T=tuple<int,int,int>;
//
// priority_queue<T, vector<T>, greater<T>> pq;
// 
// T t=pq.top(); pq.pop();
// res.push_back({nums1[get<1>(t)], nums2[get<2>(t)]});
// if(++get<2>(t) < nums2.size())
// {
//     get<0>(t)=nums1[get<1>(t)] + nums2[get<2>(t)];
//     pq.push(move(t));
// }



// type pair []int
// func (p pair) sum() int {
// 	return p[0]+p[1]
// }
// type mHeap []pair
// func (h mHeap) Len() int {return len(h)}
// func (h mHeap) Less(i, j int) bool {return h[i].sum() > h[j].sum()}
// func (h mHeap) Swap(i, j int) {h[i],h[j] = h[j],h[i]}
// func (h *mHeap) Push(x interface{}) {*h = append(*h, x.(pair))}
// func (h *mHeap) Pop() interface{} {
// 	l := len(*h)
// 	v := (*h)[l-1]
// 	*h = (*h)[:l-1]
// 	return v
// }



// PriorityQueue<int[]> que = new PriorityQueue<>((a,b)->a[0]+a[1]-b[0]-b[1]);




// Runtime: 131 ms, faster than 61.36% of Go online submissions for Find K Pairs with Smallest Sums.
// Memory Usage: 8.9 MB, less than 100.00% of Go online submissions for Find K Pairs with Smallest Sums.
// 需要一个 pri_q。。。 <int,int,int> <和，下标，下标> 。 初始 00下标，下一个 可能是 01 也可能是 10, 都入队。
// 00 之后 01 或 10， 如果是 10, 那么 10之后 是 11或 20  但是 11 必然大于 01 的， 所以是 01 或 20  。但是怎么剔除11？。  还需要去重。 01->11  10->11 ,11重复了。
// 剔除和去重 用 map。  一个map记录 已经加入到 queue的 下标1×10000+下标2 。  每个数组各一个map，<下标，另一个数组的最大下标>， 有了 后者，似乎不需要前者。 后者就去重了。
// 但是少个 最主要的 pri_q......
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    ans := [][]int{}
    sz1, sz2 := len(nums1), len(nums2)

    pq := PriorityQueue{}

	heap.Init(&pq)

    heap.Push(&pq, &Item{ idx1: 0, idx2: 0, priority: nums1[0] + nums2[0] })
    map2 := map[int]bool{}
    for i := 0; i < k; i++ {
        if len(pq) == 0 {
            break
        }
        item := heap.Pop(&pq).(*Item)
        t1, t2 := item.idx1, item.idx2
        ans = append(ans, []int{ nums1[t1], nums2[t2] })
        if t1 + 1 < sz1 {
            t3 := (t1 + 1) * 10000 + t2
            if _, ok := map2[t3]; !ok {
                map2[t3] = true
                heap.Push(&pq, &Item{ idx1: t1 + 1, idx2 : t2, priority: nums1[t1 + 1] + nums2[t2] })
            }
        }
        if t2 + 1 < sz2 {
            t4 := t1 * 10000 + t2 + 1
            if _, ok := map2[t4]; !ok {
                map2[t4] = true
                heap.Push(&pq, &Item{ idx1: t1, idx2: t2 + 1, priority: nums1[t1] + nums2[t2 + 1] })
            }
        }
    }

    return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
    idx1 int
    idx2 int
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}




// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main1123() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}

func main_LT0373_20220121() {
// func main() {

    fmt.Println("ans:")

    // arr1 := []int{1,7,11}
    // arr2 := []int{2,4,6}
    // k := 3

    // arr1 := []int{1,1,2}
    // arr2 := []int{1,2,3}
    // k := 2

    arr1 := []int{1,2}
    arr2 := []int{3}
    k := 3

    ans := kSmallestPairs(arr1, arr2, k)

    fmt.Println(ans)

}
