package main

import (
	"fmt"
)

func main() {
	fmt.Println(UniquePermutations("aabb")) // ["aabb", "abab", "abba", "baab", "baba", "bbaa"]
	fmt.Println(UniquePermutations("abc"))  // ["abc", "acb", "bac", "bca", "cab", "cba"]
	fmt.Println(UniquePermutations("aaa"))  // ["aaa"]
}

func UniquePermutations(s string) []string {
	permutations := make(map[string]struct{}) // Используем множество для уникальных перестановок
	generatePermutations([]rune(s), 0, permutations)

	// Преобразуем множество в срез
	result := make([]string, 0, len(permutations))
	for p := range permutations {
		result = append(result, p)
	}
	return result
}

// Рекурсивная функция для генерации перестановок
func generatePermutations(runes []rune, index int, permutations map[string]struct{}) {
	if index == len(runes)-1 {
		permutations[string(runes)] = struct{}{} // Добавляем перестановку в множество
		return
	}

	seen := make(map[rune]bool) // Отслеживаем символы, которые уже были использованы на этой позиции
	for i := index; i < len(runes); i++ {
		if seen[runes[i]] {
			continue // Пропускаем, если символ уже использовался
		}
		seen[runes[i]] = true

		// Меняем символы местами
		runes[index], runes[i] = runes[i], runes[index]
		generatePermutations(runes, index+1, permutations)
		// Возвращаем символы обратно
		runes[index], runes[i] = runes[i], runes[index]
	}
}
