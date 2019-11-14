// NextGreatestLetter
package main

func main() {
    letters := []byte{'c', 'f', 'j'}
    println(nextGreatestLetter(letters, 'c'))
    println(nextGreatestLetter(letters, 'f'))
    println(nextGreatestLetter(letters, 'j'))
    println(nextGreatestLetter(letters, 'a'))

}

func nextGreatestLetter(letters []byte, target byte) byte {
    for i := 0; i < len(letters); i++ {
        if v := letters[i]; v > target {
            return v
        }
    }

    return letters[0]
}
