
func allPathsSourceTarget(graph [][]int) [][]int {
results := make([][]int, 0)
    n := len(graph)
    var dfs func(int, []int)

    dfs = func(node int, path []int) {
        if node == n-1 {
            tmp := make([]int, len(path))
            copy(tmp, path)
            results = append(results, tmp)
            return
        }

        for _, v := range graph[node] {
            path = append(path, v)
            dfs(v, path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, []int{0})
    return results
}
