package main

import (
    "fmt"
)

func main() {
    fmt.Println(isRobotBounded("GGLLGG") == true)
    fmt.Println(isRobotBounded("GG") == false)
    fmt.Println(isRobotBounded("GL") == true)
    fmt.Println(isRobotBounded("GLRLLGLL") == true)
}

// 因为是循环跑, 所以如果保证在一个圈子里就必须保证要不一次循环回到原点, 要不就方向不一致,
// 否则会冲着一个方向出去
func isRobotBounded(instructions string) bool {

    r := &Robot{
        X: 0,
        Y: 0,
        To: 0,
    }

    for _, v := range []byte(instructions) {
        r.Exe(v)
        // fmt.Println(r)
    }

    if r.X == r.Y && r.X == 0 {
        return true
    }

    if r.To != 0 {
        return true
    }

    return false
}


type Robot struct {
    X int
    Y int
    To int
}

func (this *Robot) Exe(cmd byte) *Robot {
    if cmd == 'L' {
        return this.Turn(-1)
    }

    if cmd == 'R' {
        return this.Turn(1)
    }

    return this.Go()
}

func (this *Robot) Go() *Robot {
    // 南北方向变动y
    if this.To%2 == 0 {
        if this.To - 2 == 0 {
            this.Y -= 1
        } else {
            this.Y += 1
        }
        return this
    }

    // 东西方向变动x

    if this.To - 2 > 0 {
        this.X -= 1
    } else {
        this.X += 1
    }

    return this
}

func (this *Robot) Turn(dir int) *Robot {
    if dir > 0 {
        this.To = (this.To + 1) % 4
    } else {
        this.To = (this.To - 1 + 4) % 4
    }

    return this
}
