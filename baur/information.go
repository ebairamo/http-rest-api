package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func informationCardNumber(cardNumber string) {
	firstSixDigit := cardNumber[:6]
	cardInformations := readAndFind(firstSixDigit)
	fmt.Fprintln(os.Stdout, cardNumber)
	fmt.Fprintln(os.Stdout, "Correct: yes")
	fmt.Fprintln(os.Stdout, "Card Brand:", cardInformations[1])
	fmt.Fprintln(os.Stdout, "Card Issuer:", cardInformations[0])
}

func readAndFind(firstSixDigit string) []string {
	var cardInformations []string

	brandFile, _ := os.Open("brands.txt")
	issuerFile, _ := os.Open("issuers.txt")

	issuerScanner := bufio.NewScanner(issuerFile)

	for issuerScanner.Scan() {
		line := issuerScanner.Text()
		fields := strings.Split(line, ":")

		if len(fields) == 2 {
			bin := strings.TrimSpace(fields[1])
			if bin == firstSixDigit {
				issuer := strings.TrimSpace(fields[0])
				cardInformations = append(cardInformations, issuer)
			}

		}
	}

	brandScanner := bufio.NewScanner(brandFile)

	for brandScanner.Scan() {
		line := brandScanner.Text()
		fields := strings.Split(line, ":")

		if len(fields) == 2 {
			bin := strings.TrimSpace(fields[1])
			if strings.Contains(firstSixDigit, bin) {
				fmt.Println(bin)
				brand := strings.TrimSpace(fields[0])
				cardInformations = append(cardInformations, brand)
			}
		}
	}
	fmt.Println(cardInformations)

	return cardInformations
}
