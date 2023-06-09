// package sdq
package main

import (
    "fmt"
)

// D D

// map 保存了到目前为止，所有可能的lcm 以及 个数。

// int subarrayLCM(vector<int>& nums, int k) {
//     int res = 0;
//     unordered_map<int, int> m;
//     for (int n : nums) {
//         unordered_map<int, int> m1;
//         if (k % n == 0) {
//             ++m[n];    
//             for (auto &[d, cnt] : m)
//                 m1[lcm(d, n)] += cnt;
//             res += m1[k];
//         }
//         swap(m, m1);
//     }
//     return res;
// }


// lcm(a,b,c,d) = lcm(a, lcm(b, lcm(c, d)))



// Runtime58 ms
// Beats
// 33.33%
// Memory3.1 MB
// Beats
// 33.33%

// sub array
// 最小公倍数
// 1000
// [3,4] -> 12
// a*b*c*..x / gcd(a,b,c,...x)  ? no  , [2,2,3]  2*2*3 / 1 = 12  but 6
// [2,3,4]  == [3,4]
// 多个数的最小公倍数的方法：把每一个数分成质数相乘,找出每个算式的最大质数的个数,再把这些质数相乘的积就是他们的最小公倍数
// 1000 * 1000
func subarrayLCM(nums []int, k int) int {
    sz1 := len(nums)
    ans := 0
    if k == 1 {
        not1 := -1
        for i, v := range nums {
            if v == 1 {
                ans += i - not1
            } else {
                not1 = i
            }
        }
        return ans
    }

    // mx := nums[0]
    // for _, val := range nums {
    //     if val > mx {
    //         mx = val
    //     }
    // }
    prm := mk_prm(k + 1)
    sz2 := len(prm)
    arr := make([]int, sz2)
    for i, v := range prm {
        for k % v == 0 {
            k /= v
            arr[i]++
        }
    }

    prm2 := []int{}
    cnt2 := []int{}
    for i, v := range arr {
        if v != 0 {
            prm2 = append(prm2, prm[i])
            cnt2 = append(cnt2, v)
        }
    }
    // fmt.Println(prm2)
    // fmt.Println(cnt2)

    // vvi := make([][]int, sz1)
    // for i, v := range nums {

    // }

    for i := range nums {
        cnt3 := make([]int, len(cnt2))
        for j := i; j < sz1; j++ {
            t2 := nums[j]
            for i, v := range prm2 {
                t3 := 0
                for t2 % v == 0 {
                    t2 /= v
                    t3++
                }
                if t3 > cnt3[i] {
                    if t3 > cnt2[i] {
                        goto AAA
                    }
                    cnt3[i] = t3
                }
            }
            if t2 != 1 {
                goto AAA
            }
            for i, v := range cnt3 {
                if cnt2[i] != v {
                    goto BBB
                }
            }
            // fmt.Println(i, j)
            ans++

            BBB:
            continue
        }
        AAA:
        continue
    }

    return ans
}

func mk_prm(mx int) []int {
    var prm []int = []int{}
    vb := make([]bool, mx)
    for i := 2; i < mx; i++ {
        if vb[i] == false {
            for j := i + i; j < mx; j += i {
                vb[j] = true;
            }
        }
    }
    for i := 2; i < mx; i++ {
        if vb[i] == false {
            prm = append(prm, i)
        }
    }
    return prm
}

func main_LT2470_20230603() {
// func main() {

    var vi []int = []int{3,6,2,7,1}
    k := 6

    // vi := []int{3}
    // k := 2

    fmt.Println("ans:", subarrayLCM(vi, k))


}
