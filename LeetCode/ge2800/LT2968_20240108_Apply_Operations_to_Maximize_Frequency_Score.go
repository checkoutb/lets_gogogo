// package sdq
package main

import (
    "fmt"
    "sort"
)


// D D


// func maxFrequencyScore(nums []int, k int64) int {
//     slices.Sort(nums)
// 	ans, left := 0, 0
// 	s := int64(0)
// 	for right, x := range nums {
// 		s += int64(x - nums[(left+right)/2])
// 		for s > k {
// 			s += int64(nums[left] - nums[(left+right+1)/2])
// 			left++
// 		}
// 		ans = max(ans, right-left+1)
// 	}
// 	return ans
// }

// 。。 想不出来，为什么 直接用 头/尾 和 median 就可以了。


// Runtime174 ms
// Beats
// 28.30%
// Memory9.7 MB
// Beats
// 26.41%

/*

1        |       100
 2       |      90
    10   |   60 
      22 | 40

| is the final element
operation is 100-1 + 90-2 + 60-10 + 40-22

1      |       100
 2     |      90
   10  |     60
       | 22 40

operation is 100-1 + 90-2 + 60-10 + 40-x + 22-x

40-x + 22-x  vs  40-22,  for x < 22, so  40-x + 22-x > 40-22



*/
func maxFrequencyScore(nums []int, k int64) int {
    sort.Ints(nums);

    pfxs := make([]int64, len(nums));
    pfxs[0] = int64(nums[0]);

    for i := 1; i < len(nums); i++ {
        pfxs[i] = pfxs[i - 1] + int64(nums[i]);
    }

    getPre := func(i int) int64 {
        if i == 0 {
            return 0;
        }
        return pfxs[i - 1];
    }

    // fmt.Println(pfxs);

    ans := 0;
    en := -1;       // [st, en]
    for st := range nums {
        for en + 1 < len(nums) {
            md := (st + en + 1) / 2;
            t2 := int64(nums[md]);
            t2 = t2 * int64(md - st + 1) - (pfxs[md] - getPre(st)) + (pfxs[en + 1] - getPre(md)) - t2 * int64(en - md + 2);

            // fmt.Println(st, en, md, t2);

            if (t2 > k) {
                break;
            }
            en++;
        }
        if (en - st + 1 > ans) {

            // fmt.Println(st, en);

            ans = en - st + 1;
        }
    }

    return ans;


    // t2 := nums[0];
    // st := 0;
    // en := 0;
    // var op int64 = 0;
    // ans := 0;
    // for i, num := range nums {

    //     op += (num - t2) * (i - st) - (num - t2) * (en - i + 1);
    //     for op < k {

    //     }

    //     if (en - st + 1 > ans) {
    //         ans = en - st + 1;
    //     }
    //     t2 = num;
    // }

}


func main_LT2968_20240108() {
// func main() {

    arr := []int{1,2,6,4};
    k := int64(3);


    fmt.Println("ans:", maxFrequencyScore(arr, k));


}
