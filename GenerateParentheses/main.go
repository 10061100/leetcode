// GenerateParentheses
package main

import "fmt"
import "strings"

func main()  {

    fmt.Println(strings.Join(generateParenthesis(0), ","))
}

var lresult []string

func generateParenthesis(n int) []string {
    lresult = []string{}
    r := make([]string, n)
    str := make([]rune, n * 2)
    for i := 0; i < n * 2; i++ {
        str[i] = ' '
    }
    c := 0

    create(&r, n, n, 0, str, &c)

    // println(c)

    return lresult
}


func create(result *[]string, left int, right int, cur int, str []rune, c *int) {
    if( left == 0 && right == 0 ){
        // println(string(str))
        *c++
        // print(*c)
        lresult = append(lresult, string(str))
        return
    }

    if( left == right ){
        str[cur] = '('
        create(result, left-1, right, cur+1, str, c)
    } else {
        if left != 0 {
            str[cur] = '('
            create(result, left-1, right, cur+1, str, c)
        }

        if right > left {
            str[cur] = ')'
            create(result, left, right-1, cur+1, str, c)
        }
    }

}