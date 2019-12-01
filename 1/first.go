package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1")
	partOne()

	fmt.Println("Part 2")
	partTwo()
}

func partTwo() {
	file, err := os.Open("first.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		v := scanner.Text()

		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		// Fuel required to launch a given module is based on its mass.
		// Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
		i = i/3 - 2

		sum += i

		// Part 2: for each module mass, calculate its fuel and add it to the total.
		// Then, treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative.
		for i > 0 {
			i = i/3 - 2

			if i > 0 {
				sum += i
			}
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partOne() {
	file, err := os.Open("first.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		v := scanner.Text()

		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		// Fuel required to launch a given module is based on its mass.
		// Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
		sum += i/3 - 2
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
