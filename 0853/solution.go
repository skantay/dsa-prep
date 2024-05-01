package leetcode

type car struct {
    position int
    speed int
}


func carFleet(target int, position []int, speed []int) int {
    cars := make([]car, 0, len(position))

    for i := range position {
        car := car{
            position: position[i],
            speed: speed[i],
        }

        cars = append(cars, car)
    }

    slices.SortFunc(cars, func(a, b car) int {
        if n := cmp.Compare(a.position, b.position); n != 0 {
			return n
		}

		return cmp.Compare(a.speed, b.speed)
    })

    stack := make([]float64, 0, len(position))

    for i := range cars {
        time := float64(target - cars[i].position) / float64(cars[i].speed)

        for len(stack) != 0 && time >= stack[len(stack)-1]  {
            stack = stack[:len(stack) - 1]
        }

        stack = append(stack, time)
    } 

    return len(stack)
}
