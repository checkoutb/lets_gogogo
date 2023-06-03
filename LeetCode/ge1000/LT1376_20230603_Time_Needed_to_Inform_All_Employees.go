// package sdq
package main

import (
    "fmt"
)




// Runtime191 ms
// Beats
// 79.41%
// Memory13.1 MB
// Beats
// 67.65%

type Pair struct {
    fst int
    snd int
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    arr := make([][]int, n)
    head := -1
    for i, val := range manager {   // employee id, manager id
        if val == -1 {
            head = i
            continue
        }
        arr[val] = append(arr[val], i);
    }
    que := []Pair{Pair{fst: head, snd: 0}}      // person, time used
    ans := 0
    for len(que) > 0 {
        var szq int = len(que)
        for i := 0; i < szq; i++ {
            p := que[i]
            fst := p.fst;
            //snd := p.snd;
            tm := p.snd + informTime[fst]
            if tm > ans {
                ans = tm
            }
            for _, nxt := range arr[fst] {
                que = append(que, Pair{nxt, tm})
            }
        }
        que = que[szq : ]
    }
    return ans;
}


func main_LT1376_20230603() {
// func main() {

    n := 6
    id := 2
    vi := []int{2,2,-1,2,2,2}
    v2 := []int{0,0,1,0,0,0}

    fmt.Println("ans:", numOfMinutes(n,id,vi,v2))


}
