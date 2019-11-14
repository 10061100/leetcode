// countPrimeSetBits
package main


func main() {

}

func countPrimeSetBits(L int, R int) int {
    n := 0
    for i := L; i <= R; i++ {
        if c := getCount(i); isPrime(c) {
            n++
        }
    }

    return n
}


func getCount( n int) int {
    c := 0
    for ; n != 0 ; {
        n = n & (n-1)
        c++
    }

    return c
}

func isPrime(n int) bool{
    switch n {
    case 2,3,5,7,11,13,17,19,23:
        return true
    default:
        return false
    }
}