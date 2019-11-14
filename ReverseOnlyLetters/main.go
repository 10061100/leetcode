package main

func main() {
    println(reverseOnlyLetters("a--bs"))
}

func reverseOnlyLetters(S string) string {

    bs := []byte(S)
    ls := len(bs)

    if ls <= 1 {
        return S
    }


    for i, j := 0, ls-1; i <= j; {
        if !isLetter(bs[i]) {
            i++
        } else {
            bs[i], bs[j] = bs[j], bs[i]
            j--
            i++
        }
    }

    return string(bs)

}


func isLetter(b byte) bool {
    return b <= 'z' && b >= 'A'
}