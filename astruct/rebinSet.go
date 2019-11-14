package astruct

// 并查集

type RebinSet struct {
    M map[string]string
}


func NewRebinSet() *RebinSet {
    return &RebinSet{
        M: make(map[string]string),
    }
}


func (r *RebinSet) GetRoot(s string) string {

    if _, ok := r.M[s]; !ok {
        return s // 没有的话,  就是他自己
    }

    for ; ; {
        if p, ok := r.M[s]; !ok || p == "" {
            return s
        } else {
            s = p
        }
    }
}

func (r *RebinSet) Set(parent, s string) {

    if !r.Contain(s) {
        r.Add(s)
    }

    p := r.GetRoot(parent)

    sp := r.GetRoot(s)

    if p > sp {
        r.M[sp] = p
    } else if p < sp {
        r.M[p] = sp
    }

}

func (r *RebinSet ) Contain(s string) bool {
    if _, ok := r.M[s]; !ok {
        return false
    }

    return true
}

func (r *RebinSet) Add(s string) {
    r.M[s] = "" // 自己是自己的祖先
}

func (r *RebinSet) OutPut() {
    for  k, v := range r.M {
        println(k, "=>", v)
    }
}
