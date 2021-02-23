// go version go1.15.6 windows/amd64
package main

import "fmt"

func main() {
	// карта для перевода из морза
	morza := map[string]string{
		"++":   "И",
		"+-":   "А",
		"-+":   "Н",
		"--+":  "Г",
		"---+": "Ч",
	}

	var out string

	// Накапливаем буффер, пока он не будет совпадать с какой либо из букв,
	// затем добавляем букву в вывод и обнуляем буффер,
	// повторяя все снова пока не достигнем конца строки
	var buff string
	for _, el := range []byte("-++-++--++----+") {
		buff += string(el)
		if v, ok := morza[buff]; ok {
			buff = ""
			out += v
		}
	}

	fmt.Println(out)
}
