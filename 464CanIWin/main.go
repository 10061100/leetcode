package main


func canIWin(maxChoosableInteger int, desiredTotal int) bool {
    t := (maxChoosableInteger + 1)/2 * maxChoosableInteger

    if t <= desiredTotal {
        return false
    }


    return true
}
