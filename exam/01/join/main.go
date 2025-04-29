package main

import "fmt"

func main() {
	fmt.Println(Join([]string{"salem", "student"}, " "))
	fmt.Println(Join([]string{"1", "2", "3"}, ", "))
}

func Join(elements []string, sep string) string {
	join := ""
	for i := 0; i <= len(elements)-1; i++ {
		if i == len(elements)-1 {
			join += elements[i]
			break
		}
		join += elements[i] + sep

	}
	return join
}
