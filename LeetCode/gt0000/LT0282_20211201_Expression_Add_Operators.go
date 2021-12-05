// package sdq
package main

import (
    "fmt"
    // "strings"
    "strconv"
)


// dfs(num, target, i+1, curNum, val+curNum, exp+"+"+curNumStr, ans);
// dfs(num, target, i+1, -curNum/*minus ==> add an negative number*/, val-curNum, exp+"-"+curNumStr, ans);
// dfs(num, target, i+1, prevNum*curNum, val-prevNum+prevNum*curNum, exp+"*"+curNumStr, ans);               // -pre + pre*cur


// backtrack(path + "+" + num[pos:i+1], num, target, i+1, value, cur+value)
// backtrack(path + "-" + num[pos:i+1], num, target, i+1, -value, cur-value)
// backtrack(path + "*" + num[pos:i+1], num, target, i+1, prev*value, cur-prev+prev*value)






// +-*  no /
// Runtime: 52 ms, faster than 92.86% of Go online submissions for Expression Add Operators.
// Memory Usage: 7 MB, less than 92.86% of Go online submissions for Expression Add Operators.
func addOperators(num string, target int) []string {
    nums := []int{}
    // for _, v := range num {
    //     nums = append(nums, v - "0"[0])
    // }
    for i, _ := range num {
        nums = append(nums, int(num[i] - '0'))
    }
    ans := []string{}
    // fmt.Println(nums)
    t2 := 0
    for i := 0; i < len(nums); i++ {
        t2 *= 10
        t2 += nums[i]
        dfsaa282(nums, []int{t2}, i + 1, target, &ans)
        if t2 == 0 {
            break
        }
    }
    // dfsaa282(nums, []int{}, 0, target, ans)
    return ans
}


// + - * / append, no append
func dfsaa282(nums, arr []int, idx, tar int, ans *[]string) {
    if idx == len(nums) {
        // fmt.Println("--", arr)
        if evala282(arr, tar) == tar {
            *ans = append(*ans, toStringa282(arr))
        }
        return
    }
    if nums[idx] == 0 {
        for op := 1; op < 4; op++ {
            arr = append(arr, op)
            arr = append(arr, 0)
            dfsaa282(nums, arr, idx + 1, tar, ans)
            arr = arr[0 : len(arr) - 2]
        }
        return
    }
    t2 := 0
    for i := idx; i < len(nums); i++ {
        t2 *= 10
        t2 += nums[i]
        for op := 1; op < 4; op++ {             // '/' != '/'
            arr = append(arr, op)
            arr = append(arr, t2)
            dfsaa282(nums, arr, i + 1, tar, ans)
            arr = arr[0 : len(arr) - 2]
        }
    }
}

var (
    ops = [5]string{"","+","-","*","/"}
)
func toStringa282(arr []int) string {
    ans := strconv.Itoa(arr[0])
    // fmt.Println("...", arr)
    for i := 1; i < len(arr); i++ {
        if i % 2 == 0 {
            ans += strconv.Itoa(arr[i])
        } else {
            ans += ops[arr[i]]
        }
    }
    // fmt.Println(ans, "......")
    return ans
}

// 后缀表达式， 逆波兰表达式 。。 不。
// 1+  2-  3*  4/
func evala282(arr []int, tar int) (ans int) {
    ans, pre := 0, arr[0]
    // if len(arr) > 1 && arr[1] < 3 {
    //     ans = arr[0]
    // }
    for i := 1; i < len(arr); i += 2 {
        switch arr[i] {
        case 1:
            ans, pre = ans + pre, arr[i + 1]
        case 2:
            ans, pre = ans + pre, -arr[i + 1]
        case 3:
            pre *= arr[i + 1]
        case 4:
            // pre /= arr[i]           // 3/2 panic?   long
            if pre % arr[i + 1] != 0 {
                return tar - 123
            }
            pre /= arr[i + 1]
        }
        // if arr[i] > 2 {
        //     pre *= arr[i + 1]
        // } else {
        //     ans += pre
        // }
    }
    ans += pre
    // fmt.Println(arr, ">>> ", ans)
    return
}


// func dfsb282(arr []int, idx, sum2, pre, tar int, arr2 []string, ans []string) {
//     if idx == len(num) {
//         if sum2 + pre == tar {
//             ans = append(ans, strings.Join(arr2))
//         }
//         return
//     }
// }


// // eval ...
// func dfsa282(num string, idx int, sum2 int, pre int, tar int,str string, ans *[]string) {
//     if idx == len(num) {
//         if sum2 + pre == tar {
//             str2 := string{str}
//             ans = append(ans, str2)
//         }
//         return
//     }
//     for i := idx; i < len(num); i++ {

//     }

// }


func main_LT0282_20211201() {
// func main() {

    fmt.Println("ans:")

    // s := "123"
    // t := 6

    // s := "232"
    // t := 8

    // s := "105"
    // t := 5

    // s := "00"
    // t := 0

    // s := "3456237490"
    // t := 9191

    s := "123456789"
    t := 45

    fmt.Println(addOperators(s, t))


}
