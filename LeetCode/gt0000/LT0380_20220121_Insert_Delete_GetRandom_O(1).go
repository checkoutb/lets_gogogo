// package sdq
package main

import (
    "fmt"
)




// discuss
// swap。 数组 只需要更新 最后一个元素的 map就可以了。。。
// 不需要swap，反正 最后一个元素 也是要删掉的，所以 直接 把 最后一个 放到前面 就可以了。


// 重复的是 map<int, set<int>> 保存 值-下标


type RandomizedSet struct {
    // map2 map[int]bool
    arr []int
    map3 map[int]int
}

func Constructor() RandomizedSet {
    // return RandomizedSet{ map[int]bool{} }
    return RandomizedSet{ []int{}, map[int]int{} }
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.map3[val]; !ok {
        this.arr = append(this.arr, val)
        this.map3                                   // 不行， 下标无法更新。  所以还是得 循环链表 + map[int]*Node
        return true
    }
    return false

    // if _, ok := this.map2[val]; !ok {
    //     this.map2[val] = true
    //     return true
    // }
    // return false
}

func (this *RandomizedSet) Remove(val int) bool {
    // if _, ok := this.map2[val]; ok {
    //     delete(this.map2, val)
    //     return true
    // }
    // return false
}


// 感觉是 循环链表 + map[int]*Node.. 差不多就是 数组+map。
func (this *RandomizedSet) GetRandom() int {


    // 不修改 map的话，第一位不会变。
    // for k, _ := range this.map2 {
    //     return k
    // }
    // return -1
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// func main_LT0380_20220121() {
func main() {

    fmt.Println("ans:")

    // map2 := map[int]bool{}
    // // arr := []int{}
    // for i := 1; i < 100; i++ {
    //     map2[i] = true
    //     for k, _ := range map2 {
    //         fmt.Println(k)
    //         break
    //     }
    // }

}
