// package sdq
package main

import (
    "fmt"
    // "sort"
)



// type SummaryRanges struct {
//     bmp [157]uint64
// }

// func (this *SummaryRanges) AddNum(val int)  {
//     this.bmp[val/64] |= 1<<(val%64)
// }

// for i, n := range this.bmp {
//     for bi:=0; bi<64; {
//         if in {
//             cnt := bits.TrailingZeros64(^n)
//             bi += cnt
//             if bi < 64 {
//             b = 64*i+bi-1
//             n >>= cnt
//             in = false
            
//             result = append(result, []int{a, b})
//             }
//         } else {
//             cnt := bits.TrailingZeros64(n)
//             bi += cnt
//             if bi < 64 {
//             a = 64*i+bi
//             n >>= cnt
//             in = true
//             }
//         }
//     }
// }


// TreeMap


// for i <= j
// Runtime: 94 ms, faster than 25.00% of Go online submissions for Data Stream as Disjoint Intervals.
// Memory Usage: 10.2 MB, less than 75.00% of Go online submissions for Data Stream as Disjoint Intervals.
// 压缩，应该类似 桶排序。但是 合并，组织。
// 桶就直接是结果了。
// ... 不需要桶。直接在 ans 上操作就可以了。。
type SummaryRanges struct {
    arr []Bucket
}

type Bucket struct {
    st int
    en int
}

func Constructor() SummaryRanges {
    return SummaryRanges{ []Bucket{} }
}

func bs(arr *[]Bucket, val int) {
    i, j := 0, len(*arr) - 1
    for i <= j {
        mid := (i + j) / 2
        mdbk := (*arr)[mid]
        if mdbk.st - 1 <= val && mdbk.en + 1 >= val {
            if mdbk.st <= val && mdbk.en >= val {
                return
            } else {
                if val == mdbk.st - 1 {
                    if mid > 0 {
                        if (*arr)[mid - 1].en + 1 == val {
                            (*arr)[mid - 1].en = (*arr)[mid].en
                            *arr = append((*arr)[0 : mid], (*arr)[mid + 1 : ]...)
                            return
                        }
                    }
                    (*arr)[mid].st = val
                    return
                } else {
                    if mid < len(*arr) - 1 {
                        if (*arr)[mid + 1].st - 1 == val {
                            (*arr)[mid + 1].st = (*arr)[mid].st
                            *arr = append((*arr)[0 : mid], (*arr)[mid + 1 : ]...)
                            return
                        }
                    }
                    (*arr)[mid].en = val
                    return
                }
            }
        } else if mdbk.en < val {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }

    // fmt.Println(i, j, (*arr))

    // 无法找到 就 new
    bk := Bucket{ val, val }
    // 切片的修改，感觉是个坑。。。
    arr2 := make([]Bucket, len(*arr))     // 复制杀内存啊。
    copy(arr2, *arr)
    *arr = append(append((*arr)[ : i], bk), arr2[i : ]...)
}

// 二分搜索。。
func (this *SummaryRanges) AddNum(val int)  {
    bs(&this.arr, val)
}

func (this *SummaryRanges) GetIntervals() [][]int {
    ans := [][]int{}
    for _, v := range this.arr {
        ans = append(ans, []int{ v.st, v.en })
    }
    // // fmt.Println(ans)
    // for i,j := 0,len(ans)-1; i < j; j-- {
    //     ans[i], ans[j] = ans[j], ans[i]
    //     i++
    // }
    return ans
}



// // Runtime: 824 ms, faster than 12.50% of Go online submissions for Data Stream as Disjoint Intervals.
// // Memory Usage: 9.8 MB, less than 87.50% of Go online submissions for Data Stream as Disjoint Intervals.
// type SummaryRanges struct {
//     sor bool
//     arr []int
// }

// func Constructor() SummaryRanges {
//     return SummaryRanges{ true, []int{} }
// }

// func (this *SummaryRanges) AddNum(val int)  {
//     this.arr = append(this.arr, val)        // 可以 set 去重。
//     this.sor = false
// }

// func (this *SummaryRanges) GetIntervals() [][]int {
//     ans := [][]int{}
//     if len(this.arr) == 0 {
//         return ans
//     }
//     if !this.sor {
//         sort.Ints(this.arr)
//         this.sor = true
//     }
//     // 不如缓存结果。。 但是有必要吗。 不知道怎么调用的。。
//     st := this.arr[0]
//     for i := 1; i < len(this.arr); i++ {
//         if this.arr[i] > this.arr[i - 1] + 1 {
//             ans = append(ans, []int{st, this.arr[i - 1]})
//             st = this.arr[i]
//         }
//     }
//     ans = append(ans, []int{st, this.arr[len(this.arr) - 1]})
//     return ans
// }
/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */


func main_LT0352_20220119() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,3,7,2,6}

    obj := Constructor()
    for _, v := range arr {
        fmt.Println("  --- insert ", v)
        obj.AddNum(v)
        fmt.Println(obj.GetIntervals())
    }

}
