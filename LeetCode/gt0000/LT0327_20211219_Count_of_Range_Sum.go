// package sdq
package main

import (
    "fmt"
)


// gg
// merge sort

// public int countRangeSum(int[] nums, int lower, int upper) {
//     int n = nums.length;
//     long[] sums = new long[n + 1];
//     for (int i = 0; i < n; ++i)
//         sums[i + 1] = sums[i] + nums[i];
//     return countWhileMergeSort(sums, 0, n + 1, lower, upper);
// }
// private int countWhileMergeSort(long[] sums, int start, int end, int lower, int upper) {
//     if (end - start <= 1) return 0;
//     int mid = (start + end) / 2;
//     int count = countWhileMergeSort(sums, start, mid, lower, upper) 
//               + countWhileMergeSort(sums, mid, end, lower, upper);
//     int j = mid, k = mid, t = mid;
//     long[] cache = new long[end - start];
//     for (int i = start, r = 0; i < mid; ++i, ++r) {
//         while (k < end && sums[k] - sums[i] < lower) k++;
//         while (j < end && sums[j] - sums[i] <= upper) j++;
//         while (t < end && sums[t] < sums[i]) cache[r++] = sums[t++];
//         cache[r] = sums[i];
//         count += j - k;
//     }
//     System.arraycopy(cache, 0, sums, start, t - start);
//     return count;
// }
// 。。。。。。


// tle。。。。。
// 需要一个map，来 lower_bound(now - low)
// bst？
func countRangeSum(nums []int, lower int, upper int) int {
    arr := []int{}
    ans := 0
    sum2 := 0
    for i := 0; i < len(nums); i++ {
        sum2 += nums[i]
        if sum2 >= lower && sum2 <= upper {
            ans++
        }
        t2, t3 := sum2 - lower, sum2 - upper
        i3 := lowerBounda327(arr, t3)
        i2 := 0
        if i3 != -1 {
            if arr[i3] <= t2 {
                i2 = lowerBounda327(arr, t2)
                if i2 == -1 {
                    ans += len(arr) - i3
                } else {
                    for (i2 + 1) < len(arr) && arr[i2 + 1] == t2 { i2++ }
                    if arr[i2] > t2 { i2-- }
                    ans += i2 - i3 + 1
                }
            }
        }
        // i2 := lowerBounda327(arr, t2)
        // if i2 != -1 {
        //     i3 := lowerBounda327(arr, t3)
        // }
        i := lowerBounda327(arr, sum2)          // 把 for的 i 覆盖了。。
        if i == -1 {
            arr = append(arr, sum2)
        } else {
            // arr = append(append(arr[0 : i], sum2), arr[i : ]...)
            arr2 := make([]int, len(arr) - i)
            copy(arr2, arr[i : ])
            // fmt.Println("=-========= ", arr, arr2, i)
            arr = arr[0 : i]
            // fmt.Println("===== == ", arr)
            arr = append(arr, sum2)
            arr = append(arr, arr2...)
        }

        // fmt.Println(i, " - - ", arr, i2, i3, ans, sum2, t2, t3)
    }
    return ans
}

// -1 : all less than v
func lowerBounda327(arr []int, v int) int {
    l, r := 0, len(arr) - 1
    if l > r {
        return -1
    }
    for l < r {
        mid := (l + r) / 2
        if arr[mid] >= v {
            r = mid
        } else {
            l = mid + 1
        }
    }
    if arr[l] < v {
        return -1
    }
    return l
}

func checka327(arr []int, l, u int) {
    for i := 1; i < len(arr); i++ {
        arr[i] += arr[i - 1]
    }
    ans := 0
    for i := 0; i < len(arr); i++ {
        if arr[i] >= l && arr[i] <= u {
            fmt.Println("[0, ", i)
            ans++
        }
        for j := 0; j < i; j++ {
            t2 := arr[i] - arr[j]
            if t2 >= l && t2 <= u {
                fmt.Println("[", j + 1, ", ", i, "]")
                ans++
            }
        }
        fmt.Println(i, "------- ", ans)
    }
}

func main_LT0327_20211219() {
// func main() {

    fmt.Println("ans:")

    // arr := []int{-2,5,-1}
    // l, u := -2, 2

    // arr := []int{0}
    // l,u := 0,0

    // arr := []int{2147483647,-2147483648,-1,0}           // 4  [0,1], [2], [3], [2,3]
    // l, u := -1, 0
    
    // arr := []int{0,-1,-2,-3,0,2}        // 0
    // l, u := 3, 5

    // 0,00,0,00-3,0-3,-3,00-32,0-32,-32,00-32-2,0-32-2,-32-2,2-2,-2,2-2-2,-2
    arr := []int{0,0,-3,2,-2,-2}        // 16
    l := -3
    u := 1
    

    fmt.Println(countRangeSum(arr, l, u))

    // checka327(arr, l, u)

}

// now - pre >= low    <= up
// now - low >= pre
// now - up <= pre
// 0  - -  [-5 -3 -3 -1 0 0] 2 0 17 -5 -2 -6



