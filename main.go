package main

import "fmt"

var data = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "_", "+", "=", "~"}

func main() {
	encodeTitle := encode("hello=")
	fmt.Println(encodeTitle)
	decodeTitle := decode(encodeTitle)
	fmt.Println(decodeTitle)
}

func encode(title string) string {
	var encodeTitle = ""
	for _, t := range title {
		for i := range data {
			if string(t) == data[i] {
				if i+3 > len(data) {
					encodeTitle += data[i+3-len(data)]
				} else {
					encodeTitle += data[i+3]
				}
			}
		}
	}

	return encodeTitle
}

func decode(title string) string {
	var decodeTitle = ""
	for _, t := range title {
		for i := range data {
			if string(t) == data[i] {
				if i-3 < 0 {
					decodeTitle += data[len(data)+i-3]
				} else {
					decodeTitle += data[i-3]
				}
			}
		}
	}

	return decodeTitle
}
