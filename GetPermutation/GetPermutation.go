// GetPermutation
package main

func main() {
    println(getPermutation(3, 4))
}

func getPermutation(n int, k int) string {
    // nt = (n-1)!
    nt := 1
    for i := 1; i< n; i++{
        nt = nt * i
    }

    str := make([]int, 0)
    ex  := make(map[int]int)
    for i := n-1; i >= 0 && k != 0; i--{
        index := k/nt + 1
        t := k % nt
        if t == 0 && k != 0 {
            index--
        }
        k = t
        if i > 0 {
            nt = nt/i
        }
        // 从小到大找到底index个顺序且没有出现的的字符
        c := 0
        // fmt.Println("map:", ex)
        for c = 1; c <= n && index > 0; c++ {
            if _, ok := ex[c]; ok {
                continue
            }
            index--
        }
        str = append(str, c-1)
        ex[c-1] = 1
    }

    for c := n ; c >= 1; c-- {
        if _, ok := ex[c]; ok {
            continue
        }
        str = append(str, c)
    }

    rs := make([]byte, 0)
    for i := 0; i < n ; i++ {
        rs = append(rs, '0' + byte(str[i]))
    }


    return string(rs)
}