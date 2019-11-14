package main

import (
    "fmt"
    "strconv"
)

func main() {
    // fmt.Println(fractionToDecimal(1, 411))
    fmt.Println(fractionToDecimal(1, 214748364))
    fmt.Println(fractionToDecimal(1, 30))
    // fmt.Println(fractionToDecimal(41, 200))
}

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    f := ""
    if numerator > 0 && denominator < 0 {
        f = "-"
    }

    if numerator < 0 && denominator > 0 {
        f = "-"
    }

    numerator = abs(numerator)
    denominator = abs(denominator)

    one := numerator/denominator
    two := numerator%denominator

    res := make([]byte, 0)
    c   := make(map[int]int) // 判断当前数字数字是否出现
    for ; two != 0; {

        if v, ok := c[two]; ok && v > 0 {
            break
        }

        t := two * 10

        if t < denominator {
            res = append(res, 0 + '0')
            c[two] = len(res)
            two = t
            continue
        }

        a := t/denominator
        b := t%denominator

        res = append(res, byte(a) + '0')
        c[two] = len(res)
        two = b
    }

    str := make([]byte, 0)
    if f == "-" {
        str = append(str, []byte(f)...)
    }

    str = append(str, []byte(strconv.Itoa(one))...)
    if len(res) != 0 {
        str = append(str, '.')
        if two == 0 {
            str = append(str, res...)
        } else {
            str = append(str, res[0:c[two]-1]...)
            str = append(str, '(')
            str = append(str, res[c[two]-1:]...)
            str = append(str, ')')
        }
    }

    return string(str)
}

func abs(a int) int {
    if a > 0 {return a}
    return -a
}


