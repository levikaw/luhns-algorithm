package main

import (
	"fmt"
)

func main() {
	fmt.Println(validate_credit_card_number("7762 8881 0311 1881")) // false
	fmt.Println(validate_credit_card_number("4111 1111 1111 1111")) // true
}
