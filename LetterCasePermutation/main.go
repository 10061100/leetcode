// LetterCasePermutation
package main

import "fmt"

func main() {
    fmt.Println(letterCasePermutation("1A2b"))
}

func letterCasePermutation(S string) []string {
    result  := []string{}
    bs      := []byte(S)

    fullLettes(bs, 0, &result)

    return result
}


func fullLettes(s []byte, cur int, r *[]string) {
    if cur >= len(s) {
        *r = append(*r, string(s))
        return
    }

    if isNumber(s[cur]) {
        fullLettes(s, cur+1, r)
        return
    }

    s[cur] = toLower(s[cur])
    fullLettes(s, cur+1, r)
    s[cur] = toUpper(s[cur])
    fullLettes(s, cur+1, r)
}


func isNumber( x byte) bool {
    return x <= '9' && x >= '0'
}


func toUpper(x byte) byte {
    if x <= 'Z' && x >= 'A' {
        return x
    }

    return x - 'a' + 'A'
}

func toLower(x byte) byte {
    if x <= 'z' && x >= 'a' {
        return x
    }

    return x - 'A' + 'a'
}