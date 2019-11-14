// MaxAreaOfIsland
package main

func main() {
    println(maxAreaOfIsland([][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1}}))
}

func maxAreaOfIsland(grid [][]int) int {

    h := len(grid)
    if h == 0 {
        return 0
    }

    w := len(grid[0])

    if w == 0 {
        return 0
    }

    gmap := make([][]bool, h )
    for i := 0; i < h ; i++ {
        t := make([]bool, w)
        for j := 0; j < w; j++ {
            t[j] = false
        }
        gmap[i] = t
    }

    max := 0
    for i := 0; i< h; i++ {
        for j := 0; j < w; j++ {
            if gmap[i][j] == true {
                continue
            }

            if grid[i][j] == 1 {
               if area := countLand(grid, gmap, w,h, j, i); area > max {
                   max = area
               }
            }
        }
    }

    return max
}

func countLand(grid [][]int, gmap [][]bool, w,h int, x, y int) int {
    if x >= w || y >= h {
        return 0
    }

    q := NewQueue()
    area := 0

    if grid[y][x] > 0 {
        q.Push(&TreeNodeJ{J:x, I:y})
        gmap[y][x] = true
    }

    for ; !q.IsEmpty(); {
        t := q.Out()
        if t == nil {
            break
        }

        if grid[t.I][t.J] > 0 {
            area++
            x = t.J
            y = t.I
            if y-1 >= 0 && grid[y-1][x] == 1 && gmap[y-1][x] == false {
                gmap[y-1][x] = true
                q.Push(&TreeNodeJ{J:x, I:y-1})
            }
            if y+1 < h && grid[y+1][x] == 1 && gmap[y+1][x] == false {
                gmap[y+1][x] = true
                q.Push(&TreeNodeJ{J:x, I:y+1})
            }
            if x-1 >= 0 && grid[y][x-1] == 1 && gmap[y][x-1] == false {
                gmap[y][x-1] = true
                q.Push(&TreeNodeJ{J:x-1, I:y})
            }
            if x+1 < w && grid[y][x+1] == 1 && gmap[y][x+1] == false {
                gmap[y][x+1] = true
                q.Push(&TreeNodeJ{J:x+1, I:y})
            }
        }

    }

    return area
}



type Queue struct {
    C []*TreeNodeJ
    N int
}

func NewQueue() *Queue {
    return &Queue{
        C: make([]*TreeNodeJ, 0),
        N: 0,
    }
}

func (q *Queue) Push(t *TreeNodeJ){
    q.C = append(q.C, t)
    q.N++
}

func (q *Queue) IsEmpty() bool {
    return q.N == 0
}

func (q *Queue) Out() *TreeNodeJ {
    if q.IsEmpty() {
        return nil
    }

    q.N--
    t := q.C[0]
    q.C = q.C[1:]
    return t
}

type TreeNodeJ struct {
    J, I int
}
