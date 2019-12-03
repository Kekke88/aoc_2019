package day3

import (
	"strconv"
	"strings"
)

const DIRECTION_UP = "U"
const DIRECTION_DOWN = "D"
const DIRECTION_LEFT = "L"
const DIRECTION_RIGHT = "R"

type Vector2D struct {
	X, Y int
}

type Instruction struct {
	Direction string
	Amount    int
}

type Wire struct {
	Path         []Vector2D
	Instructions []Instruction
}

func create(instructions string) *Wire {
	w := Wire{nil, nil}

	w.ParseInstructions(instructions)

	return &w
}

func (w *Wire) Trace() {
	x := 0
	y := 0
	for _, instruction := range w.Instructions {
		switch instruction.Direction {
		case DIRECTION_UP:
			for i := 0; i < instruction.Amount; i++ {
				x++
				w.Path = append(w.Path, Vector2D{x, y})
			}
		case DIRECTION_DOWN:
			for i := 0; i < instruction.Amount; i++ {
				x--
				w.Path = append(w.Path, Vector2D{x, y})
			}
		case DIRECTION_RIGHT:
			for i := 0; i < instruction.Amount; i++ {
				y++
				w.Path = append(w.Path, Vector2D{x, y})
			}
		case DIRECTION_LEFT:
			for i := 0; i < instruction.Amount; i++ {
				y--
				w.Path = append(w.Path, Vector2D{x, y})
			}
		}
	}
}

func (w Wire) TracePoint(p Vector2D) int {
	steps := 0
	x := 0
	y := 0

	for _, instruction := range w.Instructions {
		switch instruction.Direction {
		case DIRECTION_UP:
			for i := 0; i < instruction.Amount; i++ {
				x++
				steps++
				currentPoint := Vector2D{x, y}
				w.Path = append(w.Path, currentPoint)

				if p == currentPoint {
					return steps
				}
			}
		case DIRECTION_DOWN:
			for i := 0; i < instruction.Amount; i++ {
				x--
				steps++
				currentPoint := Vector2D{x, y}
				w.Path = append(w.Path, currentPoint)

				if p == currentPoint {
					return steps
				}
			}
		case DIRECTION_RIGHT:
			for i := 0; i < instruction.Amount; i++ {
				y++
				steps++
				currentPoint := Vector2D{x, y}
				w.Path = append(w.Path, currentPoint)

				if p == currentPoint {
					return steps
				}
			}
		case DIRECTION_LEFT:
			for i := 0; i < instruction.Amount; i++ {
				y--
				steps++
				currentPoint := Vector2D{x, y}
				w.Path = append(w.Path, currentPoint)

				if p == currentPoint {
					return steps
				}
			}
		}
	}
	return steps
}

func (w Wire) GetIntersections(w2 *Wire) []Vector2D {
	var intersections []Vector2D

	for _, v1 := range w.Path {
		for _, v2 := range w2.Path {
			if v1 == v2 {
				intersections = append(intersections, v1)
			}
		}
	}

	return intersections
}

func (w *Wire) ParseInstructions(instructions string) {
	parts := strings.Split(instructions, ",")
	for _, part := range parts {
		dir := string(part[0])
		num, err := strconv.Atoi(trimLeftChar(part))
		if err != nil {
			panic(err)
		}

		w.Instructions = append(w.Instructions, Instruction{dir, num})
	}
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
