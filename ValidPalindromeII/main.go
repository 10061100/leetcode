// ValidPalindromeII
package main


func main() {
    // println(validPalindrome("cbbcc"))
    println(validPalindrome("eeccccbebaeeabebccceea"))
    // println(validPalindrome("ebcbbececabbacecbbcbe"))
}


func validPalindrome(s string) bool {
    bs := []rune(s)
    cur:= 0
    last := len(bs)-1
    hasDel := false

    return isValidPalindrome(bs, cur, last, hasDel)
}


func isValidPalindrome( bs []rune, cur, last int, hasDel bool) bool {
    for ;; {
        if last < cur {
            return true
        }

        if last - cur == 0 {
            return true
        }

        if last - cur == 1 {
            if bs[last] == bs[cur] {
                return true
            }

            return !hasDel
        }

        if bs[last] == bs[cur] {
            last--
            cur++
            continue
        }

        if hasDel {
            println("diff2:",cur)
            return false
        }

        if bs[last-1] == bs[cur] {
            println("---<", last)
            if isValidPalindrome(bs, cur, last-1, true) == true {
                return true
            }
        }

        if bs[last] == bs[cur+1] {
            println("--->", cur)
            if isValidPalindrome(bs, cur+1, last, true) == true {
                return true
            }
        }

        return false
    }
}
