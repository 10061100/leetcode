package main

import (
    "fmt"
)

func main() {
    fmt.Println(letterCombinations("234"))
    fmt.Println(letterCombinations(""))
    fmt.Println(letterCombinations("2"))
}


var dic = map[int][]byte {
    2:{'a','b','c'},
    3:{'d','e','f'},
    4:{'g','h','i'},
    5:{'j','k' ,'l'},
    6:{'m','n','o'},
    7:{'p','q','r', 's'},
    8:{'t','u','v'},
    9:{'w','x','y', 'z'},
}

func letterCombinations(digits string) []string {


    dbs := []byte(digits)
    ldbs:= len(dbs)
    res := make([]string, 0)
    if ldbs == 0 {
        return res
    }

    tmp := make([]byte, ldbs)

    leeterSubOrder(dbs, ldbs, 0, tmp, &res)

    return res
}

func leeterSubOrder(dbs []byte, l, cur int, tmp []byte, res *[]string) {
    if cur >= l {
        *res = append(*res, string(tmp))
        return
    }

    num := int(dbs[cur] - '0')
    m := dic[num]

    for _, v := range m {
        tmp[cur] = v
        leeterSubOrder(dbs, l, cur+1, tmp, res)
    }

}
