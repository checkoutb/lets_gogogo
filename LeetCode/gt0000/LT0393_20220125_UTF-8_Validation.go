// package sdq
package main

import (
    "fmt"
)


// if((data[i] & 128) == 0) { // 0xxxxxxx, 1 byte, 128(10000000)
//     numberOfBytes = 1;
// } else if((data[i] & 224) == 192) { // 110xxxxx, 2 bytes, 224(11100000), 192(11000000)
//     numberOfBytes = 2;
// } else if((data[i] & 240) == 224) { // 1110xxxx, 3 bytes, 240(11110000), 224(11100000)
//     numberOfBytes = 3;
// } else if((data[i] & 248) == 240) { // 11110xxx, 4 bytes, 248(11111000), 240(11110000)
//     numberOfBytes = 4;
// if((data[i+j] & 192) != 128) return false; // 192(11000000), 128(10000000)

// if ((c >> 5) == 0b110) count = 1;
// else if ((c >> 4) == 0b1110) count = 2;
// else if ((c >> 3) == 0b11110) count = 3;
// else if ((c >> 7)) return false;
// } else {
// if ((c >> 6) != 0b10) return false;


// int a1 = 128;
// int a11 = 128 + 64;
// int a111 = 128 + 64 + 32;
// int a1111 = 128 + 64 + 32 + 16;
// int a11111 = a1111 + 8;

// int mask1 = 0b10000000, 
// mask2 = 0b11111,  val2 = 0b11011111, 
// mask3 = 0b1111,   val3 = 0b11101111, 

// } else if (!((data[i] >> 4) ^ 14)) { 

// const b1_m, b1_v = 0b1000_0000, 0b0000_0000

// byteData := fmt.Sprintf("%08b", data[i])
// if numBytes == 0 {
//     if string(byteData[0]) == "0" {
//         continue
//     }
//     if string(byteData[0:3]) == "110" {
//         numBytes = 1
//     } else if string(byteData[0:4]) == "1110" {
//         numBytes = 2


// Runtime: 14 ms, faster than 37.50% of Go online submissions for UTF-8 Validation.
// Memory Usage: 5.4 MB, less than 37.50% of Go online submissions for UTF-8 Validation.
// 1byte, 首位0
// 2byte，首位110
// 3byte，开头1110
// 4byte，开头11110
// 2,3,4 的 后续 1,2,3个 都是 10.
// 0xxx x xxx           < 8
// 110x x xxx        >=12 <=13
// 1110 x xxx        14
// 1111 0 xxx        15
// 10xx x xxx     >=8 <= 11    << 4
func validUtf8(data []int) bool {
    // const b10 = 0x80         // 1000 0000
    // const arr = []int{0x80, 0xC0, 0xE0, 0xF0}      // 10000.. 110000.. 111000.. 1111000...
    const b1111 = 0xF0          // 1111
    cnt := 0        // skip 10xx xxxx
    for i := 0; i < len(data); i++ {
        t2 := b1111 & data[i]
        if cnt > 0 {
            if t2 < 128 || t2 > 176 {
                fmt.Println(i, t2)
                return false
            }
            cnt--
            continue
        }
        if t2 < 128 {
            cnt = 0
        } else if t2 >= 192 && t2 <= 208 {
            cnt = 1
        } else if t2 == 224 {
            cnt = 2
        } else if t2 == 240 {
            if data[i] & (1<<3) != 0 {          // xxxx 0xxx
                return false
            }
            cnt = 3
        } else {
            fmt.Println(i)
            return false
        }
    }
    fmt.Println(cnt)
    return cnt == 0
}

func main_LT0393_20220125() {
// func main() {

    fmt.Println("ans:")

    // fmt.Println(8 << 4)         // 128
    // fmt.Println(11 << 4)        // 176
    // fmt.Println(12 << 4)        // 192
    // fmt.Println(13 << 4)        // 208
    // fmt.Println(14 << 4)        // 224
    // fmt.Println(15 << 4)        // 240

    // arr := []int{197,130,1}
    // arr := []int{235,140,4}
    arr := []int{248,130,130,130}
    // arr := []int{228,189,160,229,165,189,13,10}

    for _, v := range arr {
        fmt.Printf("%b, ", v)
    }

    fmt.Println(validUtf8(arr))

}
