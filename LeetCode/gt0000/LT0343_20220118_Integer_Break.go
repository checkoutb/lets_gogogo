// package sdq
package main

import (
    "fmt"
)





// If an optimal product contains a factor f >= 4, 
// then you can replace it with factors 2 and f-2 without losing optimality, 
// as 2*(f-2) = 2f-4 >= f. 
// So you never need a factor greater than or equal to 4


// while(n>4){
//     product*=3;
//     n-=3;
// }
// product*=n;


// for(int i = 2; i <= n; i ++) {
//     for(int j = 1; j < i; j ++) {
//         dp[i] = Math.max(dp[i], (Math.max(j,dp[j])) * (Math.max(i - j, dp[i - j])));
//     }
// }


// if(n==2)return 1;
// if(n==3)return 2;
// if(n==4)return 4;
// if(n==5)return 6;
// int dp[n+1];
// dp[1]=1,dp[2]=2,dp[3]=3,dp[4]=4,dp[5]=6;
// for(int i=6;i<=n;i++){
//     dp[i]=i;
//     for(int j=1;j<i;j++){
//         dp[i]=max(dp[i],dp[j]*dp[i-j]);
//     }
// }
// return dp[n];



// while (i <= n)
// {
//     t1 = 1;
//     t3 = n;
//     for (int j = i; j > 0; j--)
//     {
//         t2 = t3 / j;
//         t3 -= t2;
//         t1 *= t2;
// 分成 j 个整数， 然后分成 j-1 个整数


// Runtime: 2 ms, faster than 21.15% of Go online submissions for Integer Break.
// Memory Usage: 2 MB, less than 85.58% of Go online submissions for Integer Break.
// 我猜 是 /3 。。。 不不不，反正肯定是 分解后 的 值 越集中越好。
func integerBreak(n int) int {
    ans := 1
    for i := 2; i < n; i++ {        // 分成多少个。
        t2 := n / i             // value
        t3 := n - n / i * i     // count(value+1)
        t4 := i - t3            // count(value)
        a := 1
        for j := 0; j < t4; j++ {
            a *= t2
        }
        for j := 0; j < t3; j++ {
            a *= (t2 + 1)
        }
        if a > ans {
            ans = a
        }
    }
    return ans
}

func main_LT0343_20220118() {
// func main() {

    fmt.Println("ans:")


}
