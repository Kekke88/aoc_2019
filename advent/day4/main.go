package day4

import (
	"fmt"
	"strconv"
)

func PartTwo() {
	possiblePasswords := 0

	for i := 156218; i < 652527; i++ {
		// Cast to string and compare
		pass := strconv.Itoa(i)

		if hasDoubleDigitNotInGroup(pass) && isNotDecreasing(pass) {
			possiblePasswords++
		}
	}

	fmt.Printf("Number of passwords matching criteria: %d\n", possiblePasswords)
}

func PartOne() {
	possiblePasswords := 0

	for i := 156218; i < 652527; i++ {
		// Cast to string and compare
		pass := strconv.Itoa(i)

		if hasDoubleDigit(pass) && isNotDecreasing(pass) {
			possiblePasswords++
		}
	}

	fmt.Printf("Number of passwords matching criteria: %d\n", possiblePasswords)
}

func hasDoubleDigitNotInGroup(pass string) bool {
	for index, _ := range pass {
		if len(pass) > index+1 {
			if pass[index] == pass[index+1] {
				// Has double digit, make sure its not in a group
				if len(pass) > index+2 {
					if pass[index] != pass[index+2] {
						if index <= 0 {
							// If we have a double digit and we have no digit to check before then return true since its a double digit not part of a group
							return true
						} else if pass[index] != pass[index-1] {
							// If index is higher than 0 we need to make sure the previous char is not part of the double digit
							return true
						}
					}
				} else if index > 0 {
					// This case only happens when its the second last (index 5 with 6 length string), when index 5 and 6 match, make sure that 4 does not match
					if pass[index-1] != pass[index] {
						return true
					}
				}
			}
		}
	}

	return false
}

func isNotDecreasing(pass string) bool {
	var prevChar rune

	for index, char := range pass {
		if index == 0 {
			prevChar = char
		}

		if char < prevChar {
			return false
		}

		prevChar = char
	}

	return true
}

func hasDoubleDigit(pass string) bool {
	var prevChar rune

	for index, char := range pass {
		if index > 0 && char == prevChar {
			return true
		}

		prevChar = char
	}

	return false
}
