package main


func main() {
    println(kthGrammar(4,3) == 1)
}


func kthGrammar(N int, K int) int {
    tmp := make([]int, N)

    K--
    for t := N; t > 1 ; t-- {
        if K % 2 == 0 {
            tmp[t-1] = 0
        } else {
            tmp[t-1] = 1
        }

        K /= 2
    }

    t := 0
    var tar []int
    for i := 1; i < N; i++ {
        if t == 0 {
            tar = []int{0,1}
        } else {
            tar = []int{1,0}
        }

        t = tar[tmp[i]]
    }

    return t
}