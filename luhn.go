package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Luhn's Algorithm
func validate_credit_card_number(number string) bool {

	checkSum := 0

	// for loop by every string element without spaces
	for idx, digit := range strings.ReplaceAll(number, " ", "") {

		// parse rune to int
		intVal, err := strconv.Atoi(string(digit))

		// if rune is not int string return false
		if err != nil {
			fmt.Println("Error:", err)
			return false
		}

		// every even digit
		if idx%2 == 0 {
			multiplication := 2 * intVal

			if multiplication > 9 {
				checkSum += (multiplication - 9)
			} else {
				checkSum += (multiplication)
			}
		} else {
			checkSum += intVal
		}
	}

	return checkSum%10 == 0
}
