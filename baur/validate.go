package main

import (
	"strconv"
	"strings"
)

func validateCardNumber(cardNumber string) string {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	isNum := false
	sum := 0
	alternate := false

	for _, c := range cardNumber {
		if c < 0 || c > 9 {
			isNum = false
		}
		isNum = true
	}

	if len(cardNumber) < 13 || isNum == false {
		return "INCORRECT"
	}
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}

	if sum%10 == 0 {
		return "OK"
	}

	return "INCORRECT"
}
