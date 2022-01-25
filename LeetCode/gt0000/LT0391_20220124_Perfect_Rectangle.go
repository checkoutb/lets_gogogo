// package sdq
package main

import (
    "fmt"
    "sort"
)


// func isRectangleCover(rectangles [][]int) bool {
// 	sort.Sort(Rectangles(rectangles))
// 	// fmt.Printf("jlog3 rectangles %v \n", rectangles)
// 	l := len(rectangles)
// 	sumArea := 0
// 	farLeft := 10001
// 	farBot := 10001
// 	farRight := -10001
// 	farTop := -10001
// 	for i := 0; i < l; i++ {
// 		sumArea += (rectangles[i][3] - rectangles[i][1]) * (rectangles[i][2] - rectangles[i][0])
// 		if farLeft > rectangles[i][0] {
// 			farLeft = rectangles[i][0]
// 		}
// 		if farBot > rectangles[i][1] {
// 			farBot = rectangles[i][1]
// 		}
// 		if farRight < rectangles[i][2] {
// 			farRight = rectangles[i][2]
// 		}
// 		if farTop < rectangles[i][3] {
// 			farTop = rectangles[i][3]
// 		}
// 	}
// 	// fmt.Printf("jlog1 %d, %d, %d, %d \n", farLeft, farBot, farRight, farTop)
//     if (farRight-farLeft)*(farTop-farBot) != sumArea {
//         return false
//     }
//     for i := 0; i < l; i++ {
//         for j := i + 1; j < l && (rectangles[j][0] < rectangles[i][2]); j++ {
// 			if rectangles[j][1] < rectangles[i][3] && rectangles[j][3] > rectangles[i][1] {
// 	            // fmt.Printf("jlog2 i %d, j %d \n", i, j)
// 				return false
// 			}
// 		}
//     }
// 	return true
// }
// 计算面积，面积没有问题，那么无法cover的时候，必然有空白。 不，这里好像 在找 重复。。





// func isRectangleCover(rects [][]int) bool {
    
//     type pt struct { x, y int }
    
//     var (
//         left, right, top, bottom int
//         area int    
        
//         set = make(map[pt]int)
//     )    
    
//     for i, r := range rects {
//         x, y, xx, yy := r[0], r[1], r[2], r[3]
//         if i == 0 || x < left {
//             left = x
//         }
//         if i == 0 || y < bottom {
//             bottom = y
//         }
//         if i == 0 || xx > right {
//             right = xx
//         }
//         if i == 0 || yy > top {
//             top = yy
//         }
//         area += (xx-x)*(yy-y)
        
//         set[pt{x,y}]++                         // //////////////
//         if set[pt{x,y}] == 2 {
//             delete(set, pt{x,y})
//         }
//         set[pt{x,yy}]++                         // //////////////
//         if set[pt{x,yy}] == 2 {
//             delete(set, pt{x,yy})
//         }
//         set[pt{xx,y}]++                         // //////////////
//         if set[pt{xx,y}] == 2 {
//             delete(set,pt{xx,y})
//         }
//         set[pt{xx,yy}]++                         // //////////////
//         if set[pt{xx,yy}] == 2 {
//             delete(set, pt{xx,yy})
//         }
//     }
    
//     return len(set) == 4 &&
//         set[pt{left, bottom}] == 1 &&
//         set[pt{left, top}] == 1 &&
//         set[pt{right, bottom}] == 1 &&
//         set[pt{right, top}] == 1 &&
//         area == (right-left)*(top-bottom)    
    
// }
// 这个是  每个点 必然 出现2次 或 4次。 (除了4个 顶点)。



// def isRectangleCover(self, rectangles):
//     area = 0
//     corners = set()
//     a, c = lambda: (X - x) * (Y - y), lambda: {(x, y), (x, Y), (X, y), (X, Y)}
//     for x, y, X, Y in rectangles:
//         area += a()
//         corners ^= c()
//     x, y, X, Y = (f(z) for f, z in zip((min, min, max, max), zip(*rectangles)))
//     return area == a() and corners == c()




// Runtime: 116 ms, faster than 80.00% of Go online submissions for Perfect Rectangle.
// Memory Usage: 7.7 MB, less than 60.00% of Go online submissions for Perfect Rectangle.
func isRectangleCover(rectangles [][]int) bool {
    sort.Slice(rectangles, func(i, j int) bool {
        return rectangles[i][1] < rectangles[j][1]
    })

    // fmt.Println(rectangles)

    mnx, mxx := 1000000, -1000000
    mny := 10000000
    for i := 0; i < len(rectangles); i++ {
        if rectangles[i][0] < mnx {
            mnx = rectangles[i][0]
        }
        if rectangles[i][2] > mxx {
            mxx = rectangles[i][2]
        }
        if rectangles[i][1] < mny {
            mny = rectangles[i][1]
        }
    }
    arr := make([]int, mxx - mnx)
    for i, _ := range arr {
        // arr[i] = -12345678
        arr[i] = mny
    }
    for i := 0; i < len(rectangles); i++ {
        x1,y1,x2,y2 := rectangles[i][0] - mnx,rectangles[i][1],rectangles[i][2] - mnx,rectangles[i][3]
        for j := x1; j < x2; j++ {
            // if arr[j] == -12345678 || arr[j] == y1 {
            if arr[j] == y1 {
                arr[j] = y2
            } else {
                return false
            }
        }
    }
    for i := 1; i < len(arr); i++ {
        if arr[i] != arr[i - 1] {
            return false
        }
    }
    return true
}

func main_LT0391_20220124() {
// func main() {

    fmt.Println("ans:")

    // arr := [][]int{{1,1,3,3},{3,1,4,2},{3,2,4,4},{1,3,2,4},{2,3,3,4}}
    // arr := [][]int{{1,1,2,3},{1,3,2,4},{3,1,4,2},{3,2,4,4}}
    // arr := [][]int{{1,1,3,3},{3,1,4,2},{1,3,2,4},{2,2,4,4}}
    // arr := [][]int{{0,0,4,1},{7,0,8,2},{6,2,8,3},{5,1,6,3},{4,0,5,1},{6,0,7,2},{4,2,5,3},{2,1,4,3},{0,1,2,2},{0,2,2,3},{4,1,5,2},{5,0,6,1}}
    arr := [][]int{{0,0,4,1},{7,0,8,2},{6,2,8,3},{5,1,6,3},{6,0,7,2},{4,2,5,3},{2,1,4,3},{0,1,2,2},{0,2,2,3},{4,1,5,2},{5,0,6,1}}

    ans := isRectangleCover(arr)

    fmt.Println(ans)

}
