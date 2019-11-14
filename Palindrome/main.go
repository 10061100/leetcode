// Palindrome
package main


func main(){
    println(isPalindrome(1002201))
    // println(maxBase(9))
}


func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    if x < 10 {
        return true
    }

    maxBase := maxBase(x)
    minBase := 10

    // print(maxBase)
    for ; ;  {
        maxBaseBit := x/maxBase
        minBaseBit := x % minBase

        if maxBaseBit != minBaseBit {
            return false
        }


        x = x % maxBase
        x = x / minBase

        maxBase = maxBase/100

        if maxBase < minBase {
            return true
        }
    }


    return true
}


func maxBase(x int ) int {
    n := 1
    for ; x > 0; n = n * 10 {
        x = x/10
    }

    return n/10
}
