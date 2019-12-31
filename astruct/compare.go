package astruct

func compareArray(a []interface{}, b []interface{}) bool {
    if len(a) != len(b) {
        return false
    }

    for k, v := range a {
        if v != b[k] {
            return false
        }
    }

    return true
}

