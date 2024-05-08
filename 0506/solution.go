package leetcode

type p struct {
    i int
    score int
}

func findRelativeRanks(score []int) []string {
    m := make([]p, 0)

    for i, v := range score {
        para := p{
            i: i,
            score: v,
        }

        m = append(m, para)
    }

    result := make([]string, len(m))

    slices.SortFunc(m, func(a, b p) int{
        if a.score > b.score {
            return -1
        }

        return 1
    })

    for i := 0; i < len(m); i++ {
        if i == 0 {
            result[m[i].i] = "Gold Medal"
        } else if i == 1 {
            result[m[i].i] = "Silver Medal"
        } else if i == 2 {
            result[m[i].i] = "Bronze Medal"
        } else {
            result[m[i].i] = strconv.Itoa(i+1)
        }
    }

    return result
}
