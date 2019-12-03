package day3

import (
	"fmt"
)

func PartTwo() {
	w := create("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w.Trace()

	w2 := create("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	w2.Trace()

	intersections := w.GetIntersections(w2)
	shortestDistance := 0

	for _, intersection := range intersections {
		steps := w.TracePoint(intersection) + w2.TracePoint(intersection)

		if shortestDistance <= 0 || shortestDistance > steps {
			shortestDistance = steps
		}
	}

	fmt.Printf("Shortest distance to an intersection is %d steps.\n", shortestDistance)
}

func PartOne() {
	w := create("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w.Trace()

	w2 := create("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	w2.Trace()

	intersections := w.GetIntersections(w2)

	closestDistance := 0

	for _, intersection := range intersections {
		if closestDistance <= 0 || closestDistance > intersection.X+intersection.Y {
			closestDistance = intersection.X + intersection.Y
		}
	}

	fmt.Printf("Shortest manhattan distance is: %d\n", closestDistance)
}
