package main

import (
    "fmt"
)

func main() {
    fmt.Println(removeOuterParentheses("((()))") == "(())")
    fmt.Println(removeOuterParentheses("(())") == "()")
    fmt.Println(removeOuterParentheses("(()())(())") == "()()()")
    fmt.Println(removeOuterParentheses("(()())(())(()(()))") == "()()()()(())")
    fmt.Println(removeOuterParentheses("()") == "")
}

func removeOuterParentheses(S string) string {
    // bs := []byte(S)
    ls := len(S)

    if ls == 0 {
        return S
    }

    count := 0
    res := make([]byte, 0)
    for i := 0; i < ls; i++ {
        if count == 0 {
            count++
            continue
        }

        v := S[i]
        if v == ')' {
            count--
            if count != 0 {
                res = append(res, v)
            }
        } else {
            count++
            res = append(res, v)
        }
    }


    return string(res)
}


