// Неправильно
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count(strings.Repeat("A", 100), strings.Repeat("A", 50)))
}
