// package sdq
package main

import (
    "fmt"
)




// for n%4 == 0 {
//     n /= 4
// }
// if n%8 == 7 {
//     return 4
// }
// for x := 0; x*x <= n; x++ {
//     y := int(math.Sqrt(float64(n - x*x)))
//     if x*x+y*y == n {
//         if x == 0 || y == 0 {
//             return 1
//         } else {
//             return 2
//         }
//     }
// }
// return 3


// int numSquaresReal(int n) {
//     while (n % 4 == 0)
//         n /= 4;
//     if (n % 8 == 7)
//         return 4;
//     int a = -1, b = sqrt(n);
//     n -= b * b;
//     b += b + 1;
//     while (a <= b) {
//         if (n < 0)
//             n += b -= 2;
//         else if (n > 0)
//             n -= a += 2;
//         else
//             return a < 0 ? 1 : 2;
//     }
//     return 3;
// }


// dp := make([]int, n + 1)
// dp[1] = 1
// squares := []int{1}
// for i := 2; i*i <= n; i ++ {
//     squares = append(squares, i*i)
// }
// for i := 2; i <= n; i++ {
//     dp[i] = i // sum of one's
//     for j := 0; j < len(squares) && i - squares[j] >= 0; j++ {
//         if dp[i - squares[j]] + 1 < dp[i] {
//             dp[i] = dp[i - squares[j]] + 1
//         }
//     }
// }
// return dp[n]


// static vector<int> dp {0};
// int m = dp.size();
// dp.resize(max(m, n+1), INT_MAX);
// for (int i=1, i2; (i2 = i*i)<=n; ++i)
//     for (int j=max(m, i2); j<=n; ++j)
//         if (dp[j] > dp[j-i2] + 1)
//             dp[j] = dp[j-i2] + 1;
// return dp[n];


// 11/30/2021 21:31	Accepted	500 ms	6.3 MB	golang
// 02/04/2021 10:01	Time Limit Exceeded	N/A	N/A	cpp
// 02/04/2021 09:48	Wrong Answer	N/A	N/A	cpp
// 02/04/2021 09:45	Time Limit Exceeded	N/A	N/A	cpp
// 02/04/2021 09:44	Time Limit Exceeded	N/A	N/A	cpp
// 11/29/2018 10:31	Time Limit Exceeded	N/A	N/A	java
// 11/29/2018 10:16	Wrong Answer	N/A	N/A	java
// 11/29/2018 09:56	Time Limit Exceeded	N/A	N/A	java
// 11/29/2018 09:56	Time Limit Exceeded	N/A	N/A	java
// Runtime: 500 ms, faster than 17.81% of Go online submissions for Perfect Squares.
// Memory Usage: 6.3 MB, less than 71.22% of Go online submissions for Perfect Squares.
func numSquares(n int) int {
    arr := make([]int, n + 1)
    // arr[1] = 1
    for i := 1; i <= n; i++ {
        if i * i <= n {
            arr[i * i] = 1
        }
        if arr[i] == 1 {
            continue
        }
        arr[i] = n
        for j := 1; j <= i / 2; j++ {
            if arr[j] + arr[i - j] < arr[i] {
                arr[i] = arr[j] + arr[i - j]
            }
        }
    }
    return arr[n]
}


func main_LT0279_20211130() {
// func main() {

    fmt.Println("ans:")

    fmt.Println(numSquares(12))

}
