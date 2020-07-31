package crypto

import (
	"fmt"
)

func AsciiEncode(data string) string {
	var result string
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		if i == 0 {
			result += fmt.Sprint(array[i])
		} else {
			result += fmt.Sprint(",", array[i])
		}
	}
	return result
}

func AsciiSUM(data string) string {
	var result rune
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		result += array[i]
	}
	return fmt.Sprint(result)
}
