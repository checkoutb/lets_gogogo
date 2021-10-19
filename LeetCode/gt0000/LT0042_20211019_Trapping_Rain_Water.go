// package sdq
package main

import (
    "fmt"
)


// while(left<=right)
// 想起来了。。。  如果 左侧 低于 右侧，用左侧， 因为有 右侧兜底。。。

// left, right := 0, len(height)-1
// leftMax, rightMax := 0, 0
// for left < right {
//     if height[left] < height[right] {
//         if height[left] > leftMax {
//             leftMax = height[left]
//         }else {
//             ans += leftMax - height[left]
//         }
//         left++
//     }else {
//         if height[right] > rightMax {
//             rightMax = height[right]
//         }else {
//             ans += rightMax - height[right]
//         }
//         right--
//     }
// }

// st := list.New()
// top := st.Back().Value.(int)
// pop := st.Remove(st.Back()).(int)
// top = st.Back().Value.(int)
// st.PushBack(i)


// Runtime: 16 ms, faster than 19.87% of Go online submissions for Trapping Rain Water.
// Memory Usage: 3.8 MB, less than 56.70% of Go online submissions for Trapping Rain Water.

// 从左往右 遍历，查看 左侧 最高的，然后 跳到最高的上。  需要一个 stack<pair<index, height>>   bu, array<pair<idx, height>> 也可以， 好像不需要stack的特性
// 好像不是。。记录最高的就行，只是为了 确保 自己。
// 也可以直接 向前搜索，直到一个 大于等于自己的。或者最大值。
func trap(height []int) int {
    ans, sz1 := 0, len(height)
    for i := 0; i < sz1; i++ {
        if height[i] != 0 {
            mx := 0
            mxidx := i
            for j := i + 1; j < sz1; j++ {
                if height[j] > mx {
                    mxidx = j
                    mx = height[j]
                    if mx >= height[i] {
                        break
                    }
                }
            }
            t2 := height[i]
            if mx < t2 {
                t2 = mx
            }
            i++
            for i < mxidx {
                ans += t2 - height[i]
                i++
            }
            i--
        }
    }
    return ans
}

// type Pair struct {
//     i int
//     h int
// }

func main_LT0042_20211019() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}
    // arr := []int{4,2,0,3,2,5}
    arr := []int{0,0,0,5,3,1,2,0}
    // arr := []int{4,2,3}

    ans := trap(arr)

    fmt.Println(ans)

}
