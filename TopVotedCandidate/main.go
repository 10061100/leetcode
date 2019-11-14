package main

type TopVotedCandidate struct {
    times []int
    maxVots map[int]int // 某个时刻最多任务的数量
    maxtime int
}


func Constructor(persons []int, times []int) TopVotedCandidate {
    r := TopVotedCandidate{
        times: times,
        maxVots: make(map[int]int),
        maxtime: times[len(times)-1],
    }

    mc := make(map[int]int) // 当前每个候选人的数量
    curp := 0 // 当前是哪个候选人
    maxvotes := 0
    for k, v := range times {
        if _, ok := mc[persons[k]]; !ok {
            mc[persons[k]] = 0
        }
        mc[persons[k]]++

        if mc[persons[k]] >= maxvotes {
            maxvotes = mc[persons[k]]
            curp = persons[k]
        }

        r.maxVots[v] = curp
    }

    return r
}


func (this *TopVotedCandidate) Q(t int) int {

    index := binarySearch(this.times, t)
    ctime := this.times[index]

    return this.maxVots[ctime]
}

func main() {
    println(binarySearch([]int{1,3,5,7,9}, 10))
}

// 返回第一个比target大的数
func binarySearch(arr []int, target int) int {
    start, end := 0, len(arr) - 1
    mid := 0
    for ; start <= end; {
        mid = (start+end)/2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            start = mid + 1
        } else {
            end = mid - 1
        }
    }

    return end
}
