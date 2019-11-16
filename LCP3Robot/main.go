package main

import (
    "fmt"
)

func main()  {
    fmt.Println(robot("URR", [][]int{}, 3, 2) == true)
    fmt.Println(robot("URR", [][]int{{2, 2}}, 3, 2) == false)
    fmt.Println(robot("URR", [][]int{{4, 2}}, 3, 2) == true)
}


func robot(command string, obstacles [][]int, x int, y int) bool {
    m := NewM()
    for _ ,v := range obstacles {
        m.Set(v[0], v[1])
    }

    n := len(command)
    curx, cury := 0, 0
    for i := 0 ; curx <= x && cury <= y; i++ {

        if m.IsExist(curx, cury) {
            return false
        }

        if curx == x && cury == y {
            return true
        }

        cmd := command[i%n]
        if cmd == 'U' {
            cury++
        } else {
            curx++
        }

        // fmt.Println(curx, cury)
    }

    // fmt.Println(curx)
    // fmt.Println(cury)
    return false
}


func NewM() *M{
    return &M{
        m: make(map[int]map[int]bool),
    }
}

type M struct {
    m map[int]map[int]bool
}

func (this *M) Set(x, y int)  {
    if _, ok := this.m[x]; !ok {
        this.m[x] = make(map[int]bool)
    }
    this.m[x][y] = true
}

func (this *M) IsExist(x, y int) bool {
    if _, ok := this.m[x]; !ok {
        return false
    }

    if _, ok := this.m[x][y]; !ok {
        return false
    }

    return true
}