// package sdq
package main

import (
    "fmt"
)


// for (int i = 2; i <= n; i++) {
//     int first = Integer.valueOf(s.substring(i - 1, i));
//     int second = Integer.valueOf(s.substring(i - 2, i));
//     if (first >= 1 && first <= 9) {
//        dp[i] += dp[i-1];  
//     }
//     if (second >= 10 && second <= 26) {
//         dp[i] += dp[i-2];
//     }
// }


// int[] memo = new int[n+1];
// memo[n]  = 1;
// memo[n-1] = s.charAt(n-1) != '0' ? 1 : 0;
// for (int i = n - 2; i >= 0; i--)
//     if (s.charAt(i) == '0') continue;
//     else memo[i] = (Integer.parseInt(s.substring(i,i+2))<=26) ? memo[i+1]+memo[i+2] : memo[i+1];
// return memo[0];


// for(int i=n-1;i>=0;i--)
//     if(s.charAt(i)!='0') {
//         dp[i]=dp[i+1];
//         if(i<n-1&&(s.charAt(i)=='1'||s.charAt(i)=='2'&&s.charAt(i+1)<'7')) 
//             dp[i]+=dp[i+2];
//     }



// Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode Ways.
// Memory Usage: 2 MB, less than 88.34% of Go online submissions for Decode Ways.
func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }
    ans := 0
    pre, ppre := 1, 1       // 1, 1
    lst := s[0] - '0'
    
    for i := 1; i < len(s); i++ {
        t2 := s[i] - '0'
        t3 := lst * 10 + t2
        t5 := 0

        // 能不能和前面 合并
        // 能不能 独立

        if t3 >= 10 && t3 <= 26 {       // 合并
            t5 += ppre
        }
        if t2 != 0 {        // duli
            t5 += pre
        }
        ppre, pre = pre, t5

        // fmt.Println(ppre, ", ", pre)

        // // t4 := pre
        // if t3 == 0 {
        //     return 0
        // } else if t3 < 10 {             // 102
        //     ppre = pre
        //     // pre = 
        //     // ppre = 0
        //     // ;
        //     // prev = prev
        // } else if t3 > 26 {             // 132
        //     ppre = pre
        // } else {
        //     if t2 == 0 {

        //     } else {

        //     }
        // }
        lst = t2
        // // ppre = t4
    }

    ans = pre
    return ans
}

func main_LT0091_20211101() {
// func main() {

    fmt.Println("ans:")

    // s := "12"
    // s := "226"
    s := "05"

    ans := numDecodings(s)

    fmt.Println(ans)

}
