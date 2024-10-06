package main

import "fmt"

func main() {
	words := []string{"go", "java", "go", "python", "go", "java"}
	wordCountMap := map[string]int{}

	for i := range words {
		_, ok := wordCountMap[words[i]]
		if ok {
			wordCountMap[words[i]] += 1
		} else {
			wordCountMap[words[i]] = 1
		}
	}
	for k, v := range wordCountMap {
		fmt.Printf("Number of occurence of word \"%s\" in slice words is: %d.\n", k, v)
	}
}
