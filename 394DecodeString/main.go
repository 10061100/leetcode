package main

func main() {
    println(decodeString("abc") == "abc")
    // println(decodeString("a2[bc]") )
    // println(decodeString("3[a]") )
    // println(decodeString("3[a]2[bc]"))
    println(decodeString("3[a2[c]]") == "accaccacc")
    println(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
}


func decodeString(s string) string {
    ws := []byte(s)

    bs, _ := decodeStringSubOrder(ws, 0, 1)
    return string(bs)
}

func decodeStringSubOrder(ws []byte, cur int, count int) (r []byte, c int) {
    bs := make([]byte, 0)
    // fmt.Println(string(ws[cur:]), ",", cur, ",",count)
    // defer func() {
    //     println(string(r))
    // }()

    xc := 0
    for i := cur; i < len(ws); i++ {
        if ws[i] <= '9' && ws[i] >= '0' {
            xc = xc * 10 + int(ws[i] - '0')
        } else if ws[i] == '[' {
            res, newend := decodeStringSubOrder(ws, i+1, xc)
            bs = append(bs, res...)
            i  = newend
            xc = 0
        } else if ws[i] == ']' {
            return multiword(bs, count), i
        } else {
            bs = append(bs, ws[i])
        }
    }

    return multiword(bs, count), len(ws)
}

func multiword(word []byte, multi int) []byte {
    res := make([]byte, 0)

    for i := 1; i <= multi; i++ {
        res = append(res, word...)
    }

    return res
}

