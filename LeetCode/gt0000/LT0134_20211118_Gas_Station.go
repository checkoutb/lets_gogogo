// package sdq
package main

import (
    "fmt"
)


// for (int i=0; i < n; i++){
//     /*add the gas avail at station i and
//     subtract the cost to reach station i.*/
//        total += (gas[i] - cost[i]);
//        cur += (gas[i] - cost[i]);
//        /*update curr_gas_tank and starting position if its < 0.*/
//        if (cur < 0){
//            pos = i +1;
//            cur = 0;
//        }
//    }




// 应该是 后移到 一个 gas.part.sum > cost 的地方。保证 中间油多。
// bu, 移动后 如果 gas更少了， 那么现在 也是过不了的。

// Runtime: 372 ms, faster than 9.09% of Go online submissions for Gas Station.
// Memory Usage: 10.1 MB, less than 25.00% of Go online submissions for Gas Station.
// 是不是 只要 sum(gas) >= sum(cost) 就必然可以？ ... return int...
// 硬算也可以， 从0开始，直到走不动，然后 0-1 继续尝试。 就是一个窗口。 0-1的话 可以知道 tank 减少了 gas[0] 增加了cost[0]。。。不，中间可能油不够。 
func canCompleteCircuit(gas []int, cost []int) int {




    // tle...
    sz1 := len(gas)
    for i := 0; i < sz1; i++ {

        if i > 0 && gas[i] == gas[i - 1] && cost[i] == cost[i - 1] {            // AC .....
            continue
        }

        gs, cs, j := 0, 0, 0
        for ; j <= sz1 && gs >= cs; j++ {
            gs += gas[(i + j) % sz1]
            cs += cost[(i + j) % sz1]
        }
        if j > sz1 && gs >= cs {
            return i
        }
    }
    return -1


    // sz1 := len(gas)
    // for i := 0; i < sz1; i++ {
    //     gs, cs, j, b2 := gas[i], cost[i], (i + 1) % sz1, true
    //     for gs >= cs && (j != (i + 2) % sz1 || b2) {

    //         fmt.Println(i, gs, cs, j)

    //         b2 = false
    //         gs += gas[j]
    //         cs += cost[j]
    //         j = (j + 1) % sz1
    //     }
    //     if j == (i + 1) % sz1 && gs >= cs {
    //         return i
    //     }
    // }
    // return -1
    

    // sum1, sum2 := 0, 0
    // for i, v := range gas {
    //     sum1 += v
    //     sum2 += cost[i]
    // }
    // return sum1 >= sum2
}


func main_LT0134_20211118() {
// func main() {

    fmt.Println("ans:")

    // arr1 := []int{5,1,2,3,4}
    // arr2 := []int{4,4,1,5,1}

    arr1 := []int{1,2,3,4,5}
    arr2 := []int{3,4,5,1,2}

    ans := canCompleteCircuit(arr1, arr2)

    fmt.Println(ans)

}
