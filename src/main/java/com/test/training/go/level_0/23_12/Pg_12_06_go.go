package main

import "fmt"

func Pg_12_06(todo_list []string, finished []bool) []string {
	result := []string{}

	for i, check := range finished {
		if check == false {
			result = append(result, todo_list[i])
		}
	}
	return result
}

func main() {
	todo_list := []string{"problemsolving", "practiceguitar", "swim", "studygraph"}
	finished := []bool{true, false, true, false}

	fmt.Println(Pg_12_06(todo_list, finished))
}
