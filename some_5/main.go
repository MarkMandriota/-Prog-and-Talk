package main

import (
	"fmt"
	"strings"
)

func main() {
	res := make(map[byte]struct{})

	str := strings.Join([]string{"asdaads", "aldsakdkaskd", "fgkekdkwkd", "wkskakska", "dsdawferg", "opffodokokk", "jgvhncncn"}, "")
	for _, ch := range []byte(str) {
		res[ch] = struct{}{}
	}

	out := make([]byte, 0, len(res))
	for k := range res {
		out = append(out, k)
	}

	fmt.Println(string(out))
}
