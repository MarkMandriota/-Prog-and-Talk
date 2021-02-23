package main

import "fmt"

func main() {
	res := make([]bool, 0)

	for _, el := range []int{54, 65, 23, 76, 89, 21, 55, 65, 20, 11, 40} {
		res = append(res, check(el))
	}

	fmt.Println(res)
}

func check(n int) bool {
	return n%10 <= n/10
}
