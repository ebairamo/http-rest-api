package main

import "fmt"

func main() {
	fmt.Println(FindLongestWord("The quick brown fox jumps over the lazy dog")) // "quick"
	fmt.Println(FindLongestWord("Alem School of Programming"))                  // "Programming"
	fmt.Println(FindLongestWord("Go is fun"))                                   // "fun"
}

func FindLongestWord(s string) string {
	var df string
	var dfs []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			df += string(s[i])
		}
		if s[i] == ' ' {
			dfs = append(dfs, df)
			df = ""
		}
	}
	dfs = append(dfs, df)
	dfe := 0
	dfm := ""
	for i := 0; i < len(dfs); i++ {
		if len(dfs[i]) > dfe {
			dfe = len(dfs[i])
			dfm = dfs[i]
		}
	}
	return dfm
}
