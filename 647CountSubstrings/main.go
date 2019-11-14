package main

func main()  {

    println(countSubstrings("dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg"))
    println(countSubstrings("abc") == 3)
    println(countSubstrings("aaa") == 6)
}

func countSubstrings(s string) int {
    bs := []byte(s)
    ls := len(bs)

    if ls <= 1 {
        return ls
    }

    news := make([]byte, 2*ls + 1)
    for i := 0; i < ls; i++ {
        news[2*i] = '#'
        news[2*i+1] = bs[i]
    }

    news[2*ls] = '#'

    p := make([]int, 2*ls + 1)
    right := 0
    midle := 0
    count := 0
    p[0] = 1

    for i := 1; i < 2*ls + 1; i++ {
        // fmt.Println("==========")
        k := 2*midle - i
        start := i+1
        if i >= right {
            p[i] = 1
        } else {
            p[i] = p[k]

            kleft := k-p[k]+1
            // fmt.Println("kleft:", kleft, ", k:", k , ", left:", midle - p[midle] + 1)
            if left := midle - p[midle] + 1; left < kleft {
                // fmt.Println("left:", left)
                start = -1
            } else {
                start = right + 1
                p[i] = p[k] - (left - kleft)
            }
        }

        // fmt.Println(start)
        // fmt.Println(p)
        for ; start >= 0 && start <= 2*ls && 2*i - start >= 0 && news[start] == news[2*i-start]; start++ {
            p[i]++
        }
        // fmt.Println(p)
        if newright := i + p[i] - 1; newright > right {
            midle = i
            right = newright
        }

        // k = (p[i]-1)/2
        // count += k
        // if news[i] != '#' {
        //     count++
        // }

        count += p[i]/2

    }

    // fmt.Println(string(news))
    // fmt.Println(p)
    return count
}