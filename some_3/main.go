package main

import "fmt"

func main() {
	arr := []interface{}{124, 122, 42, 21, 45, 1}
	fmt.Println(append(arr, "hello", "nick", "good", "win"))
}
