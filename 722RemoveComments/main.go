package main


func removeComments(source []string) []string {

    res := make([]string, 0)

    for _, v := range source {
        res = append(res, v)
    }

    return res
}
