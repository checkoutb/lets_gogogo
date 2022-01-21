// package sdq
package main

import (
    "fmt"
    "sort"
)





// List<Integer>[] bucket = new List[nums.length + 1];
// for (int key : frequencyMap.keySet()) {
//     int frequency = frequencyMap.get(key);
//     if (bucket[frequency] == null) {
//         bucket[frequency] = new ArrayList<>();
//     }
//     bucket[frequency].add(key);
// }



// Runtime: 25 ms, faster than 19.49% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 5.6 MB, less than 56.99% of Go online submissions for Top K Frequent Elements.
func topKFrequent(nums []int, k int) []int {
    map2 := map[int]int{}
    for _, v := range nums {
        map2[v]++
    }
    arr := []int{}
    for _, v := range map2 {
        arr = append(arr, v)
    }
    sort.Ints(arr)
    ans := []int{}
    t2 := arr[len(arr) - k]
    // fmt.Println(arr, t2)
    for k, v := range map2 {
        if v >= t2 {
            ans = append(ans, k)
        }
    }
    return ans
}


func main_LT0347_20220118() {
// func main() {

    fmt.Println("ans:")

    arr := []int{1,1,1,2,2,3}
    k := 2

    ans := topKFrequent(arr, k)

    fmt.Println(ans)

}
