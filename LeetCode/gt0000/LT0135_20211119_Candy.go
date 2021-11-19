// package sdq
package main

import (
    "fmt"
)



// candys := make([]int, len(ratings))
// candys[0] = 1
// candys[len(ratings)-1] = 1
// for i:=0; i<len(candys); i++ {
//     if i>0 && ratings[i] > ratings[i-1] && candys[i] <= candys[i-1]{
//         candys[i] = candys[i-1]+1
//     } else if candys[i] == 0 {
//         candys[i] = 1
//     }
// }
// for i:=len(candys)-1; i>=0; i-- {
//     if i<len(candys)-1 && ratings[i] > ratings[i+1] && candys[i] <= candys[i+1]{
//         candys[i] = candys[i+1]+1
//     } else if candys[i] == 0 {
//         candys[i] = 1
//     }
// }
// for _, v := range candys {
//     res += v
// }


// for i := range ratings{
//     cnt[i] = 1
//     if i > 0 && ratings[i] > ratings[i-1]{
//         cnt[i] = cnt[i-1] + 1
//     }
// }
// for i := len(ratings)-1; i >= 0; i--{
//     if i != len(ratings)-1 && ratings[i] > ratings[i+1]{
//         cnt[i] = max(cnt[i+1]+1, cnt[i])
//     }
//     total += cnt[i]
// }
// return total


// for (int i = 0; i < n; i++) {
//     count[i] = 1;
// }
// for (int i = 1; i < n; i++) {
//     if (ratings[i] > ratings[i - 1]) {
//         count[i] = count[i - 1] + 1;
//     }
// }
// for (int i = n - 2; i >= 0; i--) {
//     if (ratings[i] > ratings[i + 1]) {
//         count[i] = Math.max(count[i + 1] + 1, count[i]);
//     }
// }


// vector<int>left(n, 1), right(n, 1);
// //traverse left to right and compare curr value with left side value
// for(int i = 1; i < n; i++)
//     if(ratings[i] > ratings[i-1]) left[i] = left[i-1]+1;
//  //traverse right to left and compare curr value with right side value
// for(int i = n-2; i >= 0; i--)
//     if(ratings[i] > ratings[i+1]) right[i] = right[i+1]+1;
// int ans = 0;
// for(int i = 0; i < n; i++)
//     ans += max(left[i], right[i]);


// for (int i = 1; i < ratings.length; i++) {
//     if (ratings[i] > ratings[i - 1]) {
//         if (i > 1 && ratings[i - 1] < ratings[i - 2]) inc = 1;
//         ans += ++inc;
//         dec = 1;
//     }
//     else if (ratings[i] < ratings[i - 1]) {
//         dec++;
//         ans += dec - 1;
//         if (dec > inc) ans++;
//     }
//     else {
//         ans++;
//         inc = 1;
//         dec = 1;
//     }
// }


// Runtime: 28 ms, faster than 23.81% of Go online submissions for Candy.
// Memory Usage: 6.8 MB, less than 15.87% of Go online submissions for Candy.
func candy(ratings []int) int {
    sz1 := len(ratings)
    arr := make([]int, sz1)

    for i := 0; i < sz1 - 1; i++ {
        if ratings[i] <= ratings[i + 1] {
            arr[i] = 1
            for t2 := i - 1; t2 >= 0; t2-- {
                if ratings[t2] > ratings[t2 + 1] {
                    if arr[t2] <= arr[t2 + 1] {
                        arr[t2] = arr[t2 + 1] + 1
                    } else {
                        // break        // 感觉这个 和 下面的break 下标 就差一位？
                    }
                } else {
                    break
                }
            }
            t2 := i + 1
            for ; t2 < sz1; t2++ {
                if ratings[t2] > ratings[t2 - 1] {
                    if arr[t2] <= arr[t2 - 1] {
                        arr[t2] = arr[t2 - 1] + 1
                    }
                } else {
                    break
                }
            }
            i = t2 - 1
        }
    }
    if arr[sz1 - 1] == 0 {
        arr[sz1 - 1] = 1
        for t2 := sz1 - 1 - 1; t2 >= 0; t2-- {
            if ratings[t2] > ratings[t2 + 1] {
                if arr[t2] <= arr[t2 + 1] {
                    arr[t2] = arr[t2 + 1] + 1
                } else {
                    // break        // 感觉这个 和 下面的break 下标 就差一位？
                }
            } else {
                break
            }
        }
    }
    fmt.Println(arr)
    ans := 0
    for _, v := range arr {
        ans += v
    }
    return ans
}


func main_LT0135_20211119() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,3,2,2,1}

    ans := candy(arr)

    fmt.Println(ans)

}
