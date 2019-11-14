package main


func smallestSubsequence(text string) string {

    if text == "" {
        return text
    }

    // 统计每个字符出现的数量
    c := [26]int{}
    for i := 0; i < 26; i++ {
        c[i] = -1
    }
    for k, v := range text {
        c[int(v-'a')] = k
    }

    r := make([]byte, 0)
    for _ , v := range []byte(text) {
        t := int(v - 'a')

        // 如果只出现一次, 肯定要在序列里面的
        if c[t] == 1 {
            r = append(r, v)
            c[t] = -1 // 设置为已经访问了
            continue
        }

        // 如果不止出现一次, 就需要看右边有没有比自己小的, 如果有比自己晓得且可以选


    }

    return string(r)
}
