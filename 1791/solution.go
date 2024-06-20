func findCenter(edges [][]int) int {
    numMap := make(map[int]struct{})
    for _, edge := range edges {
        edge1 := edge[0]
        edge2 := edge[1]
        if _, exists := numMap[edge1]; exists {
            return edge1
        }
        if _, exists := numMap[edge2]; exists {
            return edge2
        }

        numMap[edge1] = struct{}{}
        numMap[edge2] = struct{}{}
    }
    return 1
}
