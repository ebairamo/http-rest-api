package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "ERROR: Command not provided")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "validate":
		validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
		stdin := validateCmd.Bool("stdin", false, "Read input from standard input (stdin)")
		validateCmd.Parse(os.Args[2:])
		handleValidate(*stdin)

	case "generate":
		generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		pick := generateCmd.Bool("pick", false, "Флаг для выбора определенной логики генерации")
		generateCmd.Parse(os.Args[2:])
		handleGenerate(*pick)
		// Вызов функции генерации карты

	case "information":
		issuerCmd := flag.NewFlagSet("issue", flag.ExitOnError)
		brands := issuerCmd.String("brands", "brands.txt", "Allows selecting a brand from a file")
		issuers := issuerCmd.String("issuers", "issuers.txt", "Allows selecting a issuer from a file")
		brand := issuerCmd.String("brand", "", "Allows selecting a specific brand")
		issuer := issuerCmd.String("issuer", "", "Allows selecting a specific issuer")
		issuerCmd.Parse(os.Args[2:])
		handleInformation(*brands, *issuers, *brand, *issuer)

	case "issue":
		issuerCmd := flag.NewFlagSet("issue", flag.ExitOnError)
		brands := issuerCmd.String("brands", "brands.txt", "Allows selecting a brand from a file")
		issuers := issuerCmd.String("issuers", "issuers.txt", "Allows selecting a issuer from a file")
		brand := issuerCmd.String("brand", "", "Allows selecting a specific brand")
		issuer := issuerCmd.String("issuer", "", "Allows selecting a specific issuer")
		issuerCmd.Parse(os.Args[2:])
		handleIssue(*brands, *issuers, *brand, *issuer)

	default:
		fmt.Fprintln(os.Stderr, "ERROR: Unknown command")
		os.Exit(1)
	}
}

func handleValidate(stdin bool) {
	var cardNumbers []string

	if stdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			cardNumbers = append(cardNumbers, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: Reading from stdin failed:", err)
			os.Exit(1)
		}
	} else {
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "ERROR: No card numbers provided")
			os.Exit(1)
		}
		cardNumbers = os.Args[2:]
	}

	for _, cardNumber := range cardNumbers {
		fmt.Println("Validating card number:", cardNumber)
		result := validateCardNumber(cardNumber)
		if result {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)

		}
	}
}

func handleGenerate(pick bool) {

	if pick {
		cardNumber := os.Args[3]
		fmt.Println("Опция pick активирована.")
		// Здесь можно добавить специфичную логику, если pick установлен
		result := generateCard(cardNumber)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(len(result) - 1)
		fmt.Println(result[num])

	} else {
		cardNumber := os.Args[2]
		result := generateCard(cardNumber)
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}
}
