package main

import "fmt"

func main() {
	var word string
	fmt.Println("Enter a word to check!")
	fmt.Scanf("%s", &word)
	index := firstNonRepeatingIndex(word)
	fmt.Println("Index of first non-repeating character:", index)
}

func firstNonRepeatingIndex() int {
	type charData struct {
		count int
		index int
	}

	results := make(map[rune]charData)
	var order []rune

	for i, r := range word {
		if data, ok := results[r]; ok {
			data.count++
			results[r] = data
		} else {
			results[r] = charData{count: 1, index: i}
			order = append(order, r)
		}
	}

	for _, char := range order {
		if results[char].count == 1 {
			return results[char].index
		}
	}

	return -1
}

