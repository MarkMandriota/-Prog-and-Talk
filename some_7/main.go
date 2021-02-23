package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 44, 5, 65, 22, 34, 987, 87, 86885, 885485, 94648, 84584, 88454}

	var sum int
	for _, n := range arr {
		if n%2 == 0 {
			sum += n * n
		}
	}

	fmt.Println(sum)
}
