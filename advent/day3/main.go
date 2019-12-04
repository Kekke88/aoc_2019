package day3

// Challenge URL: https://adventofcode.com/2019/day/3

import (
	"fmt"
)

func PartTwo() {
	firstWire := create("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	firstWire.Trace()

	secondWire := create("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	secondWire.Trace()

	intersections := firstWire.GetIntersections(secondWire)
	shortestDistance := 0

	for _, intersection := range intersections {
		steps := firstWire.TracePoint(intersection) + secondWire.TracePoint(intersection)

		if shortestDistance <= 0 || shortestDistance > steps {
			shortestDistance = steps
		}
	}

	fmt.Printf("Shortest distance to an intersection is %d steps.\n", shortestDistance)
}

func PartOne() {
	firstWire := create("R8,U5,L5,D3")
	firstWire.Trace()

	secondWire := create("U7,R6,D4,L4")
	secondWire.Trace()

	intersections := firstWire.GetIntersections(secondWire)

	closestDistance := 0

	for _, intersection := range intersections {
		if closestDistance <= 0 || closestDistance > intersection.X+intersection.Y {
			closestDistance = intersection.X + intersection.Y
		}
	}

	fmt.Printf("Shortest manhattan distance is: %d\n", closestDistance)
}
