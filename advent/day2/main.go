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
		currentOpcode := 0

		// Brute force
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				parts[1] = i
				parts[2] = j

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

					// Move 4 positions forward
					currentOpcode += 4
				}
				// End opcode instructions
				// Reset memory
				if parts[0] == 19690720 {
					fmt.Printf("Noun: %d, Verb: %d, Answer: %d", i, j, 100*i+j)
				}

				// Reset memory
				currentOpcode = 0
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
		currentOpcode := 0

		// Read [0], do calcs then move 4 positions forward to [4]
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

			// Move 4 positions forward
			currentOpcode += 4
		}

		fmt.Println(parts[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
