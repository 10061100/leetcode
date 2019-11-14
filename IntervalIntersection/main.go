package main



type Interval struct {
	   Start int
	   End   int
}


func intervalIntersection(A []Interval, B []Interval) []Interval {
    r := make([]Interval, 0)
    lA := len(A)
    lB := len(B)

    if lA == 0 || lB == 0 {
        return r
    }

    for a, b := 0, 0; a < lA && b < lB; {
        ta, tb := A[a], B[b]
        start := max(ta.Start, tb.Start)
        end   := min(ta.End, tb.End)

        if end >= start {
            r = append(r, Interval{
                Start: start,
                End: end,
            })
        }

        if tb.End < ta.End {
            b++
        } else if tb.End > ta.End {
            a++
        } else {
            a++
            b++
        }
    }

    return r
}

func min( a, b int) int {
    if a < b {
        return a
    }

    return b
}

func max( a, b int) int {
    if a > b {
        return a
    }

    return b
}