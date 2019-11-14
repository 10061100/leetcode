// IsOneBitCharacter
package main

func main() {

    println(isOneBitCharacter([]int{1, 0, 0}) == true)
    println(isOneBitCharacter([]int{1, 1, 1, 0}) == false)
}


func isOneBitCharacter(bits []int) bool {

    length := len(bits)

    for i := 0 ; i < length; {
        if v := bits[i]; v == 0 {
            i++
            continue
        } else {
            if i == length -1 || i == length - 2{
                return false
            }

            i += 2
        }
    }

    return true
}