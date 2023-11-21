package main

import (
	"fmt"
)

func Pg_11_22(arr []int, n int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		if length%2 != 0 && i%2 == 0 {
			arr[i] += n
		}
		if length%2 == 0 && i%2 != 0 {
			arr[i] += n
		}
	}
	return arr
}

func main() {
	fmt.Println(Pg_11_22([]int{49, 12, 100, 276, 33}, 27)) // [76 12 127 276 60]
	fmt.Println(Pg_11_22([]int{444, 555, 666, 777}, 100))  // [444 655 666 877]
}
