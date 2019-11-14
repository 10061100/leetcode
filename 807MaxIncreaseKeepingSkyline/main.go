package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
    y := len(grid)
    x := len(grid[0])

    xmax := make([]int, x)
    ymax := make([]int, y)

    for i := 0; i < y; i++ {
        r := grid[i][0]
        for j := 0; j < x; j++ {
            r = max(grid[i][j], r)
        }
        ymax[i] = r
    }

    for i := 0; i < x; i++ {
        r := grid[0][i]
        for j := 0; j < y; j++ {
            r = max(grid[i][j], r)
        }
        xmax[i] = r
    }

    sum := 0
    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
            m := min(xmax[j], ymax[i])
            if m > grid[i][j] {
                sum += m - grid[i][j]
            }
        }
    }

    return sum
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a , b int) int {
    if a < b {
        return a
    }

    return b
}
