package main

import (
    "fmt"
)

func main()  {
    fmt.Println(predictPartyVictory("DR"))
    fmt.Println(predictPartyVictory("RDD"))
}

func predictPartyVictory(senate string) string {
    sb := []byte(senate)
    cr := 0 // 天辉已经遍历的数量
    cd := 0 // 夜宴已经遍历的数量

    for ;; {
        end := 0
        for i := 0; i < len(sb); i++ {
            if sb[i] == 'R' {
                if cd > 0 {
                    cd--
                } else {
                    sb[end] = sb[i]
                    end++
                    cr++
                }
            }

            if sb[i] == 'D' {
                if cr > 0 {
                    cr--
                } else {
                    sb[end] = sb[i]
                    end++
                    cd++
                }
            }
        }

        if end == len(sb) {
            break
        }

        sb = sb[:end]
    }

    if sb[0] == 'D' {
        return "Dire"
    }

    return "Radiant"
}
