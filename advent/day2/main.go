package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func runInstructions(pos1 int, pos2 int, parts []int) int {
	currentOpcode := 0
	parts[1] = pos1
	parts[2] = pos2
	for {
		if parts[currentOpcode] == 99 {
			break
		}

		// Read instruction positions
		pos1 := parts[currentOpcode+1]
		pos2 := parts[currentOpcode+2]
		sumPos := parts[currentOpcode+3]

		// Read values
		val1 := parts[pos1]
		val2 := parts[pos2]

		if parts[currentOpcode] == 1 {
			parts[sumPos] = val1 + val2
		} else if parts[currentOpcode] == 2 {
			parts[sumPos] = val1 * val2
		}

		// Move to next opcode
		currentOpcode += 4
	}

	return parts[0]
}

func PartTwo() {
	file, err := os.Open("advent/day2/input.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		v := scanner.Text()

		backupParts := strings.Split(v, ",")
		parts, err := sliceAtoi(backupParts)
		if err != nil {
			panic(err)
		}

		// Brute force
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				val := runInstructions(i, j, parts)

				if val == 19690720 {
					fmt.Printf("Noun: %d, Verb: %d, Answer: %d", i, j, 100*i+j)
				}

				parts, err = sliceAtoi(backupParts)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func PartOne() {
	file, err := os.Open("advent/day2/input.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		v := scanner.Text()

		input := strings.Split(v, ",")
		parts, err := sliceAtoi(input)
		if err != nil {
			panic(err)
		}

		val := runInstructions(12, 2, parts)

		fmt.Println(val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
