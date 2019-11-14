package main


func main() {
    println(decodeAtIndex("leet2code3", 10))
}


func decodeAtIndex(S string, K int) string {
    bs := []byte(S)
    ls := len(bs)

    tc := 0
    for i := 0 ; i < ls; i++ {
        if isDigital(bs[i]) {
            tc *= int(bs[i] - '0')
        } else {
            tc++
        }
    }

    for i := ls-1; i >= 0; i-- {
        // 这一段是减少遍历范围
        if isDigital(bs[i]) {
            tc /= int(bs[i] - '0')
            K = K % tc
            continue
        }

        if K == 0 || K == tc {
            return string([]byte{bs[i]})
        } else {
            tc--
        }

    }

    return ""
}

func isDigital(c byte) bool {
    return c >= '0' && c <= '9'
}