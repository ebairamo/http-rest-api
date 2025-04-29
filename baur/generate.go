package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func validateAsterisksPlacement(cardNumber string) bool {
	startIndex := strings.Index(cardNumber, "*")
	if startIndex == -1 {
		fmt.Fprintln(os.Stderr, "ERROR: asterisks not found")
		os.Exit(1)
	}

	return strings.Trim(cardNumber[startIndex:], "*") == ""
}

func generateCardNumber(cardNumber string, pick bool) {
	if len(cardNumber) < 13 || len(cardNumber) > 16 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid card length")
		os.Exit(1)
	}
	asterisksCount := 0
	digitCount := 0
	for _, digit := range cardNumber {
		if digit >= '0' && digit <= '9' {
			digitCount++
		} else if digit == '*' {
			asterisksCount++
		} else {
			fmt.Fprintln(os.Stderr, "ERROR: Invalid input")
			os.Exit(1)
		}
	}
	firstSixDigit := cardNumber[:6]
	cardInformations := readAndFind(firstSixDigit)
	fmt.Println(cardInformations)
	if cardInformations == nil || len(cardInformations) < 1 {
		fmt.Fprintln(os.Stderr, "ERROR: brand not found")
		os.Exit(1)
	}

	brand := cardInformations[0]

	if !strings.Contains("VISA, MASTERCARD, AMEX, DISCOVER, JCB, DinersClubCarteBlanche, DinersClubInternational, InstaPayment", brand) {
		fmt.Fprintln(os.Stderr, "ERROR: brand not found")
		os.Exit(1)
	}

	if !validateAsterisksPlacement(cardNumber) {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid asterisks placement")
		os.Exit(1)
	}

	cardLength := digitCount + asterisksCount

	if cardLength < 13 || cardLength > 16 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid card length")
		os.Exit(1)
	} else if asterisksCount < 1 || asterisksCount > 4 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid asterisks length")
		os.Exit(1)
	}

	switch {
	case brand == "VISA" && cardLength == 13 || cardLength == 16:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "MASTERCARD" && cardLength == 16:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "AMEX" && cardLength == 15:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "DISCOVER" && cardLength == 16:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "JCB" && cardLength == 16:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "DinersClubCarteBlanche" && cardLength == 14:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "DinersClubInternational" && cardLength == 14:
		generate(asterisksCount, cardLength, cardNumber, pick)
	case brand == "InstaPayment" && cardLength == 16:
		generate(asterisksCount, cardLength, cardNumber, pick)
	default:
		fmt.Fprintln(os.Stderr, "ERROR: Invalid card length")
		os.Exit(1)
	}
}

func generateRandom(cardNumberWithAsterisks string) string {
	cardNumber := ""

	for _, char := range cardNumberWithAsterisks {
		randomDigit := rand.Intn(10)
		if char == '*' {
			cardNumber += string(strconv.Itoa(randomDigit))
		} else {
			cardNumber += string(char)
		}
	}
	return cardNumber
}

func generate(asterisksCount int, cardLength int, cardNumberWithAsterisks string, pick bool) {
	// Generate random validate card number if used --pick flag
	if pick {
		cardNumber := ""
		for {
			cardNumber = generateRandom(cardNumberWithAsterisks)
			if validateCardNumber(cardNumber) == "OK" {
				fmt.Fprintln(os.Stdout, cardNumber)
				break
			}
		}
	}
	// count := 0
	if !pick {
		switch {
		case asterisksCount == 1:
			for a := 0; a < 10; a++ {
				cardNumber := cardNumberWithAsterisks
				cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(a), 1)
				if validateCardNumber(cardNumber) == "OK" {
					fmt.Println(cardNumber)
					// count++
				}
			}
		case asterisksCount == 2:
			for a := 0; a < 10; a++ {
				for b := 0; b < 10; b++ {
					cardNumber := cardNumberWithAsterisks
					cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(a), 1)
					cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(b), 1)
					if validateCardNumber(cardNumber) == "OK" {
						fmt.Println(cardNumber)
						// count++
					}
				}
			}
		case asterisksCount == 3:
			for a := 0; a < 10; a++ {
				for b := 0; b < 10; b++ {
					for c := 0; c < 10; c++ {
						cardNumber := cardNumberWithAsterisks
						cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(a), 1)
						cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(b), 1)
						cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(c), 1)

						if validateCardNumber(cardNumber) == "OK" {
							fmt.Println(cardNumber)
							// count++
						}

					}
				}
			}
		case asterisksCount == 4:
			for a := 0; a < 10; a++ {
				for b := 0; b < 10; b++ {
					for c := 0; c < 10; c++ {
						for d := 0; d < 10; d++ {
							cardNumber := cardNumberWithAsterisks
							cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(a), 1)
							cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(b), 1)
							cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(c), 1)
							cardNumber = strings.Replace(cardNumber, "*", strconv.Itoa(d), 1)

							if validateCardNumber(cardNumber) == "OK" {
								fmt.Println(cardNumber)
								// count++
							}
						}
					}
				}
			}
		}
	}
	// print("\n Count:", count, "\n")
}
