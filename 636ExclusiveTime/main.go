package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // fmt.Println(exclusiveTime(2, []string{"0:start:0","0:start:2","0:end:5","1:start:6","1:end:6","0:end:7"}))
    // fmt.Println(exclusiveTime(1, []string{"0:start:0","0:start:2","0:end:5","0:start:6","0:end:6","0:end:7"}))
    // fmt.Println(exclusiveTime(2, []string{"0:start:0","0:start:2","0:end:5","1:start:7","1:end:7","0:end:8"}))
    fmt.Println(exclusiveTime(8, []string{"0:start:0","1:start:5","2:start:6","3:start:9","4:start:11","5:start:12","6:start:14","7:start:15","1:start:24","1:end:29","7:end:34","6:end:37","5:end:39","4:end:40","3:end:45","0:start:49","0:end:54","5:start:55","5:end:59","4:start:63","4:end:66","2:start:69","2:end:70","2:start:74","6:start:78","0:start:79","0:end:80","6:end:85","1:start:89","1:end:93","2:end:96","2:end:100","1:end:102","2:start:105","2:end:109","0:end:114"}))
}

func exclusiveTime(n int, logs []string) []int {
    res := make([]int, n)
    // timestack := NewStack(len(logs))
    idstack:= NewStack(len(logs))
    if len(logs) == 0 {
        return res
    }
    id, _, _ := formatLog(logs[0])
    idstack.Push(id)
    for i := 1; i < len(logs); i++ {

        curid, curmode, curtime := formatLog(logs[i])
        _, _, pretime := formatLog(logs[i-1])


        if curmode == "start" {
            if !idstack.IsEmpty() {
                res[idstack.Top()] += curtime-pretime
            }
            idstack.Push(curid)
        } else {
            res[curid] += curtime - pretime
            idstack.Pop()
        }

        // fmt.Println(res)
    }

    return res
}

func formatLog(s string) (id int, mode string, time int) {
    t := strings.Split(s, ":")
    id, _ = strconv.Atoi(t[0])
    mode  = t[1]
    time, _= strconv.Atoi(t[2])

    if mode == "end" {
        time++
    }

    return
}


type Stack struct {
    stack []int
    top int
    size int
}

func NewStack(n int) *Stack {
    s := make([]int, n)

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}


func ( s *Stack) Top() int {
    if s.top == 0 {
        return 0
    }
    return s.stack[s.top-1]
}

func (s * Stack ) Pop() int {
    if s.top == 0 {
        return 0
    }

    s.top--
    r := s.stack[s.top]

    return r
}


func (s *Stack) Push(x int) *Stack {
    if s.top >= s.size {
        s.stack = append(s.stack, x)
    }

    s.stack[s.top] = x
    s.top++

    return s
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}

func (s * Stack) Len() int {
    return s.top
}
