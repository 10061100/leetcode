package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    fmt.Println(math.Sin(math.Pi/4))
}

type Solution struct {
    Radius float64
    X float64
    Y float64
}


func Constructor(radius float64, x_center float64, y_center float64) Solution {
    return Solution{
        Radius: radius,
        X : x_center,
        Y : y_center,
    }
}


func (this *Solution) RandPoint() []float64 {
    the := rand.Float64() * 2
    r := rand.Float64() * this.Radius

    x := this.X + r * math.Cos(the)
    y := this.Y + r * math.Sin(the)

    return []float64{x, y}
}
