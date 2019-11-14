package main

import (
    "fmt"
)

func main() {
    fmt.Println(isMatch("heeelllooo", "hello") == true)
    fmt.Println(isMatch("abc", "abcd") == false)
    fmt.Println(isMatch("abcd", "abc") == false)
    fmt.Println(isMatch("heeellloo", "helo") == false)
    fmt.Println(isMatch("lee", "le") == false)
    fmt.Println(isMatch("aaa", "aaaa") == false)
    fmt.Println(isMatch("aaaa", "aaa") == true)
}


func expressiveWords(S string, words []string) int {
    c := 0

    for _, v :=range words {
        if isMatch(S, v) {
            c++
        }
    }

    return c
}


func isMatch(s string, word string) bool {
    if s == word {
        return true
    }

    if s == "" {
        return false
    }

    if word == "" {
        return false
    }

    i, j := 0, 0
    for ; i < len(s) && j < len(word) ; {
        if s[i] != word[j] {
            return false
        }

        // 如果s[i] == word[j], 分别往后跳过所有的相同字符
        c1, c2 := 1,1
        for i = i+1; i < len(s) && s[i] == s[i-1]; i++ {c1++}
        for j = j+1; j < len(word) && word[j] == word[j-1]; j++ {c2++}

        // fmt.Println(c1)
        // fmt.Println(c2)
        if c2 > c1 {
            return false
        }

        if c1 != c2 && c1 <= 2 {
            return false
        }
    }

    // fmt.Println(i, ",", j)
    return i >= len(s)  && j >= len(word)
}
