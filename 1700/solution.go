package leetcode

func countStudents(students []int, sandwiches []int) int {
    var stdInd, sndwInd int

    haveEaten := make(map[int]struct{})

    for i := 0; i < len(students) && len(students) != len(haveEaten); {
        _, ok := haveEaten[stdInd]
        if students[stdInd] == sandwiches[sndwInd] && !ok {
            i = 0

            haveEaten[stdInd] = struct{}{}

            sndwInd++
        } else {
            i++
        }

        stdInd++

        if stdInd == len(students) {
            stdInd = 0
        }
    }

    return len(students) - len(haveEaten)
}
