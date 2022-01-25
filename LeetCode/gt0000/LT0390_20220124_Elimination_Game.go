// package sdq
package main

import (
    "fmt"
)



// return ((Integer.highestOneBit(n) - 1) & (n | 0x55555555)) + 1;


// return n == 1 ? 1 : 2 * (1 + n / 2 - lastRemaining(n / 2));


// Stack<Integer> stack = new Stack<>();
// while (n > 1) {
//     n /= 2;
//     stack.push(n);
// }
// int result = 1;
// while (!stack.isEmpty()) {
//     result = 2 * (1 + stack.pop() - result);
// }
// return result;


// boolean left = true;
// int remaining = n;
// int step = 1;
// int head = 1;
// while (remaining > 1) {
//     if (left || remaining % 2 ==1) {
//         head = head + step;
//     }
//     remaining = remaining / 2;
//     step = step * 2;
//     left = !left;
// }
// return head;





// gg


// 1
// 1
// 1 2
// 2
// 1 2 3
// 2
// 1 2 3 4
// 2 4
// 2
// 1 2 3 4 5
// 2 4
// 2
// 1 2 3 4 5 6
// 2 4 6
// 4
// 1 2 3 4 5 6 7
// 2 4 6
// 4
// 
// 1 2 3 4 5 6 7 8 9
// 2 4 6 8
// 2 6
// 6
// ... dp ?????
// 1个的时候 就是 这个      -> + 1
// 2个的时候 第二个     -> 2
// 3个的时候 第二个
// 4个的时候 第二个
// 需要一个方向。
// 9个数字 ->  等于 4个数字 <- 等于 2个数字->  。。 但我不知道 最后剩的是 2,6 。。。
// 
// 2的倍数， 4的倍数， ？？？
// 
func lastRemaining(n int) int {
    
}


func main_LT0390_20220124() {
// func main() {

    fmt.Println("ans:")


}
