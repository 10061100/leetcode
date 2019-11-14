package main


func main() {
    println(canFinish(2, [][]int{{1,0}}) == true)
    println(canFinish(2, [][]int{{1,0}, {0,1}}) == false)
    println(canFinish(3, [][]int{{1,0}, {0,1}, {1,2}}) == false)
    println(canFinish(4, [][]int{{1,0}, {0,1}, {1,2}, {2,3}, {1,3}}) == false)
    println(canFinish(4, [][]int{{1,0}, {1,2}, {2,3}, {1,3}}) == true)
}


func canFinish(numCourses int, prerequisites [][]int) bool {
    m := make(map[int][]int)

    for i := 0; i < len(prerequisites); i++ {
        pre, next := prerequisites[i][0], prerequisites[i][1]

        if _, ok := m[pre]; !ok {
            m[pre] = make([]int, 0)
        }

        m[pre] = append(m[pre], next)
    }

    visited := make([]int, numCourses)
    for i := 0; i < numCourses; i++ {
        if visited[i] != 0 {
            continue
        }

        if dfs(m, visited, i) {
            return false
        }
    }

    return true
}

func dfs(m map[int][]int, visited []int, start int) bool {

    // 当前节点已经访问了, 直接返回状态
    if visited[start] == 2 {
        return false
    }

    if _, ok := m[start]; !ok {
        visited[start] = 2
        return false // 没有环
    }

    visited[start] = 1 // 表示当前正在被访问
    queue := m[start]

    for _, next := range queue {
        if visited[next] == 2 {
            // 如果当前节点已经访问, 且后面没有环, 进行下一个遍历
            continue
        } else if visited[next] == 1 {
            // 存在环, 直接返回成功
            return true
        } else if dfs(m, visited, next) {
            return true
        } else {
            // 便利了之后,并没有循环,则当前节点也没有
        }
    }

    visited[start]=2

    return false
}