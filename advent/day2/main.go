package day2

// Challenge URL: https://adventofcode.com/2019/day/2

import (
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadFile("advent/day2/input.dat")
	if err != nil {
		panic(err)
	}

	backupParts := strings.Split(string(data), ",")
	parts, err := sliceAtoi(backupParts)
	if err != nil {
		panic(err)
	}

	// Brute force
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			elfComputer := create(parts)
			elfComputer.SetMemState(1, i)
			elfComputer.SetMemState(2, j)
			elfComputer.Run()

			if elfComputer.Value() == 19690720 {
				fmt.Printf("Noun: %d, Verb: %d, Answer: %d\n", i, j, 100*i+j)
			}

			parts, err = sliceAtoi(backupParts)
			if err != nil {
				panic(err)
			}
		}
	}
}

func PartOne() {
	data, err := ioutil.ReadFile("advent/day2/input.dat")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), ",")
	parts, err := sliceAtoi(input)
	if err != nil {
		panic(err)
	}

	elfComputer := create(parts)
	elfComputer.SetMemState(1, 12)
	elfComputer.SetMemState(2, 2)
	elfComputer.Run()

	fmt.Println(elfComputer.Value())
}
