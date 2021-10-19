// package sdq
package main

import (
    "fmt"
)




// int[] products = new int[n1 + n2];
// for (int i = n1 - 1; i >= 0; i--) {
//     for (int j = n2 - 1; j >= 0; j--) {
//         int d1 = num1.charAt(i) - '0';
//         int d2 = num2.charAt(j) - '0';
//         products[i + j + 1] += d1 * d2;
//     }
// }


// for (int i = num1.size() - 1; 0 <= i; --i) {
//     int carry = 0;
//     for (int j = num2.size() - 1; 0 <= j; --j) {
//         int tmp = (sum[i + j + 1] - '0') + (num1[i] - '0') * (num2[j] - '0') + carry;
//         sum[i + j + 1] = tmp % 10 + '0';
//         carry = tmp / 10;
//     }
//     sum[i] += carry;
// }
// size_t startpos = sum.find_first_not_of("0");
// if (string::npos != startpos) {





// Runtime: 4 ms, faster than 58.73% of Go online submissions for Multiply Strings.
// Memory Usage: 2.8 MB, less than 65.08% of Go online submissions for Multiply Strings.

//     1111
//      111
// ---
//     1111
//    1111
//   1111
//   123321

// 最后一位等于 最后×最后                                   1+1                 2
// 最后第二位=  最后第二×最后 + 最后×最后第二 + carry           2+1  1+2           3
// 3 最后第三×最后，最后第二×最后第二，最后×最后第三 + carry    3+1 2+2 1+3           4
// 4 4*最后，3×最后2, 2×最后3 1×最后4(没有4)。              4+1             5

func multiply(num1 string, num2 string) string {
    crr := 0
    ans := []byte{}
    sz1, sz2 := len(num1), len(num2)
    arr1 := make([]int, sz1)
    arr2 := make([]int, sz2)
    for i := 0; i < sz1; i++ {
        arr1[sz1 - i - 1] = int(num1[i] - '0')
    }
    for i := 0; i < sz2; i++ {
        arr2[sz2 - 1 - i] = int(num2[i] - '0')          // 低位 在 低位
    }
    // fmt.Println(arr1, " - - ", arr2)
    for i := 2; i <= (sz1 + sz2); i++ {
        t5 := crr
        for j := 1; j < i; j++ {
            t2 := j - 1
            t3 := (i - j) - 1
            if t2 >= sz1 || t3 >= sz2 {
                continue                // not break ...
            }
            t5 += arr1[t2] * arr2[t3]
        }
        // fmt.Println(t5, ", ", i, ", ", crr)
        // fmt.Println(t5)
        ans = append(ans, byte(t5 % 10 + '0'))
        crr = t5 / 10
    }
    if crr > 0 {
        ans = append(ans, byte(crr + '0'))
    }
    // fmt.Println(len(ans), ", ", ans[0])
    l, r := 0, len(ans) - 1
    for l < r {
        ans[l], ans[r] = ans[r], ans[l]
        l++
        r--
    }
    for len(ans) > 1 && ans[0] == '0' {
        ans = ans[1 : len(ans)]
    }
    return string(ans)
}


func main_LT0043_20211019() {
// func main() {

    fmt.Println("ans:")

    // s1, s2 := "2", "3"
    // s1, s2 := "123", "456"
    // s1, s2 := "1111", "111"
    s1, s2 := "100", "200"
    // s1, s2 := "123", "0"
    // s1, s2 := "0", "0"

    ans := multiply(s1, s2)

    fmt.Println(ans)

}

//         123
//         456
//       738
//      615
//     492
// 56088
