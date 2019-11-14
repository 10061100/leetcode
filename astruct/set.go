package astruct



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