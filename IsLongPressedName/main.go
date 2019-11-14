package main


func isLongPressedName(name string, typed string) bool {
    bname := []byte(name)
    btyped:= []byte(typed)

    if len(btyped)  < len(bname) {
        return false
    }

    ni,ti := 0, 0
    for ; ti < len(btyped) && ni < len(bname); ti++ {
        if bname[ni] == btyped[ti] {
            ni++
        } else {
            if ti <= 0 {
                return false
            }

            if btyped[ti] != btyped[ti-1] {
                return false
            }
        }
    }

    // 如果 typed遍历完了, 但是 name没有遍历玩
    if ti == len(btyped) && ni < len(bname) {
        return false
    }

    // 如果name遍历完了, 但是typed没有遍历完
    if ni == len(bname) && ti < len(btyped) {
        for ; ti < len(btyped); ti++ {
            if btyped[ti] != btyped[ti-1] {
                return false
            }
        }
    }

    return true
}