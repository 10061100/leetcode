package main

func main() {
    // print(countDigitOne(12312312))
    print(countDigitOne(13))
}

func countDigitOne(n int) int {
    r := 0
    for i := 1; i < n; i = i * 10 {
        t := i * 10
        r += (n/t) * i
        mod := n % t
        if mod >= i * 2 {
            r += i
        } else if mod >= i {
            r += mod - i + 1
        }
    }

    return r
}


