package main

import (
    "fmt"
)

func main() {
    fmt.Println(validTicTacToe([]string{"XOX", "O O", "XOX"}) == true)
    fmt.Println(validTicTacToe([]string{"XXX", "   ", "OOO"}) == false)
    fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}) == false)
    fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}) == false)
}


func validTicTacToe(board []string) bool {

    // 首先统计x和o的数量
    x , o := 0, 0
    for _, v := range board {
        for _, c := range v {
            if c == 'X' {
                x++
            }
            if c == 'O' {
                o++
            }
        }
    }

    if o > x {
        return false
    }

    if x - o >= 2 {
        return false
    }

    // 接下来统计赢得数量
    win := make(map[uint8]int)
    win["X"[0]] = 0
    win["O"[0]] = 0
    // 横排 + 宗排
    for i := 0; i < 3; i++ {
        if checkWin([][]int{{i, 0}, {i,1}, {i, 2}}, board) {
            win[board[i][0]]++
        }

        if checkWin([][]int{{0, i}, {1,i}, {2, i}}, board) {
            win[board[0][i]]++
        }
    }

    // 斜的

    if checkWin([][]int{{0, 0}, {1,1}, {2, 2}}, board) {
        win[board[1][1]]++
    }

    if checkWin([][]int{{0, 2}, {1,1}, {2, 0}}, board) {
        win[board[1][1]]++
    }

    if x == o {
        // x 必然是没哟赢得
        return  win[uint8('X')] == 0
    }

    if x > o {
        // 这个时候, O一定不是赢得
        return win[uint8('O')] == 0
    }

    return true
}

func checkWin(array [][]int, board []string) bool {
    c := board[array[0][0]][array[0][1]]

    for i := 1; i < 3; i++ {
        if board[array[i][0]][array[i][1]] != c {
            return false
        }
    }

    return true
}
