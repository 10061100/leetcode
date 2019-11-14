package main


func main() {
    println(wordBreak("leetcode", []string{"leet", "code"}))
    println(wordBreak("applepenapple", []string{"apple", "pen"}))
    println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
    wordset := NewSetArray(wordDict)

    bs := []byte(s)
    ls := len(bs)
    dp := make([]bool, ls+1)

    dp[0] = true
    for i := 1; i <= ls; i++ {
        dp[i] = false
        for j := 0 ; j < i && !dp[i]; j++ {
            dp[i] = dp[j] && wordset.IsExist(string(bs[j:i]))
        }
    }

    return dp[ls]
}


type Set struct {
    set map[string]int
}

func NewSetArray(array []string) *Set {

    s := NewSet()

    for _, v := range array {
        s.Set(v)
    }

    return s
}

func NewSet()  *Set {
    return &Set{
        set: make(map[string]int),
    }
}

func (s *Set) Del(key string) *Set{
    delete(s.set, key)
    return s
}

func (s *Set) Set(key string) *Set {
    if _, ok := s.set[key]; !ok {
        s.set[key] = 1
    }

    return s
}


func ( s * Set) IsExist(key string) bool {
    if _, ok := s.set[key]; ok {
        return true
    }

    return false
}
