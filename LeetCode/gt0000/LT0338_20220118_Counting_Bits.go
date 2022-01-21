// package sdq
package main

import (
    "fmt"
)



// answer[i] = answer[i - 1] - (numOfOne(i - 1) - 1);

// for(int i = 0; i <= n; i++){
//     a[i] = a[i/2];
//     if(i&1) a[i] += 1;
// }

// for(int i=1;i<=n;i++){
//     vec[i]=vec[i/2]+(i%2);
// }

// for(int i=1;i<=n;i++){
//     dp[i] = dp[i >> 1] + (i&1);
// }

// if(i==pow(2,k))
// {
//     dp.push_back(1);
//     k++;
// }
// else
// {
//     int z=pow(2,k-1);
//     dp.push_back(dp[z]+dp[i-z]);
// }





// Runtime: 16 ms, faster than 8.94% of Go online submissions for Counting Bits.
// Memory Usage: 4.9 MB, less than 33.74% of Go online submissions for Counting Bits.
// x&(x-1)
func countBits(n int) []int {
    ans := make([]int, n + 1)
    cnt := 0
    for i, _ := range ans {
        cnt = 0
        i2 := i
        for i > 0 {
            i &= (i - 1)
            cnt++
        }
        ans[i2] = cnt
    }
    return ans
}

func main_LT0338_20220118() {
// func main() {

    fmt.Println("ans:")

    // n := 2
    n := 5

    ans := countBits(n)

    fmt.Println(ans)

}
