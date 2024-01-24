// BG - 철벽 보안 알고리즘

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWordCount(scanner *bufio.Scanner) int {
	scanner.Scan()
	var N int
	fmt.Sscanf(scanner.Text(), "%d", &N)
	return N
}

func createKeyMapping(scanner *bufio.Scanner) map[string]int {
	scanner.Scan()
	key := strings.Fields(scanner.Text())

	mapKey := make(map[string]int)
	for i, word := range key {
		mapKey[word] = i
	}
	return mapKey
}

func mapKeyToIndex(scanner *bufio.Scanner, mapKey map[string]int, N int) []int {
	scanner.Scan()
	key := strings.Fields(scanner.Text())

	normal := make([]int, N)
	for i, word := range key {
		normal[i] = mapKey[word]
	}
	return normal
}

func readCipherText(scanner *bufio.Scanner, N int) []string {
	scanner.Scan()
	return strings.Fields(scanner.Text())
}

func decryptCipherText(cipherText []string, normal []int, N int) []string {
	result := make([]string, N)
	for i := 0; i < N; i++ {
		result[normal[i]] = cipherText[i]
	}
	return result
}

func printDecryptedText(decryptedText []string) {
	for _, word := range decryptedText {
		fmt.Print(word + " ")
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var T int
	fmt.Sscanf(scanner.Text(), "%d", &T)

	for tc := 1; tc <= T; tc++ {
		N := readWordCount(scanner)
		mapKey := createKeyMapping(scanner)
		normal := mapKeyToIndex(scanner, mapKey, N)
		cipherText := readCipherText(scanner, N)
		decryptedText := decryptCipherText(cipherText, normal, N)

		printDecryptedText(decryptedText)
	}
}
