// package sdq
package main

import (
    "fmt"
)

// if height[l] < height[r] {
//     high = height[l]
//     l++
// } else {
//     high = height[r]
//     r--
// }
// 移动小的那端。



// error...
// 有高于自己的就不必计算了
func maxArea(height []int) int {
    sz1 := len(height)
    ans := 0
    mx := 0
    for i := 0; i < sz1; i++ {
        if (height[i] < mx) {
            continue
        }
        t2 := height[i]
        mx = t2
        for j := 0; j < i; j++ {
            if (height[j] >= t2) {
                t3 := t2 * (i - j)
                if (ans < t3) {
                    ans = t3
                }
                break
            }
        }
        for j := sz1 - 1; j > i; j-- {
            if (height[j] >= t2) {
                t3 := t2 * (j - i)
                if (ans < t3) {
                    ans = t3
                }
                break
            }
        }
    }
    return ans;
}


// Runtime: 223 ms, faster than 6.81% of Go online submissions for Container With Most Water.
// Memory Usage: 9.3 MB, less than 19.06% of Go online submissions for Container With Most Water.
func maxArea3(height []int) int {
    sz1 := len(height)
    ans := 0
    for i := 0; i < sz1; i++ {
        t2 := height[i]
        for j := 0; j < i; j++ {
            if (height[j] >= t2) {
                t3 := t2 * (i - j)
                if (ans < t3) {
                    ans = t3
                }
                break
            }
        }
        for j := sz1 - 1; j > i; j-- {
            if (height[j] >= t2) {
                t3 := t2 * (j - i)
                if (ans < t3) {
                    ans = t3
                }
                break
            }
        }
    }
    return ans
}


// tle
func maxArea2(height []int) int {
    sz1 := len(height)
    ans := 0
    for i := 0; i < sz1; i++ {
        for j := i + 1; j < sz1; j++ {
            t2 := (j - i) * min(height[i], height[j])
            if t2 > ans {
                ans = t2
            }
        }
    }
    return ans
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}


func main_LT0011_20211006() {
// func main() {

    fmt.Println("ans:")


}
