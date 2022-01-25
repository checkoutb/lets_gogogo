// package sdq
package main

import (
    "fmt"
)


// int n;
// vector<int> result;
// void solve(int num){
//     for(int i = 0;i <= 9;i++){
//         if(num*10+i <= n){
//             result.push_back(num*10+i);
//             solve(num*10+i);
//         }
//     }
// }
// vector<int> lexicalOrder(int n) {
//     this->n = n;
//     for(int i = 1;i <= 9;i++){
//         if(i > n) return result;
//         result.push_back(i);
//         solve(i);
//     }
//     return result;
// }


// void fun(int x, int n, vector<int> &ans){
//     if(x>n) return;
//     ans.push_back(x);
//     fun(x*10, n, ans);
//     if(x%10!=9) fun(x+1, n, ans);
// }


// func dfs(x, n int, res *[]int) {
// 	limit := (x + 10) / 10 * 10
// 	for x <= n && x < limit {
// 		*res = append(*res, x)
// 		if x*10 <= n {
// 			dfs(x*10, n, res)
// 		}
// 		x++
// 	}
// }


// for (int i = 1; i <= n; i++) {
//     list.add(curr);
//     if (curr * 10 <= n) {
//         curr *= 10;
//     } else if (curr % 10 != 9 && curr + 1 <= n) {
//         curr++;
//     } else {
//         while ((curr / 10) % 10 == 9) {
//             curr /= 10;
//         }
//         curr = curr / 10 + 1;
//     }
// }



// Runtime: 8 ms, faster than 92.59% of Go online submissions for Lexicographical Numbers.
// Memory Usage: 6.5 MB, less than 85.19% of Go online submissions for Lexicographical Numbers.
// O(1) 空间，不可能吧，毕竟 ans 就是 O(n)了。 估计是 不能用辅助数组。
// O(n) 时间。。就不可能sort了。 自定义compare 不能用了。
// 1 10 100 1000 1001 1002 1003 1004 .... 1999 -> 2000
// 碰到 X000 直接 X,X0,X00,X000 然后继续 +1
// max 4567
// 1 10 100 1000 1001 ... 1999 2 20 200 2000 2001 2002 ... 2999 3 30 300
// ...bushi
// 1 10 101 1011 11 111 1111 12 121 1211 13 131 14 15 16 17 18 19          20
// 优先 ans[i-1]*10 + 0, 不行 就 ans[i-1]+1 .. 还有 1999 + 1 -> 2000 需要 变成 2.  %10啊。  119 + 1 = 120 -> 12
// 还有 *10 > n && +1 > n 需要    23 1 10 11 12... 19 2 20 21 22 23 3
// 234  234 24 25 26 29 30-3
func lexicalOrder(n int) []int {
    ans := make([]int, n)
    ans[0] = 1
    for i := 1; i < n; i++ {
        t2 := ans[i - 1] * 10
        if t2 > n {
            t2 = ans[i - 1] + 1
            if t2 > n {
                // t2 /= 10             // ... 1496...  14959+1 -> 14960  /10 1496 ++ -> 1497.... 跳过了 1496.。
                t2 = ans[i - 1] / 10
                t2++
            }
            for t2 % 10 == 0 {
                t2 /= 10
            }
        }
        ans[i] = t2
    }
    return ans
}


func main_LT0386_20220124() {
// func main() {

    fmt.Println("ans:")


}
