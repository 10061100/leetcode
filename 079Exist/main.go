package main


var dir = [][]int{{0,1},{0,-1},{1,0},{-1,0}}


func main() {
    board := [][]byte{
        {'A','B','C','E'},
        {'S','F','C','S'},
        {'A','D','E','E'},
    }

    println(exist(board, "ABCCED"))
    println(exist(board, "SEE"))
    println(exist(board, "SEES"))
    println(exist(board, "ABCEC"))
    println(exist(board, "ABCESEEEFS"))


}


func exist(board [][]byte, word string) bool {
    ws := []byte(word)
    ls := len(ws)

    if ls == 0 {
        return true
    }

    h := len(board)
    if h == 0 {
        return false
    }

    w := len(board[0])
    if w == 0 {
        return false
    }

    for i := 0; i < h; i++ {
        for j := 0; j < w; j++ {
            if c := board[i][j]; c != ws[0] {
                // 第一个字母就不匹配直接失败
                continue
            }

            visited := make([]int, w*h)
            visited[getHash(w, i, j)] = 1
            if recurSearch(board, h, w, i, j, ws, ls, 1, visited) {
                return true
            }
        }
    }

    return false
}


func recurSearch(board [][]byte, height, width int, curh, curw int, word []byte, wl, curl int, visited []int) bool {
    // 如果word已经没有值了, 说明遍历完了
    if curl >= wl {
        return true
    }
    // println(curh, curw, curl)
    // 否则, 看一下上下左右四个值是否可用
    for _, curd := range dir {
        newh := curh + curd[0]
        neww := curw + curd[1]

        if newh >= height || newh < 0 {
            continue
        }

        if neww >= width || neww < 0 {
            continue
        }

        // 判断是否已经进入hash域了
        newhash := getHash(width, newh, neww)
        if visited[newhash] == 1 {
            continue
        }

        newc := word[curl]
        if newc != board[newh][neww] {
            // 当前不等
            continue
        }

        visited[newhash] = 1
        // 后面也相等, 艾玛成功了
        if recurSearch(board, height, width, newh, neww, word, wl, curl+1, visited) == true {
            return true
        }
    }

    return false
}

func getHash(width int, curh, curw int) int {
    return curh * width + curw
}