// MyAtoi
package main

func main() {
    println(myAtoi("-1"))
}

const INT_MAX = 2147483647;
const INT_MIN = -2147483648;

func myAtoi(str string) int {
    r := 0
    op:= 1
    setOp := false
    delNum:= false

    bs := []byte(str)

    for i := 0; i < len(bs); i++ {
        c := bs[i]
        if r == INT_MAX || r == INT_MIN {
            break
        }

        if !isOp(c) && !isNumber(c) {
            if isWhiteSpace(c) {
                if delNum || setOp {
                    break
                }
            } else {
                break
            }

            continue
        }

        if isOp(c) {
            if setOp {
                return 0
            }

            if isNeg(c) {
                op = -1
                if r > 0 {
                    r = r * -1
                }
            }

            setOp = true
        }


        if isNumber(c) {
            delNum = true
            n := int(c - '0')
            if r > INT_MAX/10 || (r == INT_MAX/10 && n > 7) {
                r = INT_MAX
            } else if r < INT_MIN/10 || ( r == INT_MIN/10 && n > 8 ){
                r = INT_MIN
            } else {
                r = r * 10
                if op == -1 {
                    r = r - n
                } else {
                    r = r + n
                }
            }
        }

    }

    return r
}

func isWhiteSpace(c byte) bool {
    return c == ' ' || c == '\t'
}

func isNeg(c byte) bool {
    return c == '-'
}

func isOp(c byte) bool {
    return c == '-' || c == '+'
}


func isNumber(c byte) bool {
    return c >= '0' && c <= '9'
}