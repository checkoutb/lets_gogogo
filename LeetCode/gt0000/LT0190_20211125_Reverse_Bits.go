// package sdq
package main

import (
    "fmt"
)




// for(int i = 0; i < 32; i++) {
//     res = (res << 1) + (n & 1);
//     n = n >> 1;
// }


// public int swapBits(int n, int i, int j) {
// 	int a = (n >> i) & 1;
// 	int b = (n >> j) & 1;
// 	if ((a ^ b) != 0) {
// 		return n ^= (1 << i) | (1 << j);
// 	}
// 	return n;
// }


// uint32_t reverseBits(uint32_t n) {
//     n = (n >> 16) | (n << 16);
//     n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8);
//     n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4);
//     n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2);
//     n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1);
//     return n;
// }


// Runtime: 4 ms, faster than 23.57% of Go online submissions for Reverse Bits.
// Memory Usage: 2.8 MB, less than 5.44% of Go online submissions for Reverse Bits.
func reverseBits(num uint32) uint32 {
    var ans uint32 = 0
    for i := 0; i < 32; i++ {
        ans |= (num & (1 << i)) >> i << (31 - i)            // << (31 - i - i)
        // var t1 uint32 = num & (1 << i)
        // var t2 uint32 = num & (1 << (31 - i))
        // ans |= 
    }
    return ans
}


func main_LT0190_20211125() {
// func main() {

    fmt.Println("ans:")


}
