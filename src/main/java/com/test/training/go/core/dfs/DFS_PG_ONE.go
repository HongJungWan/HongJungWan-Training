// PG - 타켓 넘버
package main

import "fmt"

func solution(numbers []int, target int) int {
	return dfs(numbers, target, 0, 0)
}

func dfs(numbers []int, target, index, currentSum int) int {
	if index == len(numbers) {
		if currentSum == target {
			return 1
		}
		return 0
	}
	return dfs(numbers, target, index+1, currentSum+numbers[index]) +
		dfs(numbers, target, index+1, currentSum-numbers[index])
}

func main() {
	numbers := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println("target number:", solution(numbers, target))

	numbers2 := []int{4, 1, 2, 1}
	target2 := 4
	fmt.Println("target number:", solution(numbers2, target2))
}
