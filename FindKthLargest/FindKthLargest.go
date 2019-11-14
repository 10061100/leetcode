// FindKthLargest
package main

// import "fmt"

func main()  {
    println(isAdditiveNumber("123") == true)
    println(isAdditiveNumber("112358") == true)
    println(isAdditiveNumber("199100199") == true)
    println(isAdditiveNumber("1023") == false)
    println(isAdditiveNumber("211738") == true)
    println(isAdditiveNumber("198019823962") == true)
}


func isAdditiveNumber(num string) bool {
    bs := []byte(num)
    a,b := 0,0
    for i := 0; i < len(bs) - 2 ; i++{
        if i > 0 && a == 0 {
            return false
        }
        a = a * 10 + int(bs[i] - '0')
        b = 0
        for j := i + 1; j < len(bs) - 1; j++ {
            if j > i + 1 && b == 0 {
                break
            }
            b = b * 10 + int(bs[j] - '0')
            if isR(a, b, bs, j+1) {
                return true
            }
        }
    }

    return false
}


func isR(a int, b int, s []byte, start int) bool {
    r := 0
    for j := start ; start < len(s); start++ {
        if start > j && r == 0 {
            return false
        }
        r = r * 10 + int(s[start] - '0')
        // println(a, b, r)
        if a + b == r && start == len(s) - 1 {
            return true
        } else if a + b == r && isR(b, r, s, start+1) {
            return true
        }
    }

    return false;
}