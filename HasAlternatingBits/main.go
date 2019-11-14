// HasAlternatingBits
package main

func main() {
    println(hasAlternatingBits(5) == true)
    println(hasAlternatingBits(6) == false)
}


func hasAlternatingBits(n int) bool {
    if n == 0 || n == 1 {
        return true
    }

    k := 1
    if isEven(n) {
        k = 0
    }
    n = n >> 1

    for ; n != 0 ; {
        if isEven(n) && k == 1 {
            k = 0
            n = n >> 1
            continue
        } else if !isEven(n) && k == 0 {
            k = 1
            n = n >> 1
            continue
        } else {
            return false
        }
    }

    return true
}

func isEven( n int ) bool {
    return n%2 == 0;
}