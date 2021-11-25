// package sdq
package main

import (
    "fmt"
)




// For Example: A = [2,1,3,4,6,3,8,9,10,12,56], w=4
//     partition the array in blocks of size w=4. The last block may have less then w.
//     2, 1, 3, 4 | 6, 3, 8, 9 | 10, 12, 56|
//     Traverse the list from start to end and calculate max_so_far. Reset max after each block boundary (of w elements).
//     left_max[] = 2, 2, 3, 4 | 6, 6, 8, 9 | 10, 12, 56
//     Similarly calculate max in future by traversing from end to start.
//     right_max[] = 4, 4, 4, 4 | 9, 9, 9, 9 | 56, 56, 56
//     now, sliding max at each position i in current window, sliding-max(i) = max{right_max(i), left_max(i+w-1)}
//     sliding_max = 4, 6, 6, 8, 9, 10, 12, 56

    // max_left[0] = in[0];
    // max_right[in.length - 1] = in[in.length - 1];
    // for (int i = 1; i < in.length; i++) {
    //     max_left[i] = (i % w == 0) ? in[i] : Math.max(max_left[i - 1], in[i]);
    //     final int j = in.length - i - 1;
    //     max_right[j] = (j % w == 0) ? in[j] : Math.max(max_right[j + 1], in[j]);
    // }
    // final int[] sliding_max = new int[in.length - w + 1];
    // for (int i = 0, j = 0; i + w <= in.length; i++) {
    //     sliding_max[j++] = Math.max(max_right[i], max_left[i + w - 1]);
    // }
// too hard to understand....sad
// 感觉是  以这个点为最后一个元素 的范围的 最大值  和 以这个点为第一个元素的范围内的最大值。



// int[] lmax = new int[nums.length];
// int[] rmax = new int[nums.length];
// lmax[0] = nums[0];
// rmax[nums.length - 1] = nums[nums.length - 1];
// for(int i = 1, j = nums.length - 2; i < nums.length && j > -1; i++, j--) {
//     lmax[i] = i % k == 0 ? nums[i] : Math.max(nums[i], lmax[i - 1]);
//     rmax[j] = j % k == 0 ? nums[j] : Math.max(nums[j], rmax[j + 1]);
// }
// int[] slidingWindow = new int[nums.length + 1 - k];
// for(int i = 0; i < slidingWindow.length; i++) {
//     slidingWindow[i] = Math.max(rmax[i], lmax[i + k - 1]);
// }



// for(int i=0;i<n;i++){
//     if(!que.isEmpty() && i-k>=0 && que.peekFirst() == i-k){
//         que.pollFirst();
//     }
//     while(!que.isEmpty() && (nums[que.peekLast()] <= nums[i] || que.size() >= k) ){
//         que.pollLast();
//     }
//     que.addLast(i);
//     if(i-k + 1>=0){
//         ret[idx++] = nums[que.peekFirst()];         
//     }
// }


// queue := make([]int, 0)
// result := make([]int, 0, len(nums) - k + 1)
// for i, num := range nums {
//     for len(queue) > 0 && nums[queue[len(queue) - 1]] <= num {
//         queue = queue[:len(queue) - 1]
//     }
//     queue = append(queue, i)
//     for queue[len(queue) - 1] - queue[0] + 1 > k {
//         queue = queue[1:]
//     }
//     if i >= k - 1 {
//         result = append(result, nums[queue[0]])
//     }
// }




// Runtime: 260 ms, faster than 59.45% of Go online submissions for Sliding Window Maximum.
// Memory Usage: 10.5 MB, less than 42.40% of Go online submissions for Sliding Window Maximum.
func maxSlidingWindow(nums []int, k int) []int {
    que := [][]int{}       // <index, num>
    ans := []int{}
    for i := 0; i < len(nums); i++ {
        t2 := i - k
        for len(que) > 0 && que[0][0] <= t2 {
            que = que[1 : ]
        }
        for len(que) > 0 && que[len(que) - 1][1] <= nums[i] {
            que = que[0 : len(que) - 1]
        }
        que = append(que, []int{i, nums[i]})
        if t2 >= -1 {
            ans = append(ans, que[0][1])
        }
    }
    return ans
}


func main_LT0239_20211125() {
// func main() {

    fmt.Println("ans:")


}
