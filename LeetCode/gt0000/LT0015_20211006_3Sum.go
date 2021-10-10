// package sdq
package main

import (
    "fmt"
    "sort"
    "strconv"
)



// 左右逼近。这个。。
// if i > 0 && val == nums[i-1] {
//     // don't repeat values
//     continue
// }        
// l,r := i+1, len(nums)-1
// for l < r {
//     sum := val + nums[l] + nums[r]
//     if sum > 0 {
//         r--
//     } else if sum < 0 {
//         l++
//     } else {
//         triplets = append(triplets, []int{val, nums[l], nums[r]})
//         j := l+1
//         for j < r && nums[l] == nums[j] {                // 跳过重复的
//             j++
//         }
//         l = j
//     }
// }




// Runtime: 1136 ms, faster than 14.83% of Go online submissions for 3Sum.
// Memory Usage: 9.4 MB, less than 19.60% of Go online submissions for 3Sum.
func threeSum(nums []int) [][]int {
    sz1 := len(nums)
    ans := [][]int{}
    if sz1 <= 2 {
        return ans
    }
    sort.Ints(nums)
    map3 := make(map[int]int)
    for _, val := range nums {
        map3[val]++
    }
    mx := nums[sz1 - 1]
    map4 := map[string]int{}
    for i := 0; i < sz1; i++ {
        for j := i + 1; j < sz1; j++ {
            t2 := nums[i] + nums[j]
            t2 = -t2
            if t2 > mx {
                break;
            }
            cnt2 := 1
            if t2 == nums[i] {
                cnt2++
            }
            if t2 == nums[j] {
                cnt2++
            }
            if map3[t2] >= cnt2 {
                arr2 := []int{nums[i], nums[j], t2}
                sort.Ints(arr2)
                key2 := strconv.Itoa(arr2[0]) + "," + strconv.Itoa(arr2[1]) + "," + strconv.Itoa(arr2[2])
                _, ok := map4[key2]
                if !ok {
                    map4[key2] = 1
                    ans = append(ans, arr2)
                }
            }
        }
    }
    return ans
}


// 。。。。 这里是 a+b==c... 
// sort.
// 1 + 0 = 1.  记录最后一次出现。 bu.  负数的存在，导致 a+b < b
// 2 2 2 2 4 4
func threeSum2(nums []int) [][]int {
    // map2 := map[int][]int
    sort.Ints(nums)
    // for idx, val := range nums {
    //     append(map2[val], idx)              // ... nil？
    // }
    map3 := make(map[int]int)
    for _, val := range nums {
        map3[val]++             // ..
        // map3[val] = map3[val] + 1
    }
    sz1 := len(nums)
    mx := nums[sz1 - 1]
    ans := [][]int{}
    map4 := map[string]int{}
    for i := 0; i < sz1; i++ {
        for j := i + 1; j < sz1; j++ {
            t2 := nums[i] + nums[j]
            if t2 > mx {
                break;
            }
            cnt2 := 1
            if t2 == nums[i] {
                cnt2++
            }
            if t2 == nums[j] {
                cnt2++
            }
            if map3[t2] >= cnt2 {
                // ok
                // struct？ equal？ ==？
                // tuple   []int
                arr2 := []int{nums[i], nums[j], t2}
                sort.Ints(arr2)
                key2 := strconv.Itoa(arr2[0]) + "," + strconv.Itoa(arr2[1]) + "," + strconv.Itoa(arr2[2])
                _, ok := map4[key2]
                if !ok {
                    map4[key2] = 1
                    ans = append(ans, arr2)
                }
            }

            // if t2 == nums[i] || t2 == nums[j] {
            //     if nums[i] == nums[j] {
            //         cnt2 := 0
            //         i2 := i
            //         for i2 >= 0 && nums[i2] == nums[i] && cnt2 < 3 {
            //             cnt2++
            //             i2--
            //         }
            //         i2 = i + 1
            //         for i2 < sz1 && nums[i2] == nums[i] && cnt2 < 3 {
            //             cnt2++
            //             i2++
            //         }
            //         if cnt2 >= 3 {
            //             ans = append(ans, []int{nums[i], nums[j], t2})
            //         }
            //     } else if nums[i] == t2 {
            //         cnt2 := 0
            //         i2 := i
            //         for i2 >= 0 && nums[i2] == nums[i] && cnt2 < 2 {
            //             i2--
            //             cnt2++
            //         }
            //         i2 = i + 1
            //         for i2
            //     } else {
            //         cnt2 := 0
            //         j2 := j
            //     }
            // }

            // if () {
            //     arr, ok := map2[val]            // 空数组 or nil？
            //     if ok {
                    
            //     }
            // }
        }
    }
    return ans
}


func main_LT0015_20211006() {
// func main() {

    fmt.Println("ans:")

    arr := []int{-1,0,1,2,-1,-4}

    arr2 := threeSum(arr)

    for _, ar := range arr2 {
        for _, a := range ar {
            fmt.Print(a)
            fmt.Print(", ")
        }
        fmt.Println()
    }

}
