package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Print("Enter fruits: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
		return
	}
	text := strings.ToLower(scanner.Text())

	words := strings.Fields(text)
	if len(words) == 0 {
		fmt.Println("No fruits entered.")
		return
	}

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	keys := make([]string, 0)

	for key := range wordCount {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, wordCount[key])
	}

	fmt.Println("Total fruits: ", len(words))
	fmt.Println("Unique fruits: ", len(keys))
}
