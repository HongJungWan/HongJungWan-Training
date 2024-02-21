// BG - Maximum Word Frequency

package main

import (
	"bufio"
	"fmt"
	"os"
)

func readWords(n int) map[string]int {
	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		word := scanner.Text()
		wordFreq[word]++
	}
	return wordFreq
}

func findMostFrequentWord(wordFreq map[string]int) (string, int) {
	var maxFreq int
	var mostFreqWord string
	for word, freq := range wordFreq {
		if freq > maxFreq || (freq == maxFreq && word > mostFreqWord) {
			mostFreqWord = word
			maxFreq = freq
		}
	}
	return mostFreqWord, maxFreq
}

func main() {
	var n int
	fmt.Scan(&n)

	wordFreq := readWords(n)
	mostFreqWord, maxFreq := findMostFrequentWord(wordFreq)

	fmt.Printf("%s %d\n", mostFreqWord, maxFreq)
}
