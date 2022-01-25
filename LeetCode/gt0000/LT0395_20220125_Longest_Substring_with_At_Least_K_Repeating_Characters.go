// package sdq
package main

import (
    "fmt"
)




// def longestSubstring(self, s, k):
//     for c in set(s):
//         if s.count(c) < k:
//             return max(self.longestSubstring(t, k) for t in s.split(c))
//     return len(s)



// int longestSubstring(string s, int k) {
//     int n = (int)s.length();
//     unordered_map<char, int> count;
//     for (char c : s) ++count[c];
//     int mid = 0;
//     while (mid < n && count[s[mid]] >= k) ++mid;
//     if (mid == n) return n;
//     int left = longestSubstring(s.substr(0, mid), k);
//     while (mid < n && count[s[mid]] < k) ++mid;
//     int right = longestSubstring(s.substr(mid), k);
//     return max(left, right);
// }



// func helper(s string, left, right, k int) int{
//    
//     if right - left < k {
//         return 0
//     }
//     count := make([]int, 26)
//     for i:=left; i< right; i++{
//         count[s[i]-'a']++
//     }
//     for i:=left; i<right; i++{   
//         if count[s[i]-'a'] < k{
//             j := i+1
//             for j < right && count[s[j] - 'a'] < k{
//                 j++
//             }
//             return Max(helper(s, left, i, k), helper(s,j,right, k))
//         }
//     }
//     return right - left
// }
// 如果 这个 char 的数量 小于 k， 那么就以这个 char 作为划分


// 可以跳过多个。

// Runtime: 380 ms, faster than 6.67% of Go online submissions for Longest Substring with At Least K Repeating Characters.
// Memory Usage: 5.1 MB, less than 20.00% of Go online submissions for Longest Substring with At Least K Repeating Characters.
// 每个char 都要>= k
// aaabbbbbaa
// sliding window 好像不行。。 不知道 是 剔除前面的 还是 继续追加。
func longestSubstring(s string, k int) int {
    sz1 := len(s)
    arr := make([][]int, sz1)
    for i, _ := range arr {         // 这种 和  for i=0 i<sz i++ 的性能区别有多大
        arr[i] = make([]int, 26)
    }
    arr[0][s[0] - 'a']++
    for i := 1; i < sz1; i++ {
        for j := 0; j < 26; j++ {
            arr[i][j] = arr[i - 1][j]
        }
        arr[i][s[i] - 'a']++
    }
    // fmt.Println(arr)
    allk := true
    // for i := 0; i < 26; i++ {
    //     if arr[sz1 - 1][i] < k && arr[sz1 - 1][i] != 0 {
    //         allk = false
    //         break;
    //     }
    // }
    // if allk {
    //     return sz1
    // }
    fmt.Println(arr)
    ans := 0
    for i := sz1 - 1; i >= 0; i-- {
        allk = true
        for m := 0; m < 26; m++ {
            if arr[i][m] != 0 && arr[i][m] < k {
                allk = false
                break
            }
        }
        if allk {
            // if i + 1 > ans {
            //     ans = i + 1
            // }
            ans = i + 1
            break
        }
    }
    // fmt.Println(ans)
    for i := 0; i < sz1; i++ {
        for j := sz1 - 1; j >= i + k - 1; j-- {
            // fmt.Println(i, " - ", j)
            allk = true
            for m := 0; m < 26; m++ {
                t2 := arr[j][m] - arr[i][m]
                // if i == 0 && j == sz1 - 1 {
                //     fmt.Println(t2, k)
                // }
                if t2 != 0 && t2 < k {
                    // fmt.Println(" - - -", t2)
                    allk = false
                    break
                }
            }
            // fmt.Println(allk)
            if allk {
                if j - i > ans {
                    // fmt.Println(j, i)
                    ans = j - i
                }
                break;
            }
        }
    }
    return ans
}

func main_LT0395_20220125() {
// func main() {

    fmt.Println("ans:")

    // s := "aaabb"
    // k := 3

    // s := "ababbc"
    // k := 2

    s := "bbaaacbd"
    k := 3
    

    fmt.Println(longestSubstring(s, k))

}
