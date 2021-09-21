
package main

import (
    "fmt";
    "sort"
)





// 很难受啊， 跟着C走， 没有 太多的 API，  看details， 基本都是 很多行的， 没有 py， ruby 那种 one-line 的。
// channel, go-routine 都是为了 工程存在的， 对于算法，没有什么价值啊。


// Runtime: 144 ms, faster than 90.00% of Go online submissions for Sell Diminishing-Valued Colored Balls.
// Memory Usage: 9 MB, less than 80.00% of Go online submissions for Sell Diminishing-Valued Colored Balls.
func lt1638c(inventory []int, orders int) int {
    sort.Ints(inventory)
    // fmt.Println(inventory)
    ans := 0
    t2 := 1
    lst := inventory[len(inventory) - 1]        // > this
    cnt2 := 0
    MOD := 1000000007             // ^  xor..
    for i := len(inventory) - 2; i >= 0; i-- {               // no ++i
        t3 := t2 * (lst - inventory[i])

        // fmt.Println(t3)

        if (cnt2 + t3 > orders) {
            break;
        }
        ans = (ans + (t2 * (lst - inventory[i]) * (inventory[i] + 1 + lst) / 2) % MOD) % MOD
        // fmt.Println(ans)
        // fmt.Println(MOD)

        cnt2 += t3
        t2++        // len() - i
        lst = inventory[i]      // inventory[i + 1]
    }

    // fmt.Println(t2)
    // fmt.Println(cnt2)
    // fmt.Println(ans)

    a2 := (orders - cnt2) / t2
    a3 := (orders - cnt2) % t2

    ans = (ans + (t2 * (lst - (lst - a2)) * (lst + (lst - a2) + 1) / 2) % MOD) % MOD
    ans = (ans + (lst - a2) * a3) % MOD
    return ans
}



type IntSlice []int

// 没有set。。map可以， 算了，直接 不去重了。
// sort 有点恐怖。。 好吧 有个简单的。
func lt1638b(inventory []int, orders int) int {
    // inventory.Sort()
    // var arr IntSlice = inventory
    // arr.Sort()

    // sort.Sort(IntSlice(inventory))

    sort.Ints(inventory)            // ......

    l, r := 0, len(inventory) - 1

    for l < r {
        mid := (l + r) / 2
        cnt2 := 0
        mv := inventory[mid]
        for _, v := range inventory {
            if (v >= mv) {
                // ... 不需要二分，， 只需要排序后， 从后往前 就可以了。。
                cnt2 += -1
            }
        }
    }
    return -1
}




// .. go.mod...


// 二分， 0 到 max()，
// priority_queue / ordered_map， 记录 值和个数。 每次把最大的 所有值 都 削减到 次大的。

// go 没有提供 优先队列。  map是乱序的。 。。 我。。
// import "container/heap"   有个heap
// 没有pair。。  得自己定义 struct。。
// 感觉很难玩啊。。。

// 二分吧。  也没有提供 max_element。。算了 直接10^9...

// 不应该 0-10^9， 应该把 切片 去重，排序， 然后 0-len(xx)   就是 下标。

func lt1638a(inventory []int, orders int) int {
    ans := 0
    l, r := 0, 10^9
    mdVal := 0
    // for l < r:           // not .py ...
    for l < r {
        mid := (l + r) / 2;
        // sum2 := 0            // 想差了， 不需要sum，主要是 累加sum 的时候 想到 溢出， 然后发现，只需要 个数就可以了。。
        cnt2 := 0
        mdVal = mid
        for _, v := range inventory {       // :=   not  in ...
            if v >= mid {           // 不写:  直接{} 有点爽。
                // sum2 += 
                cnt2 += (v - mid) + 1
            }
        }
        if cnt2 > orders {
            l = mid + 1
        } else {
            r = mid             // 强制{}
        }
    }

    var mod int = 10^9 + 7
    for _, v := range inventory {
        if (v >= mdVal) {
            // var cnt2 int64 = v - mdVal + 1       // 没有自动升级。。
            cnt2 := v - mdVal + 1           // int  64位os  就是 int64吧。
            ans = (ans + (cnt2 * (v + mdVal) / 2) % mod) % mod
        }
    }

    return ans
}


func main() {

    // var arr []int = []int{2,4,8,6,10}           // ...
    // var ord int = 20
    
    // arr := []int{1000000000}
    // ord := 1000000000

    arr := []int{3,5}
    ord := 6

    t2 := lt1638c(arr, ord)

    fmt.Println(t2)

    fmt.Println("aaa")
}
