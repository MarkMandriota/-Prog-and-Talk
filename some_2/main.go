package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{55, -5, 10, 6, 4, 1, 123, 34, 56567, 88975, 23, -3455, 3457, -467567, -675, -3234}
	sort.Ints(arr)

	fmt.Println(arr)
}
