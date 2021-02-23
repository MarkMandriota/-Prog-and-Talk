package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(map[bool]string{
		true: "YES", false: "NO",
	}[((a+b)%4-3)%3 == 0])
}
