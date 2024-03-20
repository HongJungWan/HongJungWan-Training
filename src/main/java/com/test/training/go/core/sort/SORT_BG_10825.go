package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Student struct {
	name    string
	korean  int
	english int
	math    int
}

type Students []Student

// Len은 sort.Interface를 구현하며, Students 슬라이스의 길이를 반환
func (s Students) Len() int {
	return len(s)
}

// Swap은 sort.Interface를 구현하며, Students 슬라이스의 두 원소의 위치를 바꾼다
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less는 sort.Interface를 구현하며, 정렬 기준에 따라 두 학생을 비교
func (s Students) Less(i, j int) bool {
	if s[i].korean == s[j].korean {
		if s[i].english == s[j].english {
			if s[i].math == s[j].math {
				return s[i].name < s[j].name // 이름이 사전 순으로 앞선 경우
			}
			return s[i].math > s[j].math // 수학 점수가 더 높은 경우
		}
		return s[i].english < s[j].english // 영어 점수가 더 낮은 경우
	}
	return s[i].korean > s[j].korean // 국어 점수가 더 높은 경우
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n) // N 개수 입력 받기

	students := make(Students, 0, n)
	for i := 0; i < n; i++ {
		var name string
		var korean, english, math int
		fmt.Fscanln(reader, &name, &korean, &english, &math)
		students = append(students, Student{name, korean, english, math})
	}

	// 정렬 기준에 따라 학생 정보를 정렬
	sort.Sort(students)

	for _, student := range students {
		fmt.Println(student.name)
	}
}
