package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func issueCardNumber(brands string, issuers string, brand string, issuer string) {
	// fmt.Println(brand)
	brandFile, _ := os.Open("brands.txt")
	issuerFile, _ := os.Open("issuers.txt")

	brand = strings.ToLower(brand)
	issuer = strings.ToLower(issuer)

	issuerScanner := bufio.NewScanner(issuerFile)
	brandScanner := bufio.NewScanner(brandFile)

	var issuerBins []string
	var brandBins []string

	for issuerScanner.Scan() {
		line := issuerScanner.Text()
		fields := strings.Split(line, ":")
		if strings.Contains(strings.ToLower(fields[0]), issuer) {
			issuerBins = append(issuerBins, fields[1])
		}
	}

	for brandScanner.Scan() {
		line := brandScanner.Text()
		fields := strings.Split(line, ":")

		for _, bin := range issuerBins {
			if strings.Contains(strings.ToLower(fields[0]), brand) && strings.Contains(bin, fields[1]) {
				brandBins = append(brandBins, fields[1])
			}
		}
	}
	// fmt.Println("{", strings.ToUpper(issuer), "} bin code:", issuerBins)

	if len(brandBins) == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: No matching brand bins found")
		os.Exit(1)
	}

	// Filter issuerBins to match brandBins
	var resBins []string
	for _, issuerBin := range issuerBins {
		for _, brandBin := range brandBins {
			if strings.HasPrefix(issuerBin, brandBin) {
				resBins = append(resBins, issuerBin)
				break
			}
		}
	}

	if len(resBins) == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: No matching issuer bins found")
		os.Exit(1)
	}

	// Select a random bin from resBins
	randomBinIndex := rand.Intn(len(resBins))
	firstSixBin := resBins[randomBinIndex]

	switch {
	case strings.Contains(strings.ToLower("MASTERCARD"), brand):
		validGenerateRandomNumber(firstSixBin, 10)

	case strings.Contains(strings.ToLower("INSTAPAYMENT"), brand):
		validGenerateRandomNumber(firstSixBin, 10)

	case strings.Contains(strings.ToLower("JCB"), brand):
		validGenerateRandomNumber(firstSixBin, 10)

	case strings.Contains(strings.ToLower("AMEX"), brand):
		validGenerateRandomNumber(firstSixBin, 9)

	case strings.Contains(strings.ToLower("DinersClubCarteBlanche"), brand):
		// dinnerCardLength := []int{8, 10}
		// randomIndex := rand.Intn(2)
		validGenerateRandomNumber(firstSixBin, 8)

	case strings.Contains(strings.ToLower("VISA"), brand):
		visaCardLength := []int{7, 10}
		randomIndex := rand.Intn(2)
		validGenerateRandomNumber(firstSixBin, visaCardLength[randomIndex])
	case strings.Contains(strings.ToLower("DinersClubInternational"), brand):

		validGenerateRandomNumber(firstSixBin, 8)

	case strings.Contains(strings.ToLower("DISCOVER"), brand):
		validGenerateRandomNumber(firstSixBin, 10)
	default:
		fmt.Println("EROR: no such brand was found")
		os.Exit(1)

	}
}

func generateRandomNumber(firstSixBin string, cardLength int) string {
	number := ""
	for i := 0; i < cardLength; i++ {
		digit := rand.Intn(10)
		number += strconv.Itoa(digit)
	}
	randomCardNumber := firstSixBin + number

	return randomCardNumber
}

func validGenerateRandomNumber(firstSixBin string, cardLength int) {
	for {
		cardNumber := generateRandomNumber(firstSixBin, cardLength)
		if validateCardNumber(cardNumber) == "OK" {
			fmt.Println(cardNumber)
			break
		}
	}
}
