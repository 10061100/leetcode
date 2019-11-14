// ZigZagConversion
package main

import (
    "fmt"
)

func main(){
    fmt.Println(convert("PAYPALISHIRING", 3))
}


func convert(s string, numRows int) string {
    if "" == s || 0 == numRows {
        return ""
    }

    runeStr := []byte(s)
    strlen := len(runeStr)
    result := make([]byte, strlen)

    for loop := 0 ; loop < numRows ; loop++ {
        for num := 0 ; num * (numRows + 1) + loop < strlen; num ++ {
            result = append(result, runeStr[num * (numRows + 1) + loop])
            if isMidIndex(loop, numRows) && num * (numRows + 1) + numRows < strlen  {
                result = append(result, runeStr[num * (numRows + 1) + numRows])
            }
        }
    }

    return string(result)
}


func isMidIndex(loop int, numRows int ) bool {
    return numRows/2 == loop
}