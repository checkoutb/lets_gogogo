// package sdq
package main

import (
    "fmt"
)




// Stack<Integer> stack = new Stack<>();
// int maxArea = 0;
// for (int i = 0; i <= n; i++) { // i - right bound ), include n
//     int rightHeight = i == n? 0: heights[i]; // end then h[n] = 0, so whatever stack.peek() can be included as h * n - whole length
//     while (!stack.isEmpty() && rightHeight < heights[stack.peek()]){
//         int curHeight = heights[stack.pop()];
//         int prevIdx = stack.isEmpty()? -1: stack.peek();
//         int area = curHeight * (i - prevIdx - 1);
//         maxArea = Math.max(maxArea, area);
//     }
//     stack.push(i);
// }
// .. 这种是  当前 小于 stack 就处理. .. 


// heights = append(heights, 0) // The reason we have a 0 at the end is if the given heights is a sorted ascending array, we will push everything to the stack without doing anything.
// n := len(heights)
// result := 0
// stack := []int{}
// for i := 0; i < n;{
//     if len(stack) == 0 || heights[i] >= heights[stack[len(stack) - 1]] {
//         stack = append(stack, i)
//         i++
//     } else {
//         h := heights[stack[len(stack) - 1]]
//         stack = stack[:len(stack) - 1]
//         w := i
//         if len(stack) > 0 {
//             w = i - stack[len(stack) - 1] - 1
//         }
//         area := h * w
//         if result < area {
//             result = area
//         }
//     }
// }
// .. 只处理一次....


// for (int i = 0; i <= len; i++){
//     int h = (i == len ? 0 : heights[i]);
//     if (s.isEmpty() || h >= heights[s.peek()]) {
//         s.push(i);
//     } else {
//         int tp = s.pop();
//         maxArea = Math.max(maxArea, heights[tp] * (s.isEmpty() ? i : i - 1 - s.peek()));
//         i--;
//     }
// }

// Runtime: 84 ms, faster than 89.74% of Go online submissions for Largest Rectangle in Histogram.
// Memory Usage: 8.2 MB, less than 100.00% of Go online submissions for Largest Rectangle in Histogram.
// 后续中, 比我大的无所谓, 只看比我小的 及 再小的.
// 得 倒序.
func largestRectangleArea(heights []int) int {
    stk := []int{}          // idx
    ans := 0
    for i := len(heights) - 1; i >= 0; i-- {
        t2 := heights[i]
        t3, t5 := t2, len(heights)       // next, next's next
        for len(stk) > 0 && heights[stk[len(stk) - 1]] >= t2 {
            stk = stk[0 : len(stk) - 1]
        }
        if len(stk) > 0 {
            t5 = stk[len(stk) - 1]
            if t2 * (t5 - i) > ans {
                ans = t2 * (t5 - i)
            }
            t3 = heights[stk[0]]
            if t3 * (len(heights) - i) > ans {
                ans = t3 * (len(heights) - i)
            }
        } else {
            if t2 * (len(heights) - i) > ans {
                ans = t2 * (len(heights) - i)
            }
        }
        for j := len(stk) - 1; j > 0; j-- {
            t3 = heights[stk[j]]
            t5 = stk[j - 1]
            if t3 * (t5 - i) > ans {
                ans = t3 * (t5 - i)
            }
        }
        stk = append(stk, i)
    }
    return ans
}


func main_LT0084_20211027() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{2,1,5,6,2,3}
    arr := []int{3,4}

    ans := largestRectangleArea(arr)

    fmt.Println(ans)

}
