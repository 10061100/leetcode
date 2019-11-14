package main

import "fmt"

func main() {
    fmt.Println(splitIntoFibonacci("123456579"))
    fmt.Println(splitIntoFibonacci("3611537383985343591834441270352104793375145479938855071433500231900737525076071514982402115895535257195564161509167334647108949738176284385285234123461518508746752631120827113919550237703163294909"))
}

func splitIntoFibonacci(S string) []int {
    bs := []byte(S)
    ls := len(bs)

    r := make([]int, ls)

    if ls <= 2 {
        return r
    }

    max := 1 << 31 - 1

    t := fibonacci(bs, 0, ls, r, 0, max)
    // fmt.Println(t)
    if t < 2 {
        return r[0:0]
    }

    return r[0:t+1]
}


func fibonacci(bs []byte, start, l int, r []int, cur int , max int) int {

    if start >= l {
        return cur-1
    }

    r[cur] = 0
    for i := start ; i < l; i++ {
        r[cur] = r[cur] * 10 + int(bs[i] - '0')
        if r[cur] > max {
            return -1
        }
        // fmt.Println(r)


        if cur >= 2 {
            if r[cur-2] + r[cur-1] < r[cur] {
                // 如果cur已经很大了, 说明不可能了
                return -1
            } else if r[cur-2] + r[cur-1] == r [cur] {
                // 如果两个相等的话, 说明有可能
                // fmt.Println("xxxx", r[0:cur+1], i, l)
                if i == l - 1 {
                    return cur
                }

                // 说明后面可能还有数, 判断后面是不是
                if t := fibonacci(bs, i+1, l, r, cur+1, max); t > 0{
                    return t
                }

                // 如果不是的话, 就返回遍历再走一遍
            }
        } else {
            if i == l - 1 {
                return -1
            } else {
                if t := fibonacci(bs, i+1, l, r, cur+1, max); t > 0 {
                    return t
                }
            }
        }

        if i == start && bs[start] == '0' {
            break
        }
    }

    return -1;

}