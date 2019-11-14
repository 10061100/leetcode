// RotatedDigits
package main

func main() {

    println(rotatedDigits(1000000000))
    // println(getLessRdCount(16, 2, true, true))
    println("====")
    println(getCount(0, 1000000000))
}


var s1 = []int{1,8}
var s2 = []int{2,5,6,9}

func rotatedDigits(n int) int {
    bit := numBits(n)
    num := 0
    for i := bit - 1; i > 0; i-- {
        num = num + getFullBitCount(i)
    }
    num = num + getLessRdCount(n, bit, false, false)
    return num;
}

func getLessRdCount(n , bc int, canZero, hasFocus bool) int {

    f := n / pow(10, bc-1) // 第一位的数

    s1count := lessCount(f, s1)
    s2count := lessCount(f, s2)

    if canZero && f > 0{
        s1count++
    }

    c := 0
    if ( in(f, s1) || f == 0 ) && bc > 1 {
        c = getLessRdCount(n%pow(10, bc-1), bc - 1, true, hasFocus)
    } else if in(f, s2) && bc > 1 {
        c = getLessRdCount(n%pow(10, bc-1), bc - 1, true, true)
    } else if bc == 1 {
        if in(n, s2) {
            c++
        }

        if hasFocus {
            if canZero && n == 0 {
                c++
            }

            if in(n, s1) {
                c++
            }
        }
    }


    if hasFocus {
        return (s1count + s2count) * pow(7, bc - 1) + c
    } else {
        return (s1count + s2count) * pow(7, bc - 1) - s1count * pow(3, bc - 1) + c
    }
}


func in(n int, array []int) bool {
    for _, v := range array {
        if v == n {
            return true
        }
    }

    return false
}

func lessCount( n int , array []int) int {
    c := 0;
    for _, v := range array {
        if v < n {
            c++
        }
    }

    return c
}

func numBits(n int) int {
    c := 0
    for ; n != 0; c++ {
        n /= 10
    }

    return c
}


func getFullBitCount(c int) int {
    return 6 * pow(7, c-1) - 2 * pow(3,c-1)
}


func pow(base, m int) int {
    result := 1
    for i := 0; i < m ; i++ {
        result *= base
    }

    return result
}

func getCount(begin, last int) int {
    c := 0
    for ; begin <= last; begin++ {
        if isValid(begin) {
            c++
        }
    }

    return c
}


func isValid(N int)  bool{
    validFound := false;
    for ;N > 0; {
        switch N%10 {
        case 2,6,5,9:
            validFound = true
            break
        case 3,4,7:
            return false;
        }
        N = N / 10;
    }
    return validFound;
}