package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func handleValidate(stdin bool) {
	var consoleInputValues []string
	isIncorrect := false
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "ERROR: Value not input")
		os.Exit(1)
	}
	if stdin {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: reading stdin:", err)
			os.Exit(1)
		}
		consoleInputValues = strings.Fields(string(input))
	} else {
		consoleInputValues = os.Args[2:]
	}
	for _, card := range consoleInputValues {
		if validateCardNumber(card) == "INCORRECT" {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			isIncorrect = true
		} else if validateCardNumber(card) == "OK" {
			fmt.Fprintln(os.Stdout, "OK")
		}
	}
	if isIncorrect {
		os.Exit(1)
	}
}

func handleGenerate(pick bool) {
	var consoleInputValues []string

	switch {
	case len(os.Args) == 4 && pick == true:
		consoleInputValues = os.Args[3:]
	case len(os.Args) == 3 && pick == false:
		consoleInputValues = os.Args[2:]
	default:
		fmt.Fprintln(os.Stderr, "ERROR: Invalid input")
		os.Exit(1)
	}

	for _, card := range consoleInputValues {
		generateCardNumber(card, pick)
	}
}

func handleInformation(brands string, issuers string, stdin bool) {
	// var consoleInputValues []string
	// fmt.Println(stdin)
	// if stdin {
	// 	input, err := io.ReadAll(os.Stdin)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, "ERROR: Reading stdin:", err)
	// 		os.Exit(1)
	// 	}
	// 	consoleInputValues = strings.Fields(string(input))
	// }

	// if len(os.Args) < 5 {
	// 	fmt.Fprintln(os.Stderr, "ERROR: Invalid input")
	// 	os.Exit(1)
	// } else {
	// 	if brands == "brands.txt" && issuers == "issuers.txt" {
	// 		consoleInputValues = os.Args[4:]
	// 	} else {
	// 		fmt.Fprintln(os.Stderr, "ERROR: File not found")
	// 		os.Exit(1)
	// 	}
	// }
	// for _, card := range consoleInputValues {
	// 	if validateCardNumber(card) == "INCORRECT" {

	// 		fmt.Fprintln(os.Stderr, card)
	// 		fmt.Fprintln(os.Stderr, "Correct: no")
	// 		fmt.Fprintln(os.Stderr, "Card Brand: -")
	// 		fmt.Fprintln(os.Stderr, "Card Issuer: -")
	// 		os.Exit(1)
	// 	} else {
	// 		informationCardNumber(card)
	// 	}
	// }
	var consoleInputValues []string

	if stdin {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: Reading stdin:", err)
			os.Exit(1)
		}
		consoleInputValues = strings.Fields(string(input))
	} else {
		if len(os.Args) < 5 {
			fmt.Fprintln(os.Stderr, "ERROR: Invalid input")
			os.Exit(1)
		}
		consoleInputValues = os.Args[4:]
	}

	if brands == "brands.txt" && issuers == "issuers.txt" {
		for _, card := range consoleInputValues {
			if validateCardNumber(card) == "INCORRECT" {
				fmt.Fprintln(os.Stderr, card)
				fmt.Fprintln(os.Stderr, "Correct: no")
				fmt.Fprintln(os.Stderr, "Card Brand: -")
				fmt.Fprintln(os.Stderr, "Card Issuer: -")
				os.Exit(1)
			} else {
				informationCardNumber(card)
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "ERROR: File not found")
		os.Exit(1)
	}
}

func handleIssue(brands string, issuers string, brand string, issuer string) {
	if len(os.Args) < 6 {

		fmt.Fprintln(os.Stderr, "ERROR: Invalid input")
		os.Exit(1)
	} else {
		if brands == "brands.txt" && issuers == "issuers.txt" && brand != "" && issuer != "" {
			issueCardNumber(brands, issuers, brand, issuer)
		} else {
			fmt.Fprintln(os.Stderr, "ERROR: Invalid input")
			os.Exit(1)
		}
	}
}
