// LongestValidParentheses
package main

func main() {
    println(longestValidParentheses(")(((((()())()()))()(()))("))
    println(longestValidParentheses(")()())()()("))
}


func longestValidParentheses(ss string) int {
    bs := []byte(ss)
    l  := len(bs)

    stack := make([]int, l)
    top := 0
    max := 0
    stack[top] = -1
    top++
    for i := 0; i < l; i++ {
        c := bs[i]
        if c == '(' {
            stack[top] = i
            top++
        } else  {
            if top > 0 {
                top--
                // t := stack[top]
            }

            if top == 0 {
                stack[top] = i
                top++
            } else {
                t := stack[top-1]
                if ll := i - t; ll > max {
                    max = ll
                    // println(i,  t, ll)
                }
            }
        }
    }


    return max
}