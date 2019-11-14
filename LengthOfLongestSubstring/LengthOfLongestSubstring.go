// LengthOfLongestSubstring
package main


func main() {
    println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {

    m       := make(map[byte]int)
    bigger1 := 0
    max     := 0
    bs      := []byte(s)

    for f, e := 0,0 ; e < len(bs); e++ {
        t := bs[e]
        if v, ok := m[t]; !ok || v == 0 {
            m[t] = 1
            bigger1 += 1
            if bigger1 > max {
                max = bigger1
            }
        } else {
            m[t]++
            if v == 1 {
                bigger1--
            }
            for ; f <= e && m[t] != 1; f++ {
                // fmt.Printf("%d => %c => %d => %c => %d\n", e, t, m[t], 0, 0)
                tf := bs[f]
                if v, ok := m[tf]; ok {
                    m[tf] = m[tf] - 1
                    if v == 2 {
                        bigger1++
                        if bigger1 > max {
                            max = bigger1
                        }
                    } else if v == 1 {
                        bigger1--
                    }
                }
            }
        }
    }

    return max
}