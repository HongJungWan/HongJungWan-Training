// PG - 완주하지 못한 선수
package main

import "fmt"

func findNonFinisher(participant, completion []string) string {
	counter := make(map[string]int)

	for _, name := range participant {
		counter[name]++
	}
	for _, name := range completion {
		counter[name]--
	}
	for name, count := range counter {
		if count > 0 {
			return name
		}
	}
	return ""
}

func main() {
	participant := []string{"leo", "kiki", "eden"}
	completion := []string{"eden", "kiki"}
	fmt.Println(findNonFinisher(participant, completion)) // "leo"

	participant = []string{"marina", "josipa", "nikola", "vinko", "filipa"}
	completion = []string{"josipa", "filipa", "marina", "nikola"}
	fmt.Println(findNonFinisher(participant, completion)) // "vinko"

	participant = []string{"mislav", "stanko", "mislav", "ana"}
	completion = []string{"stanko", "ana", "mislav"}
	fmt.Println(findNonFinisher(participant, completion)) // "mislav"
}

// 로직
// 1. 참가자 수 count
// 2. 완주자 수 빼기
// 3. 완주하지 못한 참가자 Search

// 조건 : 모든 참가자가 완주한 경우 (이 경우는 문제 조건에 따라 발생하지 않음)
