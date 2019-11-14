// Rotate
package main

import "fmt"

func main() {

    a := [][]int{
        {1,2,3,4},
        {11,12,13,14},
        {21,22,23,24},
        {31,32,33,34},
    }
    rotate(a)
    fmt.Println(a)
}

func rotate(matrix [][]int)  {
    n := len(matrix)

    for i := 0; i <= (n + 1)/2; i++ {
        start := i
        end   := n - i
        row   := i
        rotateSingle(matrix, row, start, end, n)
    }
}


func rotateSingle(matrix [][]int, row int, start, end int, N int)  {
    if end - start <= 1{
        return
    }

    N = N - 1
    tmp := make([]int, end - start)
    for i := start; i < end; i++ {
        tmp[i-start] = matrix[row][i]
    }

    // 开始进行转职

    // 把最左面一行转到最上边一行
    n := 0 + row
    for i , j:= start, N - row ; i < end ; i++ {
        matrix[n][i] = matrix[j][n];
        j--
    }

    // 把最下面的一行转到最左边一行
    n = N - row
    for i, j := start, start; i < end ; i++ {
        matrix[i][row] =  matrix[n][j]
        j++
    }

    // 把最右面的一行转到最下面一行
    n = N - row
    for i, j := start, N - row; i < end; i++ {
        matrix[n][i] = matrix[j][n]
        j--
    }

    // 吧tmp的值填充到最右面
    n = N - row
    for i, j := start, 0 ; i < end ; i++ {
        matrix[i][n] = tmp[j]
        j++
    }
}