// CanPlaceFlowers
package main

import "fmt"

func main() {
    // println(canPlaceFlowers([]int{1,0,0,0,1}, 2) == false)
    // println(canPlaceFlowers([]int{1,0,0,0,1}, 1) == true)
    // println(pacificAtlantic([][]int{{1,1},{1,1},{1,1}}))
    // println(validUtf8([]int{255}))
    printByte([]int{250,145,145,145,145})
}

func canPlaceFlowers(flowerbed []int, n int) bool {

    for i := 0; i < len(flowerbed); i++ {
        if posCanPlaceFlower(flowerbed, i) {
            flowerbed[i] = 1
            n--
        }
    }

    a := 1000000000+7;
    println(a)

    return n <= 0
}


func posCanPlaceFlower(f []int , pos int) bool {
    if f[pos] != 0 {
        return false
    }

    if (pos <= 0 || f[pos-1]  == 0) && (pos >= len(f)-1 || f[pos+1] == 0) {
        return true
    }

    return false
}


func pacificAtlantic(matrix [][]int) [][]int {
    r := make([][]int, 0)
    m := len(matrix)
    if m <= 0 {
        return r
    }

    n := len(matrix[0])
    if n <= 0 {
        return r
    }

    t := make([][]int , m)
    for i := 0; i < m; i++ {
        t[i] = make([]int, n)
    }

    // 初始化寻找可以流入太平洋的
    // 初始化可以流入大西洋的
    for i := 0; i < n; i++ {
        t[0][i] |= 1
        areaState(matrix, t, 0, i, m, n, 1)
        t[m-1][i] |= 2
        areaState(matrix, t, m-1, i, m, n, 2)
    }

    for i := 0; i < m; i++{
        t[i][0] |= 1
        areaState(matrix, t, i, 0, m, n, 1)
        t[i][n-1] |= 2
        areaState(matrix, t, i, m-1, m, n, 2)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n ; j++ {
            if t[i][j] == 3 {
                r = append(r, []int{i,j})
            }
        }
    }

    return r
}




func areaState(matrix, cache [][]int, x, y int, h, z int, v int) {
    // 逆向感染上下左右

    q := matrix[x][y]
    // 上
    if x >= 1 {
        infect(matrix, cache, x-1, y, h, z, q, v)
    }

    // 下
    if x < h - 1 {
        infect(matrix, cache, x+1, y, h, z, q, v)
    }

    // 左
    if y >= 1 {
        infect(matrix, cache, x, y-1, h, z, q, v)
    }

    // 右边
    if y < z - 1 {
        infect(matrix, cache, x, y+1, h, z, q, v)
    }
}


func infect(matrix, cache [][]int, x, y int, h, z int, q , v int ){
    t := matrix[x][y]

    // 说明无法感染
    if q > t {
        return
    }

    sc := cache[x][y]
    if sc & v == v {
        // 说明已经被感染过了, 直接返回
        return
    }

    cache[x][y] |= v
    areaState(matrix, cache, x, y, h, z, v)
}


var nums []int = []int{128, 64, 32, 16, 8, 4, 2, 1}

func validUtf8(data []int) bool {
    l := len(data)
    if l == 0 {
        return true
    }

    for i := 0; i < l; {
        v := data[i]
        c := bitCount(v)

        // fmt.Println(l, i, c)
        // 没有是1个的
        if c == 1 {
            return false
        }

        i++
        if c == 0 {
            continue
        }

        c--
        // 否则是多个字节的
        if l - i < c {
            return false
        }

        for t := 0; t < c; t++ {
            if !isSubBit(data[i]) {
                return false
            }
            i++
        }
    }

    return true
}


func bitCount(n int) int {
    c := 0
    for ; c < 8 && nums[c] & n > 0 ; c++ {}
    return c
}

func isSubBit(n int) bool {
    return n & nums[0] > 0 && n & nums[1] == 0
}


func printByte(a []int) {
    for _, v := range a {
        fmt.Printf("%b\n", v)
    }
}