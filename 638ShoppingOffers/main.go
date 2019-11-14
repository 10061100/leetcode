package main

import (
    "fmt"
)


func main() {
    fmt.Println(shoppingOffers([]int{2,5}, [][]int{{3,0,5},{1,2,10}}, []int{3,2}) == 14)
    fmt.Println(shoppingOffers([]int{2,3,4}, [][]int{{1,1,0,4},{2,2,1,9}}, []int{1,2,1}) == 11)
    fmt.Println(shoppingOffers([]int{4,3,2,9,8,8}, [][]int{{1,5,5,1,4,0,18},{3,3,6,6,4,2,32}}, []int{6,5,5,6,4,1}))
    fmt.Println(shoppingOffers([]int{4,3,2,9,8,8}, [][]int{{1,5,5,1,4,0,18},{3,3,6,6,4,2,32}}, []int{1,1,1,1,1,1}))

}


func shoppingOffers(price []int, special [][]int, needs []int) int {
    s := make([]BigPackage, 0)
    for i := 0; i < len(special); i++ {
        s = append(s, BigPackage(special[i]))
    }

    n := Need(needs)

    // fmt.Println(s)
    ret := 0
    for k, v := range needs {
        ret += v * price[k]
    }

    dfs(n, s, price, 0,  &ret)


    return ret
}

type BigPackage []int

func (b BigPackage) GetPrice() int {
    return b[len(b)-1]
}

type Need []int

func( n Need) Enough( b BigPackage) bool {
    for i := 0; i < len(n); i++ {
        if n[i] < b[i] {
            return false
        }
    }

    return true
}

func (n Need) Over() bool {
    for i := 0; i < len(n); i++ {
        if n[i] > 0 {
            return false
        }
    }

    return true
}

func (n Need) Minus(b BigPackage) {
    for i :=0 ; i < len(n); i++ {
        n[i] = n[i] - b[i]
    }
}

func (n Need) Add(b BigPackage) {
    for i := 0 ; i < len(n); i++ {
        n[i] = n[i] + b[i]
    }
}


func dfs(needs Need, special []BigPackage, d []int, value int, min *int) {
    // 如果需求已经满足, 就判断一下
    // c++
    // defer func() {
    //     c--
    // }()
    // fmt.Println(c)
    // fmt.Println(needs)
    if needs.Over() && *min > value {
        *min = value
        return
    }

    if value >= *min{
        return
    }

    has := false
    for i := 0; i < len(special); i++ {
        if !needs.Enough(special[i]) {
            continue
        }


        value += special[i].GetPrice()
        needs.Minus(special[i])
        // fmt.Println(value)
        // fmt.Println(*min)
        // mt.Println(special[i])
        // fmt.Println(needs)
        if value < *min {
            // fmt.Println(2222)
            has = true
            dfs(needs, special,d , value, min)
        }
        value -= special[i].GetPrice()
        needs.Add(special[i])
        // fmt.Println(needs)
    }

    // 如果没有还可以使用的大礼包, 剩下的就全部单买来解决
    if !has {
        // fmt.Println(value)
        // fmt.Println(needs)
        for  i := 0; i < len(needs); i++ {
            value += needs[i] * d[i]
            // needs[i] = 0
        }
        // fmt.Println(value)

        if value < *min {
            // fmt.Println(value)
            // fmt.Println(*min)
            *min = value
        }
    }
}
