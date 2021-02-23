package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []byte("2214-544245-54544-5464556-645454")

	arr := make([]int, 0)
	sign := 1

	for _, ch := range str {
		if ch == '-' {
			sign *= -1
			continue
		}

		arr = append(arr, int(ch-'0')*sign)
		sign = 1
	}

	sort.Ints(arr)

	fmt.Println(arr)
}
