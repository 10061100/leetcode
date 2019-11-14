// IsToeplitzMatrix
package main


func main() {

}


func isToeplitzMatrix(matrix [][]int) bool {
    h := len(matrix)
    w := len(matrix[0])

    for j := 0; j < h; j ++ {
        if !isSameLine(0, j, w, h, matrix) {
            return false
        }
    }

    for i := 0; i < w ; i++ {
        if !isSameLine(i, 0, w, h, matrix) {
            return false
        }
    }

    return true
}


func isSameLine(x, y int, w, h int, matrix [][]int) bool {

    for i , j := x, y;  i < w && j < h ;  {
        if matrix[y][x] != matrix[j][i] {
            return false
        }
        i++
        j++
    }

    return true
}