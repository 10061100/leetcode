package main

import (
    "fmt"
)

func main()  {
    fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}) ==3)
    fmt.Println(findLUSlength([]string{"ae", "abc", "abc", "def", "def"}))
    // fmt.Println(isSubSeq("abc",))
}


func findLUSlength(strs []string) int {
    index := -1
    for i := 0; i < len(strs); i++ {
        unique := true
        for j := 0; j < len(strs); j++ {
            if i == j {
                continue
            }

            if isSubSeq(strs[j], strs[i]) {
                unique = false
                break
            }
        }

        // 如果不是唯一的, 跳过去
        if unique == false {
            continue
        }

        if index == -1 || len(strs[i]) > len(strs[index]) {
            index = i
        }
    }

    if index == -1 {
        return -1
    }

    // fmt.Println(index)
    return len(strs[index])
}

func isSubSeq(parent, son string) bool {
    if len(son) > len(parent) {
        return false
    }

    si, pi := 0, 0
    for ; si < len(son) && pi < len(parent); {
        if son[si] == parent[pi] {
            si++
            pi++
            continue
        }

        pi++
    }

    return si >= len(son)
}
