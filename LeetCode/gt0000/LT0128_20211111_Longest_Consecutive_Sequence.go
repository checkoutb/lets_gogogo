// package sdq
package main

import (
    "fmt"
)



// while (!set.empty()) {
// while (set.find(current - i) != set.end()) {
//     set.erase(current - i);
//     i++;
//     count++;
// }
// erase...

// 最快的还是 sort， 然后 for。

// if _, ok := m[k-1]; ok {
//     continue
// }
// 有更小的就 不计算。

// 直接改 map2的value 。。


// unordered_map<int, int> m;
// int r = 0;
// for (int i : num) {
//     if (m[i]) continue;
//     r = max(r, m[i] = m[i + m[i + 1]] = m[i - m[i - 1]] = m[i + 1] + m[i - 1] + 1);
// }
// return r;
// .......... map[value]length
// 和另一个 uf 感觉差不多


// nums = set(nums)
// best = 0
// for x in nums:
//     if x - 1 not in nums:
//         y = x + 1
//         while y in nums:
//             y += 1
//         best = max(best, y - x)
// return best


// Runtime: 40 ms, faster than 85.53% of Go online submissions for Longest Consecutive Sequence.
// Memory Usage: 9 MB, less than 51.27% of Go online submissions for Longest Consecutive Sequence.
// O(n) ... 只能转set， 然后 再for一次。 但是肯定不是O(n)啊。 hash。
func longestConsecutive(nums []int) int {
    map2 := make(map[int]int)
    vst := 2100000000
    for i, v := range nums {
        if _, ok := map2[v]; ok {
            nums[map2[v]] = vst
        }
        map2[v] = i         // distinct...
    }
    ans := 0
    for _, v := range nums {
        // smaller bigger
        if v != vst {
            cnt := 1
            t2 := v - 1
            for _, ok := map2[t2]; ok; _, ok = map2[t2] {
                cnt++
                nums[map2[t2]] = vst
                t2--
            }
            t2 = v + 1
            for _, ok := map2[t2]; ok; _, ok = map2[t2] {
                cnt++
                nums[map2[t2]] = vst
                t2++
            }
            if cnt > ans {
                ans = cnt
            }
        }
    }
    return ans
}

func main_LT0128_20211111() {
// func main() {

    fmt.Println("ans:")


}
