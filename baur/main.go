package main

import (
	"flag"
	"fmt"
	"os"
)

// var (
// 	pick    bool
// 	brands  string
// 	issuers string
// 	brand   string
// 	issuer  string
// 	// stdin          bool
// 	cmdValidate    bool
// 	cmdGenerate    bool
// 	cmdInformation bool
// 	cmdIsssue      bool
// )

// func init() {
// 	flag.BoolVar(&pick, "pick", false, "Allows random selection of one entry")
// 	flag.StringVar(&brands, "brands", "", "Allows selecting a brand from a file")
// 	flag.StringVar(&issuers, "issuers", "", "Allows you to select an issuer from a file")
// 	flag.BoolVar(&stdin, "stdin", false, "Transmitting a number from the standard input (stdin)")
// 	flag.StringVar(&brand, "brand", "", "Allows selecting a specific brand")
// 	flag.StringVar(&issuer, "issuer", "", "Allows selecting a specific issuer")
// 	flag.BoolVar(&cmdValidate, "validate", false, "Generate card numbercommand")
// 	flag.BoolVar(&cmdGenerate, "generate", false, "Generate card numbercommand")
// 	flag.BoolVar(&cmdInformation, "information", false, "Generate card numbercommand")
// 	flag.BoolVar(&cmdIsssue, "issue", false, "Generate card numbercommand")
// }

func main() {
	if len(os.Args) > 1 {

		cmd := os.Args[1]

		switch cmd {
		case "validate":
			validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
			stdin := validateCmd.Bool("stdin", false, "Transmitting a number from the standard input (stdin)")
			validateCmd.Parse(os.Args[2:])
			handleValidate(*stdin)
		case "generate":
			generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
			pick := generateCmd.Bool("pick", false, "Allows random selection of one entry")
			generateCmd.Parse(os.Args[2:])
			handleGenerate(*pick)
		case "information":
			informationCmd := flag.NewFlagSet("information", flag.ExitOnError)
			brands := informationCmd.String("brands", "brands.txt", "Allows selecting a brand from a file")
			issuers := informationCmd.String("issuers", "issuers.txt", "Allows selecting a issuer from a file")
			stdin := informationCmd.Bool("stdin", false, "Transmitting a number from the standard input (stdin)")
			informationCmd.Parse(os.Args[2:])
			handleInformation(*brands, *issuers, *stdin)
		case "issue":
			issuerCmd := flag.NewFlagSet("issue", flag.ExitOnError)
			brands := issuerCmd.String("brands", "brands.txt", "Allows selecting a brand from a file")
			issuers := issuerCmd.String("issuers", "issuers.txt", "Allows selecting a issuer from a file")
			brand := issuerCmd.String("brand", "", "Allows selecting a specific brand")
			issuer := issuerCmd.String("issuer", "", "Allows selecting a specific issuer")
			issuerCmd.Parse(os.Args[2:])
			handleIssue(*brands, *issuers, *brand, *issuer)
		default:
			fmt.Fprintln(os.Stderr, "ERROR: command not found")
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "ERROR: command not use")
		os.Exit(1)
	}
}
